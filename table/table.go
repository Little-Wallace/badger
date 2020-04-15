/*
 * Copyright 2017 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package table

import (
	"fmt"
	"math"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/coocood/badger/cache"
	"github.com/coocood/badger/fileutil"
	"github.com/coocood/badger/options"
	"github.com/coocood/badger/surf"
	"github.com/coocood/badger/y"
	"github.com/coocood/bbloom"
	"github.com/golang/snappy"
	"github.com/pingcap/errors"
)

const (
	fileSuffix    = ".sst"
	idxFileSuffix = ".idx"

	intSize = int(unsafe.Sizeof(int(0)))
)

func IndexFilename(tableFilename string) string { return tableFilename + idxFileSuffix }

type tableIndex struct {
	blockEndOffsets []uint32
	baseKeysEndOffs []uint32
	baseKeys        []byte
	bf              *bbloom.Bloom
	hIdx            *hashIndex
	surf            *surf.SuRF
}

func (i *tableIndex) getBlockBaseKey(idx int) []byte {
	baseKeyStartOff := 0
	if idx > 0 {
		baseKeyStartOff = int(i.baseKeysEndOffs[idx-1])
	}
	baseKeyEndOff := i.baseKeysEndOffs[idx]
	return i.baseKeys[baseKeyStartOff:baseKeyEndOff]
}

// Table represents a loaded table file with the info we have about it
type Table struct {
	sync.Mutex

	fd      *os.File // Own fd.
	indexFd *os.File

	globalTs          uint64
	tableSize         int64
	numBlocks         int
	smallest, biggest y.Key
	id                uint64

	blockCache *cache.Cache
	blocksMmap []byte

	indexCache *cache.Cache
	index      *tableIndex
	indexMmap  []byte

	compacting int32

	compression options.CompressionType
}

// CompressionType returns the compression algorithm used for block compression.
func (t *Table) CompressionType() options.CompressionType {
	return t.compression
}

// Delete delete table's file from disk.
func (t *Table) Delete() error {
	if t.blockCache != nil {
		for blk := 0; blk < t.numBlocks; blk++ {
			t.blockCache.Del(t.blockCacheKey(blk))
		}
	}
	if t.indexCache != nil {
		t.indexCache.Del(t.id)
	}
	if len(t.blocksMmap) != 0 {
		y.Munmap(t.blocksMmap)
	}
	t.index = nil
	if len(t.indexMmap) != 0 {
		y.Munmap(t.indexMmap)
	}
	if err := t.fd.Truncate(0); err != nil {
		// This is very important to let the FS know that the file is deleted.
		return err
	}
	filename := t.fd.Name()
	if err := t.fd.Close(); err != nil {
		return err
	}
	if err := os.Remove(filename); err != nil {
		return err
	}
	return os.Remove(filename + idxFileSuffix)
}

// OpenTable assumes file has only one table and opens it.  Takes ownership of fd upon function
// entry.  Returns a table with one reference count on it (decrementing which may delete the file!
// -- consider t.Close() instead).  The fd has to writeable because we call Truncate on it before
// deleting.
func OpenTable(filename string, compression options.CompressionType, blockCache *cache.Cache, indexCache *cache.Cache) (*Table, error) {
	id, ok := ParseFileID(filename)
	if !ok {
		return nil, errors.Errorf("Invalid filename: %s", filename)
	}

	// TODO: after we support cache of L2 storage, we will open block data file in cache manager.
	fd, err := y.OpenExistingFile(filename, 0)
	if err != nil {
		return nil, err
	}

	indexFd, err := y.OpenExistingFile(filename+idxFileSuffix, 0)
	if err != nil {
		return nil, err
	}

	t := &Table{
		fd:          fd,
		indexFd:     indexFd,
		id:          id,
		compression: compression,
		blockCache:  blockCache,
		indexCache:  indexCache,
	}

	if err := t.initTableInfo(); err != nil {
		t.Close()
		return nil, err
	}

	if blockCache == nil {
		t.blocksMmap, err = y.Mmap(fd, false, t.Size())
		if err != nil {
			t.Close()
			return nil, y.Wrapf(err, "Unable to map file")
		}
	}
	return t, nil
}

// Close closes the open table.  (Releases resources back to the OS.)
func (t *Table) Close() error {
	if t.fd != nil {
		t.fd.Close()
	}
	if t.indexFd != nil {
		if len(t.indexMmap) != 0 {
			y.Munmap(t.indexMmap)
		}
		t.indexFd.Close()
	}
	return nil
}

// PointGet try to lookup a key and its value by table's hash index.
// If it find an hash collision the last return value will be false,
// which means caller should fallback to seek search. Otherwise it value will be true.
// If the hash index does not contain such an element the returned key will be nil.
func (t *Table) PointGet(key y.Key, keyHash uint64) (y.Key, y.ValueStruct, bool, error) {
	idx, err := t.getIndex()
	if err != nil {
		return y.Key{}, y.ValueStruct{}, false, err
	}
	if idx.bf != nil && !idx.bf.Has(keyHash) {
		return y.Key{}, y.ValueStruct{}, true, err
	}

	blkIdx, offset := uint32(resultFallback), uint8(0)
	if idx.hIdx != nil {
		blkIdx, offset = idx.hIdx.lookup(keyHash)
	} else if idx.surf != nil {
		v, ok := idx.surf.Get(key.UserKey)
		if !ok {
			blkIdx = resultNoEntry
		} else {
			var pos entryPosition
			pos.decode(v)
			blkIdx, offset = uint32(pos.blockIdx), pos.offset
		}
	}
	if blkIdx == resultFallback {
		return y.Key{}, y.ValueStruct{}, false, nil
	}
	if blkIdx == resultNoEntry {
		return y.Key{}, y.ValueStruct{}, true, nil
	}

	it := t.NewIterator(false)
	it.seekFromOffset(int(blkIdx), int(offset), key)

	if !it.Valid() || !key.SameUserKey(it.Key()) {
		return y.Key{}, y.ValueStruct{}, true, it.Error()
	}
	return it.Key(), it.Value(), true, nil
}

func (t *Table) read(off int, sz int) ([]byte, error) {
	if len(t.blocksMmap) > 0 {
		if len(t.blocksMmap[off:]) < sz {
			return nil, y.ErrEOF
		}
		return t.blocksMmap[off : off+sz], nil
	}
	res := make([]byte, sz)
	_, err := t.fd.ReadAt(res, int64(off))
	return res, err
}

func (t *Table) readNoFail(off int, sz int) []byte {
	res, err := t.read(off, sz)
	y.Check(err)
	return res
}

func (t *Table) initTableInfo() error {
	data, err := t.loadIndexData(false)
	if err != nil {
		return err
	}
	for d := (metaDecoder{buf: data}); d.valid(); d.next() {
		switch d.currentId() {
		case idSmallest:
			if k := d.decode(); len(k) != 0 {
				t.smallest = y.KeyWithTs(y.Copy(k), math.MaxUint64)
			}
		case idBiggest:
			if k := d.decode(); len(k) != 0 {
				t.biggest = y.KeyWithTs(y.Copy(k), 0)
			}
		case idBlockEndOffsets:
			offsets := bytesToU32Slice(d.decode())
			t.tableSize = int64(offsets[len(offsets)-1])
			t.numBlocks = len(offsets)
		}
	}
	return nil
}

func (t *Table) readTableIndex(data []byte) *tableIndex {
	idx := new(tableIndex)
	for d := (metaDecoder{buf: data}); d.valid(); d.next() {
		switch d.currentId() {
		case idBaseKeysEndOffs:
			idx.baseKeysEndOffs = bytesToU32Slice(d.decode())
		case idBaseKeys:
			idx.baseKeys = d.decode()
		case idBlockEndOffsets:
			idx.blockEndOffsets = bytesToU32Slice(d.decode())
		case idBloomFilter:
			if d := d.decode(); len(d) != 0 {
				idx.bf = new(bbloom.Bloom)
				idx.bf.BinaryUnmarshal(d)
			}
		case idHashIndex:
			if d := d.decode(); len(d) != 0 {
				idx.hIdx = new(hashIndex)
				idx.hIdx.readIndex(d)
			}
		case idSuRFIndex:
			if d := d.decode(); len(d) != 0 {
				idx.surf = new(surf.SuRF)
				idx.surf.Unmarshal(d)
			}
		}
	}
	return idx
}

func (t *Table) getIndex() (*tableIndex, error) {
	if t.index != nil {
		return t.index, nil
	}
	if t.indexCache == nil {
		idxData, err := t.loadIndexData(true)
		if err != nil {
			return nil, err
		}
		t.index = t.readTableIndex(idxData)
		return t.index, nil
	}

	index, err := t.indexCache.GetOrCompute(t.id, func() (interface{}, int64, error) {
		d, err := t.loadIndexData(false)
		if err != nil {
			return nil, 0, err
		}
		return t.readTableIndex(d), int64(len(d)), nil
	})
	if err != nil {
		return nil, err
	}
	return index.(*tableIndex), nil
}

func (t *Table) loadIndexData(useMmap bool) ([]byte, error) {
	fstat, err := t.indexFd.Stat()
	if err != nil {
		return nil, err
	}
	var idxData []byte

	if useMmap {
		idxData, err = y.Mmap(t.indexFd, false, fstat.Size())
		if err != nil {
			return nil, err
		}
		t.indexMmap = idxData
	} else {
		idxData = make([]byte, fstat.Size())
		if _, err = t.indexFd.ReadAt(idxData, 0); err != nil {
			return nil, err
		}
	}

	t.globalTs = bytesToU64(idxData[:8])
	idxData = idxData[8:]
	if t.compression != options.None {
		old := idxData
		if idxData, err = t.decompressData(idxData); err != nil {
			return nil, err
		}
		if useMmap {
			y.Munmap(old)
		}
		t.indexMmap = nil
	}
	return idxData, nil
}

type block struct {
	offset  int
	data    []byte
	baseKey []byte
}

func (b *block) size() int64 {
	return int64(intSize + len(b.data))
}

func (t *Table) block(idx int, index *tableIndex) (block, error) {
	y.Assert(idx >= 0)

	if idx >= len(index.blockEndOffsets) {
		return block{}, errors.New("block out of index")
	}

	if t.blockCache == nil {
		return t.loadBlock(idx, index)
	}

	key := t.blockCacheKey(idx)
	blk, err := t.blockCache.GetOrCompute(key, func() (interface{}, int64, error) {
		b, e := t.loadBlock(idx, index)
		if e != nil {
			return nil, 0, e
		}
		return b, int64(len(b.data)), nil
	})
	if err != nil {
		return block{}, err
	}
	return blk.(block), nil
}

func (t *Table) loadBlock(idx int, index *tableIndex) (block, error) {
	var startOffset int
	if idx > 0 {
		startOffset = int(index.blockEndOffsets[idx-1])
	}
	blk := block{
		offset: startOffset,
	}
	endOffset := int(index.blockEndOffsets[idx])
	dataLen := endOffset - startOffset
	var err error
	if blk.data, err = t.read(blk.offset, dataLen); err != nil {
		return block{}, errors.Wrapf(err,
			"failed to read from file: %s at offset: %d, len: %d", t.fd.Name(), blk.offset, dataLen)
	}

	blk.data, err = t.decompressData(blk.data)
	if err != nil {
		return block{}, errors.Wrapf(err,
			"failed to decode compressed data in file: %s at offset: %d, len: %d",
			t.fd.Name(), blk.offset, dataLen)
	}
	blk.baseKey = index.getBlockBaseKey(idx)
	return blk, nil
}

// HasGlobalTs returns table does set global ts.
func (t *Table) HasGlobalTs() bool {
	return t.globalTs != 0
}

// SetGlobalTs update the global ts of external ingested tables.
func (t *Table) SetGlobalTs(ts uint64) error {
	if _, err := t.indexFd.WriteAt(u64ToBytes(ts), 0); err != nil {
		return err
	}
	if err := fileutil.Fsync(t.indexFd); err != nil {
		return err
	}
	t.globalTs = ts
	return nil
}

func (t *Table) MarkCompacting(flag bool) {
	if flag {
		atomic.StoreInt32(&t.compacting, 1)
	}
	atomic.StoreInt32(&t.compacting, 0)
}

func (t *Table) IsCompacting() bool {
	return atomic.LoadInt32(&t.compacting) == 1
}

func (t *Table) blockCacheKey(idx int) uint64 {
	y.Assert(t.ID() < math.MaxUint32)
	y.Assert(idx < math.MaxUint32)
	return (t.ID() << 32) | uint64(idx)
}

// Size is its file size in bytes
func (t *Table) Size() int64 { return t.tableSize }

// Smallest is its smallest key, or nil if there are none
func (t *Table) Smallest() y.Key { return t.smallest }

// Biggest is its biggest key, or nil if there are none
func (t *Table) Biggest() y.Key { return t.biggest }

// Filename is NOT the file name.  Just kidding, it is.
func (t *Table) Filename() string { return t.fd.Name() }

// ID is the table's ID number (used to make the file name).
func (t *Table) ID() uint64 { return t.id }

func (t *Table) HasOverlap(start, end y.Key, includeEnd bool) bool {
	if start.Compare(t.Biggest()) > 0 {
		return false
	}

	if cmp := end.Compare(t.Smallest()); cmp < 0 {
		return false
	} else if cmp == 0 {
		return includeEnd
	}

	idx, err := t.getIndex()
	if err != nil {
		return true
	}

	if idx.surf != nil {
		return idx.surf.HasOverlap(start.UserKey, end.UserKey, includeEnd)
	}

	// If there are errors occurred during seeking,
	// we assume the table has overlapped with the range to prevent data loss.
	it := t.newIterator(false, idx)
	it.Seek(start)
	if !it.Valid() {
		return it.Error() != nil
	}
	if cmp := it.Key().Compare(end); cmp > 0 {
		return false
	} else if cmp == 0 {
		return includeEnd
	}
	return true
}

// ParseFileID reads the file id out of a filename.
func ParseFileID(name string) (uint64, bool) {
	name = path.Base(name)
	if !strings.HasSuffix(name, fileSuffix) {
		return 0, false
	}
	//	suffix := name[len(fileSuffix):]
	name = strings.TrimSuffix(name, fileSuffix)
	id, err := strconv.ParseUint(name, 16, 64)
	if err != nil {
		return 0, false
	}
	return id, true
}

// IDToFilename does the inverse of ParseFileID
func IDToFilename(id uint64) string {
	return fmt.Sprintf("%08x", id) + fileSuffix
}

// NewFilename should be named TableFilepath -- it combines the dir with the ID to make a table
// filepath.
func NewFilename(id uint64, dir string) string {
	return filepath.Join(dir, IDToFilename(id))
}

// decompressData decompresses the given data.
func (t *Table) decompressData(data []byte) ([]byte, error) {
	switch t.compression {
	case options.None:
		return data, nil
	case options.Snappy:
		return snappy.Decode(nil, data)
	case options.ZSTD:
		return decompress(data)
	}
	return nil, errors.New("Unsupported compression type")
}

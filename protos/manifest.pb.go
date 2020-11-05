// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: manifest.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ManifestChange_Operation int32

const (
	ManifestChange_CREATE    ManifestChange_Operation = 0
	ManifestChange_DELETE    ManifestChange_Operation = 1
	ManifestChange_MOVE_DOWN ManifestChange_Operation = 2
)

var ManifestChange_Operation_name = map[int32]string{
	0: "CREATE",
	1: "DELETE",
	2: "MOVE_DOWN",
}

var ManifestChange_Operation_value = map[string]int32{
	"CREATE":    0,
	"DELETE":    1,
	"MOVE_DOWN": 2,
}

func (x ManifestChange_Operation) String() string {
	return proto.EnumName(ManifestChange_Operation_name, int32(x))
}

func (ManifestChange_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{2, 0}
}

type ShardChange_Operation int32

const (
	ShardChange_CREATE   ShardChange_Operation = 0
	ShardChange_DELETE   ShardChange_Operation = 1
	ShardChange_TRUNCATE ShardChange_Operation = 2
)

var ShardChange_Operation_name = map[int32]string{
	0: "CREATE",
	1: "DELETE",
	2: "TRUNCATE",
}

var ShardChange_Operation_value = map[string]int32{
	"CREATE":   0,
	"DELETE":   1,
	"TRUNCATE": 2,
}

func (x ShardChange_Operation) String() string {
	return proto.EnumName(ShardChange_Operation_name, int32(x))
}

func (ShardChange_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{3, 0}
}

type ManifestChangeSet struct {
	// A set of changes that are applied atomically.
	Changes              []*ManifestChange `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
	Head                 *HeadInfo         `protobuf:"bytes,2,opt,name=head,proto3" json:"head,omitempty"`
	ShardChange          []*ShardChange    `protobuf:"bytes,3,rep,name=shardChange,proto3" json:"shardChange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ManifestChangeSet) Reset()         { *m = ManifestChangeSet{} }
func (m *ManifestChangeSet) String() string { return proto.CompactTextString(m) }
func (*ManifestChangeSet) ProtoMessage()    {}
func (*ManifestChangeSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{0}
}
func (m *ManifestChangeSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManifestChangeSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManifestChangeSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ManifestChangeSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestChangeSet.Merge(m, src)
}
func (m *ManifestChangeSet) XXX_Size() int {
	return m.Size()
}
func (m *ManifestChangeSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestChangeSet.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestChangeSet proto.InternalMessageInfo

func (m *ManifestChangeSet) GetChanges() []*ManifestChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

func (m *ManifestChangeSet) GetHead() *HeadInfo {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *ManifestChangeSet) GetShardChange() []*ShardChange {
	if m != nil {
		return m.ShardChange
	}
	return nil
}

type HeadInfo struct {
	Version              uint64   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	LogID                uint32   `protobuf:"varint,2,opt,name=logID,proto3" json:"logID,omitempty"`
	LogOffset            uint32   `protobuf:"varint,3,opt,name=logOffset,proto3" json:"logOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeadInfo) Reset()         { *m = HeadInfo{} }
func (m *HeadInfo) String() string { return proto.CompactTextString(m) }
func (*HeadInfo) ProtoMessage()    {}
func (*HeadInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{1}
}
func (m *HeadInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeadInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeadInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeadInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadInfo.Merge(m, src)
}
func (m *HeadInfo) XXX_Size() int {
	return m.Size()
}
func (m *HeadInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HeadInfo proto.InternalMessageInfo

func (m *HeadInfo) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HeadInfo) GetLogID() uint32 {
	if m != nil {
		return m.LogID
	}
	return 0
}

func (m *HeadInfo) GetLogOffset() uint32 {
	if m != nil {
		return m.LogOffset
	}
	return 0
}

type ManifestChange struct {
	Id                   uint64                   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Op                   ManifestChange_Operation `protobuf:"varint,2,opt,name=Op,proto3,enum=protos.ManifestChange_Operation" json:"Op,omitempty"`
	Level                uint32                   `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	CF                   int32                    `protobuf:"varint,4,opt,name=CF,proto3" json:"CF,omitempty"`
	ShardID              uint32                   `protobuf:"varint,5,opt,name=ShardID,proto3" json:"ShardID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ManifestChange) Reset()         { *m = ManifestChange{} }
func (m *ManifestChange) String() string { return proto.CompactTextString(m) }
func (*ManifestChange) ProtoMessage()    {}
func (*ManifestChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{2}
}
func (m *ManifestChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManifestChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManifestChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ManifestChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestChange.Merge(m, src)
}
func (m *ManifestChange) XXX_Size() int {
	return m.Size()
}
func (m *ManifestChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestChange.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestChange proto.InternalMessageInfo

func (m *ManifestChange) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ManifestChange) GetOp() ManifestChange_Operation {
	if m != nil {
		return m.Op
	}
	return ManifestChange_CREATE
}

func (m *ManifestChange) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *ManifestChange) GetCF() int32 {
	if m != nil {
		return m.CF
	}
	return 0
}

func (m *ManifestChange) GetShardID() uint32 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

type ShardChange struct {
	ShardID              uint32                `protobuf:"varint,1,opt,name=shardID,proto3" json:"shardID,omitempty"`
	Op                   ShardChange_Operation `protobuf:"varint,2,opt,name=Op,proto3,enum=protos.ShardChange_Operation" json:"Op,omitempty"`
	StartKey             []byte                `protobuf:"bytes,3,opt,name=startKey,proto3" json:"startKey,omitempty"`
	EndKey               []byte                `protobuf:"bytes,4,opt,name=endKey,proto3" json:"endKey,omitempty"`
	TableIDs             []uint32              `protobuf:"varint,5,rep,packed,name=tableIDs,proto3" json:"tableIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ShardChange) Reset()         { *m = ShardChange{} }
func (m *ShardChange) String() string { return proto.CompactTextString(m) }
func (*ShardChange) ProtoMessage()    {}
func (*ShardChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{3}
}
func (m *ShardChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShardChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShardChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShardChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardChange.Merge(m, src)
}
func (m *ShardChange) XXX_Size() int {
	return m.Size()
}
func (m *ShardChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardChange.DiscardUnknown(m)
}

var xxx_messageInfo_ShardChange proto.InternalMessageInfo

func (m *ShardChange) GetShardID() uint32 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *ShardChange) GetOp() ShardChange_Operation {
	if m != nil {
		return m.Op
	}
	return ShardChange_CREATE
}

func (m *ShardChange) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *ShardChange) GetEndKey() []byte {
	if m != nil {
		return m.EndKey
	}
	return nil
}

func (m *ShardChange) GetTableIDs() []uint32 {
	if m != nil {
		return m.TableIDs
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.ManifestChange_Operation", ManifestChange_Operation_name, ManifestChange_Operation_value)
	proto.RegisterEnum("protos.ShardChange_Operation", ShardChange_Operation_name, ShardChange_Operation_value)
	proto.RegisterType((*ManifestChangeSet)(nil), "protos.ManifestChangeSet")
	proto.RegisterType((*HeadInfo)(nil), "protos.HeadInfo")
	proto.RegisterType((*ManifestChange)(nil), "protos.ManifestChange")
	proto.RegisterType((*ShardChange)(nil), "protos.ShardChange")
}

func init() { proto.RegisterFile("manifest.proto", fileDescriptor_0bb23f43f7afb4c1) }

var fileDescriptor_0bb23f43f7afb4c1 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x3b, 0x93, 0x9f, 0x26, 0x37, 0x8d, 0x65, 0x06, 0x54, 0x8d, 0x10, 0x58, 0x96, 0xc5,
	0xc2, 0x1b, 0xa2, 0x62, 0xc4, 0x03, 0x80, 0xed, 0x0a, 0x8b, 0xb6, 0x96, 0xa6, 0xe1, 0x67, 0x87,
	0xa6, 0x78, 0x9c, 0x44, 0x32, 0xb6, 0xe5, 0x19, 0x55, 0xe2, 0x4d, 0x58, 0xf3, 0x30, 0x88, 0x25,
	0x5b, 0x76, 0x28, 0xbc, 0x08, 0x9a, 0xb1, 0x9d, 0xa6, 0xa8, 0x0b, 0x56, 0xf6, 0xf1, 0x39, 0xf7,
	0xf3, 0x9c, 0xab, 0x01, 0xeb, 0x33, 0x2f, 0x37, 0xb9, 0x90, 0x6a, 0x51, 0x37, 0x95, 0xaa, 0xc8,
	0xd8, 0x3c, 0xa4, 0xf7, 0x0d, 0xc1, 0xbd, 0xf3, 0xce, 0x0a, 0xd7, 0xbc, 0x5c, 0x89, 0x4b, 0xa1,
	0xc8, 0x09, 0x1c, 0x7e, 0x32, 0x42, 0x52, 0xe4, 0x0e, 0xfc, 0x59, 0x70, 0xdc, 0x8e, 0xc9, 0xc5,
	0xed, 0x2c, 0xeb, 0x63, 0xe4, 0x09, 0x0c, 0xd7, 0x82, 0x67, 0x14, 0xbb, 0xc8, 0x9f, 0x05, 0x76,
	0x1f, 0x7f, 0x2d, 0x78, 0x96, 0x94, 0x79, 0xc5, 0x8c, 0x4b, 0x5e, 0xc0, 0x4c, 0xae, 0x79, 0x93,
	0xb5, 0xd3, 0x74, 0x60, 0xd8, 0xf7, 0xfb, 0xf0, 0xe5, 0x8d, 0xc5, 0xf6, 0x73, 0xde, 0x07, 0x98,
	0xf4, 0x20, 0x42, 0xe1, 0xf0, 0x5a, 0x34, 0x72, 0x53, 0x95, 0x14, 0xb9, 0xc8, 0x1f, 0xb2, 0x5e,
	0x92, 0x07, 0x30, 0x2a, 0xaa, 0x55, 0x12, 0x99, 0x33, 0xcc, 0x59, 0x2b, 0xc8, 0x23, 0x98, 0x16,
	0xd5, 0x2a, 0xcd, 0x73, 0x29, 0x14, 0x1d, 0x18, 0xe7, 0xe6, 0x83, 0xf7, 0x1d, 0x81, 0x75, 0xbb,
	0x12, 0xb1, 0x00, 0x27, 0x59, 0xc7, 0xc6, 0x49, 0x46, 0x4e, 0x00, 0xa7, 0xb5, 0x61, 0x5a, 0x81,
	0x7b, 0xf7, 0x1a, 0x16, 0x69, 0x2d, 0x1a, 0xae, 0x36, 0x55, 0xc9, 0x70, 0x5a, 0xeb, 0x83, 0x9c,
	0x89, 0x6b, 0x51, 0x74, 0xbf, 0x6b, 0x85, 0xe6, 0x86, 0xa7, 0x74, 0xe8, 0x22, 0x7f, 0xc4, 0x70,
	0x78, 0xaa, 0x8b, 0x98, 0xc2, 0x49, 0x44, 0x47, 0x26, 0xd7, 0x4b, 0x2f, 0x80, 0xe9, 0x0e, 0x48,
	0x00, 0xc6, 0x21, 0x8b, 0x5f, 0x2e, 0x63, 0xfb, 0x40, 0xbf, 0x47, 0xf1, 0x59, 0xbc, 0x8c, 0x6d,
	0x44, 0xe6, 0x30, 0x3d, 0x4f, 0xdf, 0xc5, 0x1f, 0xa3, 0xf4, 0xfd, 0x85, 0x8d, 0xbd, 0x5f, 0x08,
	0x66, 0x7b, 0xfb, 0xd3, 0x74, 0xd9, 0xd1, 0x51, 0x4b, 0xef, 0x24, 0x79, 0xba, 0xd7, 0xe7, 0xf1,
	0x1d, 0xab, 0xff, 0xa7, 0xcc, 0x43, 0x98, 0x48, 0xc5, 0x1b, 0xf5, 0x46, 0x7c, 0x31, 0x7d, 0x8e,
	0xd8, 0x4e, 0x93, 0x63, 0x18, 0x8b, 0x32, 0xd3, 0xce, 0xd0, 0x38, 0x9d, 0xd2, 0x33, 0x8a, 0x5f,
	0x15, 0x22, 0x89, 0x24, 0x1d, 0xb9, 0x03, 0x7f, 0xce, 0x76, 0xda, 0x7b, 0xf6, 0x3f, 0xe5, 0x8e,
	0x60, 0xb2, 0x64, 0x6f, 0x2f, 0x42, 0xed, 0xe0, 0x57, 0xf6, 0x8f, 0xad, 0x83, 0x7e, 0x6e, 0x1d,
	0xf4, 0x7b, 0xeb, 0xa0, 0xaf, 0x7f, 0x9c, 0x83, 0xab, 0xf6, 0xf6, 0x3e, 0xff, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0xbb, 0x73, 0xc9, 0xd6, 0x02, 0x00, 0x00,
}

func (m *ManifestChangeSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestChangeSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ManifestChangeSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ShardChange) > 0 {
		for iNdEx := len(m.ShardChange) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ShardChange[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintManifest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Head != nil {
		{
			size, err := m.Head.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintManifest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Changes) > 0 {
		for iNdEx := len(m.Changes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Changes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintManifest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *HeadInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeadInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeadInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LogOffset != 0 {
		i = encodeVarintManifest(dAtA, i, uint64(m.LogOffset))
		i--
		dAtA[i] = 0x18
	}
	if m.LogID != 0 {
		i = encodeVarintManifest(dAtA, i, uint64(m.LogID))
		i--
		dAtA[i] = 0x10
	}
	if m.Version != 0 {
		i = encodeVarintManifest(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ManifestChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ManifestChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ShardID != 0 {
		i = encodeVarintManifest(dAtA, i, uint64(m.ShardID))
		i--
		dAtA[i] = 0x28
	}
	if m.CF != 0 {
		i = encodeVarintManifest(dAtA, i, uint64(m.CF))
		i--
		dAtA[i] = 0x20
	}
	if m.Level != 0 {
		i = encodeVarintManifest(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x18
	}
	if m.Op != 0 {
		i = encodeVarintManifest(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintManifest(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ShardChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShardChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShardChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.TableIDs) > 0 {
		dAtA3 := make([]byte, len(m.TableIDs)*10)
		var j2 int
		for _, num := range m.TableIDs {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintManifest(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.EndKey) > 0 {
		i -= len(m.EndKey)
		copy(dAtA[i:], m.EndKey)
		i = encodeVarintManifest(dAtA, i, uint64(len(m.EndKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.StartKey) > 0 {
		i -= len(m.StartKey)
		copy(dAtA[i:], m.StartKey)
		i = encodeVarintManifest(dAtA, i, uint64(len(m.StartKey)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Op != 0 {
		i = encodeVarintManifest(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x10
	}
	if m.ShardID != 0 {
		i = encodeVarintManifest(dAtA, i, uint64(m.ShardID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintManifest(dAtA []byte, offset int, v uint64) int {
	offset -= sovManifest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ManifestChangeSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for _, e := range m.Changes {
			l = e.Size()
			n += 1 + l + sovManifest(uint64(l))
		}
	}
	if m.Head != nil {
		l = m.Head.Size()
		n += 1 + l + sovManifest(uint64(l))
	}
	if len(m.ShardChange) > 0 {
		for _, e := range m.ShardChange {
			l = e.Size()
			n += 1 + l + sovManifest(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HeadInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovManifest(uint64(m.Version))
	}
	if m.LogID != 0 {
		n += 1 + sovManifest(uint64(m.LogID))
	}
	if m.LogOffset != 0 {
		n += 1 + sovManifest(uint64(m.LogOffset))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ManifestChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovManifest(uint64(m.Id))
	}
	if m.Op != 0 {
		n += 1 + sovManifest(uint64(m.Op))
	}
	if m.Level != 0 {
		n += 1 + sovManifest(uint64(m.Level))
	}
	if m.CF != 0 {
		n += 1 + sovManifest(uint64(m.CF))
	}
	if m.ShardID != 0 {
		n += 1 + sovManifest(uint64(m.ShardID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ShardChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ShardID != 0 {
		n += 1 + sovManifest(uint64(m.ShardID))
	}
	if m.Op != 0 {
		n += 1 + sovManifest(uint64(m.Op))
	}
	l = len(m.StartKey)
	if l > 0 {
		n += 1 + l + sovManifest(uint64(l))
	}
	l = len(m.EndKey)
	if l > 0 {
		n += 1 + l + sovManifest(uint64(l))
	}
	if len(m.TableIDs) > 0 {
		l = 0
		for _, e := range m.TableIDs {
			l += sovManifest(uint64(e))
		}
		n += 1 + sovManifest(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovManifest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozManifest(x uint64) (n int) {
	return sovManifest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ManifestChangeSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManifest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ManifestChangeSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestChangeSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthManifest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changes = append(m.Changes, &ManifestChange{})
			if err := m.Changes[len(m.Changes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Head", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthManifest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Head == nil {
				m.Head = &HeadInfo{}
			}
			if err := m.Head.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardChange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthManifest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShardChange = append(m.ShardChange, &ShardChange{})
			if err := m.ShardChange[len(m.ShardChange)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipManifest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManifest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManifest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HeadInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManifest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeadInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeadInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogID", wireType)
			}
			m.LogID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogOffset", wireType)
			}
			m.LogOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogOffset |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipManifest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManifest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManifest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ManifestChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManifest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ManifestChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= ManifestChange_Operation(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CF", wireType)
			}
			m.CF = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CF |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardID", wireType)
			}
			m.ShardID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShardID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipManifest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManifest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManifest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ShardChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManifest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ShardChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShardChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardID", wireType)
			}
			m.ShardID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShardID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= ShardChange_Operation(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthManifest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartKey = append(m.StartKey[:0], dAtA[iNdEx:postIndex]...)
			if m.StartKey == nil {
				m.StartKey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthManifest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndKey = append(m.EndKey[:0], dAtA[iNdEx:postIndex]...)
			if m.EndKey == nil {
				m.EndKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowManifest
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.TableIDs = append(m.TableIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowManifest
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthManifest
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthManifest
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.TableIDs) == 0 {
					m.TableIDs = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowManifest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.TableIDs = append(m.TableIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field TableIDs", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipManifest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManifest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManifest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipManifest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowManifest
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthManifest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupManifest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthManifest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthManifest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowManifest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupManifest = fmt.Errorf("proto: unexpected end of group")
)

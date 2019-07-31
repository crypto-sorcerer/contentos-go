// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_blockProducerVote.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type SoBlockProducerVote struct {
	VoterId              *prototype.BpVoterId    `protobuf:"bytes,1,opt,name=voter_id,json=voterId,proto3" json:"voter_id,omitempty"`
	WitnessId            *prototype.BpWitnessId  `protobuf:"bytes,2,opt,name=witness_id,json=witnessId,proto3" json:"witness_id,omitempty"`
	VoteTime             *prototype.TimePointSec `protobuf:"bytes,3,opt,name=vote_time,json=voteTime,proto3" json:"vote_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoBlockProducerVote) Reset()         { *m = SoBlockProducerVote{} }
func (m *SoBlockProducerVote) String() string { return proto.CompactTextString(m) }
func (*SoBlockProducerVote) ProtoMessage()    {}
func (*SoBlockProducerVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d279fa3103478f1, []int{0}
}

func (m *SoBlockProducerVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoBlockProducerVote.Unmarshal(m, b)
}
func (m *SoBlockProducerVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoBlockProducerVote.Marshal(b, m, deterministic)
}
func (m *SoBlockProducerVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoBlockProducerVote.Merge(m, src)
}
func (m *SoBlockProducerVote) XXX_Size() int {
	return xxx_messageInfo_SoBlockProducerVote.Size(m)
}
func (m *SoBlockProducerVote) XXX_DiscardUnknown() {
	xxx_messageInfo_SoBlockProducerVote.DiscardUnknown(m)
}

var xxx_messageInfo_SoBlockProducerVote proto.InternalMessageInfo

func (m *SoBlockProducerVote) GetVoterId() *prototype.BpVoterId {
	if m != nil {
		return m.VoterId
	}
	return nil
}

func (m *SoBlockProducerVote) GetWitnessId() *prototype.BpWitnessId {
	if m != nil {
		return m.WitnessId
	}
	return nil
}

func (m *SoBlockProducerVote) GetVoteTime() *prototype.TimePointSec {
	if m != nil {
		return m.VoteTime
	}
	return nil
}

type SoMemBlockProducerVoteByVoterId struct {
	VoterId              *prototype.BpVoterId `protobuf:"bytes,1,opt,name=voter_id,json=voterId,proto3" json:"voter_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SoMemBlockProducerVoteByVoterId) Reset()         { *m = SoMemBlockProducerVoteByVoterId{} }
func (m *SoMemBlockProducerVoteByVoterId) String() string { return proto.CompactTextString(m) }
func (*SoMemBlockProducerVoteByVoterId) ProtoMessage()    {}
func (*SoMemBlockProducerVoteByVoterId) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d279fa3103478f1, []int{1}
}

func (m *SoMemBlockProducerVoteByVoterId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemBlockProducerVoteByVoterId.Unmarshal(m, b)
}
func (m *SoMemBlockProducerVoteByVoterId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemBlockProducerVoteByVoterId.Marshal(b, m, deterministic)
}
func (m *SoMemBlockProducerVoteByVoterId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemBlockProducerVoteByVoterId.Merge(m, src)
}
func (m *SoMemBlockProducerVoteByVoterId) XXX_Size() int {
	return xxx_messageInfo_SoMemBlockProducerVoteByVoterId.Size(m)
}
func (m *SoMemBlockProducerVoteByVoterId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemBlockProducerVoteByVoterId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemBlockProducerVoteByVoterId proto.InternalMessageInfo

func (m *SoMemBlockProducerVoteByVoterId) GetVoterId() *prototype.BpVoterId {
	if m != nil {
		return m.VoterId
	}
	return nil
}

type SoMemBlockProducerVoteByWitnessId struct {
	WitnessId            *prototype.BpWitnessId `protobuf:"bytes,1,opt,name=witness_id,json=witnessId,proto3" json:"witness_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoMemBlockProducerVoteByWitnessId) Reset()         { *m = SoMemBlockProducerVoteByWitnessId{} }
func (m *SoMemBlockProducerVoteByWitnessId) String() string { return proto.CompactTextString(m) }
func (*SoMemBlockProducerVoteByWitnessId) ProtoMessage()    {}
func (*SoMemBlockProducerVoteByWitnessId) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d279fa3103478f1, []int{2}
}

func (m *SoMemBlockProducerVoteByWitnessId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemBlockProducerVoteByWitnessId.Unmarshal(m, b)
}
func (m *SoMemBlockProducerVoteByWitnessId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemBlockProducerVoteByWitnessId.Marshal(b, m, deterministic)
}
func (m *SoMemBlockProducerVoteByWitnessId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemBlockProducerVoteByWitnessId.Merge(m, src)
}
func (m *SoMemBlockProducerVoteByWitnessId) XXX_Size() int {
	return xxx_messageInfo_SoMemBlockProducerVoteByWitnessId.Size(m)
}
func (m *SoMemBlockProducerVoteByWitnessId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemBlockProducerVoteByWitnessId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemBlockProducerVoteByWitnessId proto.InternalMessageInfo

func (m *SoMemBlockProducerVoteByWitnessId) GetWitnessId() *prototype.BpWitnessId {
	if m != nil {
		return m.WitnessId
	}
	return nil
}

type SoMemBlockProducerVoteByVoteTime struct {
	VoteTime             *prototype.TimePointSec `protobuf:"bytes,1,opt,name=vote_time,json=voteTime,proto3" json:"vote_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemBlockProducerVoteByVoteTime) Reset()         { *m = SoMemBlockProducerVoteByVoteTime{} }
func (m *SoMemBlockProducerVoteByVoteTime) String() string { return proto.CompactTextString(m) }
func (*SoMemBlockProducerVoteByVoteTime) ProtoMessage()    {}
func (*SoMemBlockProducerVoteByVoteTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d279fa3103478f1, []int{3}
}

func (m *SoMemBlockProducerVoteByVoteTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemBlockProducerVoteByVoteTime.Unmarshal(m, b)
}
func (m *SoMemBlockProducerVoteByVoteTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemBlockProducerVoteByVoteTime.Marshal(b, m, deterministic)
}
func (m *SoMemBlockProducerVoteByVoteTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemBlockProducerVoteByVoteTime.Merge(m, src)
}
func (m *SoMemBlockProducerVoteByVoteTime) XXX_Size() int {
	return xxx_messageInfo_SoMemBlockProducerVoteByVoteTime.Size(m)
}
func (m *SoMemBlockProducerVoteByVoteTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemBlockProducerVoteByVoteTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemBlockProducerVoteByVoteTime proto.InternalMessageInfo

func (m *SoMemBlockProducerVoteByVoteTime) GetVoteTime() *prototype.TimePointSec {
	if m != nil {
		return m.VoteTime
	}
	return nil
}

type SoListBlockProducerVoteByVoterId struct {
	VoterId              *prototype.BpVoterId `protobuf:"bytes,1,opt,name=voter_id,json=voterId,proto3" json:"voter_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SoListBlockProducerVoteByVoterId) Reset()         { *m = SoListBlockProducerVoteByVoterId{} }
func (m *SoListBlockProducerVoteByVoterId) String() string { return proto.CompactTextString(m) }
func (*SoListBlockProducerVoteByVoterId) ProtoMessage()    {}
func (*SoListBlockProducerVoteByVoterId) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d279fa3103478f1, []int{4}
}

func (m *SoListBlockProducerVoteByVoterId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListBlockProducerVoteByVoterId.Unmarshal(m, b)
}
func (m *SoListBlockProducerVoteByVoterId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListBlockProducerVoteByVoterId.Marshal(b, m, deterministic)
}
func (m *SoListBlockProducerVoteByVoterId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListBlockProducerVoteByVoterId.Merge(m, src)
}
func (m *SoListBlockProducerVoteByVoterId) XXX_Size() int {
	return xxx_messageInfo_SoListBlockProducerVoteByVoterId.Size(m)
}
func (m *SoListBlockProducerVoteByVoterId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListBlockProducerVoteByVoterId.DiscardUnknown(m)
}

var xxx_messageInfo_SoListBlockProducerVoteByVoterId proto.InternalMessageInfo

func (m *SoListBlockProducerVoteByVoterId) GetVoterId() *prototype.BpVoterId {
	if m != nil {
		return m.VoterId
	}
	return nil
}

type SoUniqueBlockProducerVoteByVoterId struct {
	VoterId              *prototype.BpVoterId `protobuf:"bytes,1,opt,name=voter_id,json=voterId,proto3" json:"voter_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SoUniqueBlockProducerVoteByVoterId) Reset()         { *m = SoUniqueBlockProducerVoteByVoterId{} }
func (m *SoUniqueBlockProducerVoteByVoterId) String() string { return proto.CompactTextString(m) }
func (*SoUniqueBlockProducerVoteByVoterId) ProtoMessage()    {}
func (*SoUniqueBlockProducerVoteByVoterId) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d279fa3103478f1, []int{5}
}

func (m *SoUniqueBlockProducerVoteByVoterId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueBlockProducerVoteByVoterId.Unmarshal(m, b)
}
func (m *SoUniqueBlockProducerVoteByVoterId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueBlockProducerVoteByVoterId.Marshal(b, m, deterministic)
}
func (m *SoUniqueBlockProducerVoteByVoterId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueBlockProducerVoteByVoterId.Merge(m, src)
}
func (m *SoUniqueBlockProducerVoteByVoterId) XXX_Size() int {
	return xxx_messageInfo_SoUniqueBlockProducerVoteByVoterId.Size(m)
}
func (m *SoUniqueBlockProducerVoteByVoterId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueBlockProducerVoteByVoterId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueBlockProducerVoteByVoterId proto.InternalMessageInfo

func (m *SoUniqueBlockProducerVoteByVoterId) GetVoterId() *prototype.BpVoterId {
	if m != nil {
		return m.VoterId
	}
	return nil
}

func init() {
	proto.RegisterType((*SoBlockProducerVote)(nil), "table.so_blockProducerVote")
	proto.RegisterType((*SoMemBlockProducerVoteByVoterId)(nil), "table.so_mem_blockProducerVote_by_voter_id")
	proto.RegisterType((*SoMemBlockProducerVoteByWitnessId)(nil), "table.so_mem_blockProducerVote_by_witness_id")
	proto.RegisterType((*SoMemBlockProducerVoteByVoteTime)(nil), "table.so_mem_blockProducerVote_by_vote_time")
	proto.RegisterType((*SoListBlockProducerVoteByVoterId)(nil), "table.so_list_blockProducerVote_by_voter_id")
	proto.RegisterType((*SoUniqueBlockProducerVoteByVoterId)(nil), "table.so_unique_blockProducerVote_by_voter_id")
}

func init() {
	proto.RegisterFile("app/table/so_blockProducerVote.proto", fileDescriptor_5d279fa3103478f1)
}

var fileDescriptor_5d279fa3103478f1 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4f, 0x4b, 0x3b, 0x31,
	0x10, 0x65, 0x7f, 0x3f, 0xfc, 0xd3, 0x78, 0x5b, 0x8a, 0xac, 0x3d, 0xc9, 0x52, 0xff, 0x20, 0xba,
	0x41, 0x05, 0xbd, 0x7b, 0xeb, 0x4d, 0x8a, 0x08, 0x16, 0x21, 0x6c, 0xb2, 0x43, 0x1b, 0xdc, 0x64,
	0xe2, 0x66, 0x56, 0xe9, 0x57, 0xf3, 0xd3, 0xc9, 0xa6, 0xda, 0xb5, 0x2a, 0x5a, 0xa4, 0x97, 0x30,
	0x99, 0x99, 0xf7, 0xf2, 0xde, 0x83, 0xb0, 0x7e, 0xee, 0x1c, 0xa7, 0x5c, 0x96, 0xc0, 0x3d, 0x0a,
	0x59, 0xa2, 0x7a, 0xb8, 0xae, 0xb0, 0xa8, 0x15, 0x54, 0xb7, 0x48, 0x90, 0xb9, 0x0a, 0x09, 0xe3,
	0xb5, 0xb0, 0xd1, 0x4b, 0xc2, 0x8d, 0xa6, 0x0e, 0xb8, 0xa9, 0x4b, 0xd2, 0x42, 0x17, 0xb3, 0x85,
	0x5e, 0xb7, 0x9d, 0x34, 0xc7, 0xac, 0x9b, 0xbe, 0x44, 0xac, 0xfb, 0x1d, 0x6b, 0x7c, 0xca, 0x36,
	0x9f, 0x90, 0xa0, 0x12, 0xba, 0x48, 0xa2, 0xdd, 0xe8, 0x70, 0xeb, 0x6c, 0x3b, 0x9b, 0x33, 0x64,
	0xd2, 0x89, 0xf7, 0xe9, 0x70, 0x23, 0x54, 0x83, 0x22, 0xbe, 0x64, 0xec, 0x59, 0x93, 0x05, 0xef,
	0x1b, 0xd0, 0xbf, 0x00, 0x4a, 0x16, 0x41, 0xed, 0x7c, 0xd8, 0x79, 0xab, 0x07, 0x45, 0x7c, 0xc1,
	0x3a, 0x0d, 0x87, 0x20, 0x6d, 0x20, 0xf9, 0x1f, 0x70, 0x3b, 0x1f, 0x70, 0x4d, 0x5b, 0x38, 0xd4,
	0x96, 0x84, 0x07, 0x35, 0x0c, 0xba, 0x6e, 0xb4, 0x81, 0xf4, 0x8e, 0xf5, 0x3d, 0x0a, 0x03, 0xe6,
	0xab, 0x7e, 0x21, 0xa7, 0x73, 0x85, 0x7f, 0xf0, 0x92, 0xe6, 0x6c, 0xff, 0x27, 0xea, 0xd6, 0xc7,
	0x27, 0xd7, 0xd1, 0xd2, 0xae, 0x53, 0xc1, 0xf6, 0x7e, 0x53, 0x1f, 0x12, 0x59, 0x8c, 0x27, 0x5a,
	0x3e, 0x9e, 0x51, 0x78, 0xa0, 0xd4, 0x9e, 0x56, 0x9f, 0xcf, 0x3d, 0x3b, 0xf0, 0x28, 0x6a, 0xab,
	0x1f, 0x6b, 0x58, 0x39, 0xfb, 0xd5, 0xf1, 0xe8, 0x68, 0xac, 0x69, 0x52, 0xcb, 0x4c, 0xa1, 0xe1,
	0x0a, 0xbd, 0x9a, 0xe4, 0xda, 0x72, 0x85, 0x96, 0xc0, 0x12, 0xfa, 0x93, 0x31, 0xf2, 0xf9, 0xaf,
	0x90, 0xeb, 0x81, 0xed, 0xfc, 0x35, 0x00, 0x00, 0xff, 0xff, 0x73, 0x94, 0xb2, 0xf4, 0x29, 0x03,
	0x00, 0x00,
}

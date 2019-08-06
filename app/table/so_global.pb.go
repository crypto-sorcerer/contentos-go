// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_global.proto

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

type SoGlobal struct {
	Id                   int32                        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Props                *prototype.DynamicProperties `protobuf:"bytes,2,opt,name=props,proto3" json:"props,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoGlobal) Reset()         { *m = SoGlobal{} }
func (m *SoGlobal) String() string { return proto.CompactTextString(m) }
func (*SoGlobal) ProtoMessage()    {}
func (*SoGlobal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b138d292b065e72, []int{0}
}

func (m *SoGlobal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoGlobal.Unmarshal(m, b)
}
func (m *SoGlobal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoGlobal.Marshal(b, m, deterministic)
}
func (m *SoGlobal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoGlobal.Merge(m, src)
}
func (m *SoGlobal) XXX_Size() int {
	return xxx_messageInfo_SoGlobal.Size(m)
}
func (m *SoGlobal) XXX_DiscardUnknown() {
	xxx_messageInfo_SoGlobal.DiscardUnknown(m)
}

var xxx_messageInfo_SoGlobal proto.InternalMessageInfo

func (m *SoGlobal) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SoGlobal) GetProps() *prototype.DynamicProperties {
	if m != nil {
		return m.Props
	}
	return nil
}

type SoUniqueGlobalById struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueGlobalById) Reset()         { *m = SoUniqueGlobalById{} }
func (m *SoUniqueGlobalById) String() string { return proto.CompactTextString(m) }
func (*SoUniqueGlobalById) ProtoMessage()    {}
func (*SoUniqueGlobalById) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b138d292b065e72, []int{1}
}

func (m *SoUniqueGlobalById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueGlobalById.Unmarshal(m, b)
}
func (m *SoUniqueGlobalById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueGlobalById.Marshal(b, m, deterministic)
}
func (m *SoUniqueGlobalById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueGlobalById.Merge(m, src)
}
func (m *SoUniqueGlobalById) XXX_Size() int {
	return xxx_messageInfo_SoUniqueGlobalById.Size(m)
}
func (m *SoUniqueGlobalById) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueGlobalById.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueGlobalById proto.InternalMessageInfo

func (m *SoUniqueGlobalById) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*SoGlobal)(nil), "table.so_global")
	proto.RegisterType((*SoUniqueGlobalById)(nil), "table.so_unique_global_by_id")
}

func init() { proto.RegisterFile("app/table/so_global.proto", fileDescriptor_8b138d292b065e72) }

var fileDescriptor_8b138d292b065e72 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xcd, 0xaa, 0x83, 0x30,
	0x10, 0x85, 0x51, 0xf0, 0xc2, 0xcd, 0x85, 0xbb, 0x90, 0x52, 0x6c, 0xa1, 0x20, 0xae, 0xa4, 0xb4,
	0x09, 0xd4, 0x37, 0xe8, 0x13, 0x14, 0x97, 0xdd, 0x84, 0xfc, 0xa1, 0x03, 0x9a, 0x49, 0x4d, 0x5c,
	0xf8, 0xf6, 0x45, 0x2d, 0x2e, 0xba, 0x19, 0x98, 0x39, 0x1f, 0xdf, 0x70, 0xc8, 0x41, 0x38, 0xc7,
	0x82, 0x90, 0x9d, 0x61, 0x1e, 0x79, 0xd3, 0xa1, 0x14, 0x1d, 0x75, 0x03, 0x06, 0x4c, 0x93, 0xe5,
	0x7c, 0xdc, 0x2d, 0x5b, 0x98, 0x9c, 0x61, 0xf3, 0x58, 0xc3, 0xe2, 0x41, 0x7e, 0x37, 0x3e, 0xfd,
	0x27, 0x31, 0xe8, 0x2c, 0xca, 0xa3, 0x32, 0xa9, 0x63, 0xd0, 0x69, 0x45, 0x12, 0x37, 0xa0, 0xf3,
	0x59, 0x9c, 0x47, 0xe5, 0xdf, 0xed, 0x44, 0x37, 0x05, 0xd5, 0x93, 0x15, 0x3d, 0x28, 0x3e, 0xe7,
	0x66, 0x08, 0x60, 0x7c, 0xbd, 0xb2, 0x45, 0x49, 0xf6, 0x1e, 0xf9, 0x68, 0xe1, 0x35, 0x9a, 0x8f,
	0x98, 0xcb, 0x89, 0x83, 0xfe, 0xd6, 0xdf, 0x2f, 0xcf, 0x73, 0x03, 0xa1, 0x1d, 0x25, 0x55, 0xd8,
	0x33, 0x85, 0x5e, 0xb5, 0x02, 0x2c, 0x53, 0x68, 0x83, 0xb1, 0x01, 0xfd, 0xb5, 0x41, 0xb6, 0xd5,
	0x92, 0x3f, 0xcb, 0xf3, 0xea, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x81, 0x57, 0x62, 0xda, 0xea, 0x00,
	0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test_plugin.proto

package testdata

import (
	fmt "fmt"
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

type MarshaledSignedData struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Identity             []byte   `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarshaledSignedData) Reset()         { *m = MarshaledSignedData{} }
func (m *MarshaledSignedData) String() string { return proto.CompactTextString(m) }
func (*MarshaledSignedData) ProtoMessage()    {}
func (*MarshaledSignedData) Descriptor() ([]byte, []int) {
	return fileDescriptor_65f5abc8dfb3e1dd, []int{0}
}

func (m *MarshaledSignedData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarshaledSignedData.Unmarshal(m, b)
}
func (m *MarshaledSignedData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarshaledSignedData.Marshal(b, m, deterministic)
}
func (m *MarshaledSignedData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarshaledSignedData.Merge(m, src)
}
func (m *MarshaledSignedData) XXX_Size() int {
	return xxx_messageInfo_MarshaledSignedData.Size(m)
}
func (m *MarshaledSignedData) XXX_DiscardUnknown() {
	xxx_messageInfo_MarshaledSignedData.DiscardUnknown(m)
}

var xxx_messageInfo_MarshaledSignedData proto.InternalMessageInfo

func (m *MarshaledSignedData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MarshaledSignedData) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *MarshaledSignedData) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func init() {
	proto.RegisterType((*MarshaledSignedData)(nil), "testdata.MarshaledSignedData")
}

func init() { proto.RegisterFile("test_plugin.proto", fileDescriptor_65f5abc8dfb3e1dd) }

var fileDescriptor_65f5abc8dfb3e1dd = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x46, 0x71, 0x5b, 0x8a, 0x2b, 0xba, 0x54, 0x5d, 0x4c, 0xe9, 0x50, 0x3a, 0x75, 0xb2, 0x28,
	0xed, 0x2f, 0x28, 0x0d, 0x99, 0xb2, 0x24, 0x5b, 0x96, 0x70, 0x96, 0x2e, 0xf2, 0x81, 0x2c, 0x99,
	0xf3, 0xd9, 0xc4, 0xff, 0x3e, 0x58, 0x90, 0x64, 0xfb, 0xbe, 0xf7, 0x96, 0xa7, 0x5e, 0x04, 0x07,
	0x39, 0xf4, 0x61, 0xf4, 0x14, 0xeb, 0x9e, 0x93, 0x24, 0x5d, 0x2e, 0xc8, 0x81, 0xc0, 0xa7, 0x55,
	0xaf, 0x1b, 0xe0, 0xa1, 0x85, 0x80, 0x6e, 0x47, 0x3e, 0xa2, 0xfb, 0x07, 0x01, 0xad, 0xd5, 0xc3,
	0xa2, 0xab, 0xe2, 0xa3, 0xf8, 0x7a, 0xde, 0xe6, 0xad, 0xdf, 0xd5, 0xd3, 0x40, 0x3e, 0x82, 0x8c,
	0x8c, 0xd5, 0x5d, 0x16, 0x37, 0xa0, 0xdf, 0x54, 0x49, 0x0e, 0xa3, 0x90, 0xcc, 0xd5, 0x7d, 0x96,
	0xd7, 0xff, 0xb7, 0xde, 0xaf, 0x3c, 0x49, 0x3b, 0x36, 0xb5, 0x4d, 0x9d, 0x69, 0xe7, 0x1e, 0x39,
	0xa0, 0xf3, 0xc8, 0xe6, 0x08, 0x0d, 0x93, 0x35, 0x36, 0x31, 0x1a, 0x9b, 0xba, 0x8e, 0x44, 0x90,
	0x8d, 0x9c, 0x26, 0x08, 0xe4, 0x40, 0x12, 0x9b, 0xe9, 0xfb, 0xd7, 0x5c, 0x6a, 0x9b, 0xc7, 0x9c,
	0xff, 0x73, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x64, 0xb2, 0x7a, 0x4a, 0xd3, 0x00, 0x00, 0x00,
}

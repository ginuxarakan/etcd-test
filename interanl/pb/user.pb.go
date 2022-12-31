// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

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

type UserReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReq) Reset()         { *m = UserReq{} }
func (m *UserReq) String() string { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()    {}
func (*UserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReq.Unmarshal(m, b)
}
func (m *UserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReq.Marshal(b, m, deterministic)
}
func (m *UserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReq.Merge(m, src)
}
func (m *UserReq) XXX_Size() int {
	return xxx_messageInfo_UserReq.Size(m)
}
func (m *UserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserReq proto.InternalMessageInfo

type UserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResp) Reset()         { *m = UserResp{} }
func (m *UserResp) String() string { return proto.CompactTextString(m) }
func (*UserResp) ProtoMessage()    {}
func (*UserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResp.Unmarshal(m, b)
}
func (m *UserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResp.Marshal(b, m, deterministic)
}
func (m *UserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResp.Merge(m, src)
}
func (m *UserResp) XXX_Size() int {
	return xxx_messageInfo_UserResp.Size(m)
}
func (m *UserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserReq)(nil), "pb.UserReq")
	proto.RegisterType((*UserResp)(nil), "pb.UserResp")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xe2, 0xe4, 0x62, 0x0f, 0x2d,
	0x4e, 0x2d, 0x0a, 0x4a, 0x2d, 0x54, 0xe2, 0xe2, 0xe2, 0x80, 0x30, 0x8b, 0x0b, 0x8c, 0x2c, 0xb8,
	0xb8, 0x41, 0xec, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x4d, 0x2e, 0x1e, 0x10, 0xd7,
	0x39, 0x31, 0x27, 0x27, 0x24, 0xb5, 0xb8, 0x44, 0x88, 0x5b, 0xaf, 0x20, 0x49, 0x0f, 0xaa, 0x4f,
	0x8a, 0x07, 0xc1, 0x29, 0x2e, 0x70, 0x62, 0x8b, 0x62, 0xd1, 0xb3, 0x2e, 0x48, 0x4a, 0x62, 0x03,
	0xdb, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x9c, 0x3a, 0x0a, 0x71, 0x00, 0x00, 0x00,
}
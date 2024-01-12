// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

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

type AddRequest struct {
	N1                   int32    `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
	N2                   int32    `protobuf:"varint,2,opt,name=n2,proto3" json:"n2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func (m *AddRequest) GetN2() int32 {
	if m != nil {
		return m.N2
	}
	return 0
}

type AddReply struct {
	N1                   int32    `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddReply) Reset()         { *m = AddReply{} }
func (m *AddReply) String() string { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()    {}
func (*AddReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *AddReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReply.Unmarshal(m, b)
}
func (m *AddReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReply.Marshal(b, m, deterministic)
}
func (m *AddReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReply.Merge(m, src)
}
func (m *AddReply) XXX_Size() int {
	return xxx_messageInfo_AddReply.Size(m)
}
func (m *AddReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddReply proto.InternalMessageInfo

func (m *AddReply) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

type SubtractRequest struct {
	N1                   int32    `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
	N2                   int32    `protobuf:"varint,2,opt,name=n2,proto3" json:"n2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubtractRequest) Reset()         { *m = SubtractRequest{} }
func (m *SubtractRequest) String() string { return proto.CompactTextString(m) }
func (*SubtractRequest) ProtoMessage()    {}
func (*SubtractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *SubtractRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubtractRequest.Unmarshal(m, b)
}
func (m *SubtractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubtractRequest.Marshal(b, m, deterministic)
}
func (m *SubtractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubtractRequest.Merge(m, src)
}
func (m *SubtractRequest) XXX_Size() int {
	return xxx_messageInfo_SubtractRequest.Size(m)
}
func (m *SubtractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubtractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubtractRequest proto.InternalMessageInfo

func (m *SubtractRequest) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func (m *SubtractRequest) GetN2() int32 {
	if m != nil {
		return m.N2
	}
	return 0
}

type SubtractReply struct {
	N1                   int32    `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubtractReply) Reset()         { *m = SubtractReply{} }
func (m *SubtractReply) String() string { return proto.CompactTextString(m) }
func (*SubtractReply) ProtoMessage()    {}
func (*SubtractReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *SubtractReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubtractReply.Unmarshal(m, b)
}
func (m *SubtractReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubtractReply.Marshal(b, m, deterministic)
}
func (m *SubtractReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubtractReply.Merge(m, src)
}
func (m *SubtractReply) XXX_Size() int {
	return xxx_messageInfo_SubtractReply.Size(m)
}
func (m *SubtractReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubtractReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubtractReply proto.InternalMessageInfo

func (m *SubtractReply) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

type MultiplyRequest struct {
	N1                   int32    `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
	N2                   int32    `protobuf:"varint,2,opt,name=n2,proto3" json:"n2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplyRequest) Reset()         { *m = MultiplyRequest{} }
func (m *MultiplyRequest) String() string { return proto.CompactTextString(m) }
func (*MultiplyRequest) ProtoMessage()    {}
func (*MultiplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *MultiplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplyRequest.Unmarshal(m, b)
}
func (m *MultiplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplyRequest.Marshal(b, m, deterministic)
}
func (m *MultiplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplyRequest.Merge(m, src)
}
func (m *MultiplyRequest) XXX_Size() int {
	return xxx_messageInfo_MultiplyRequest.Size(m)
}
func (m *MultiplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplyRequest proto.InternalMessageInfo

func (m *MultiplyRequest) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func (m *MultiplyRequest) GetN2() int32 {
	if m != nil {
		return m.N2
	}
	return 0
}

type MultiplyReply struct {
	N1                   int32    `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplyReply) Reset()         { *m = MultiplyReply{} }
func (m *MultiplyReply) String() string { return proto.CompactTextString(m) }
func (*MultiplyReply) ProtoMessage()    {}
func (*MultiplyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *MultiplyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplyReply.Unmarshal(m, b)
}
func (m *MultiplyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplyReply.Marshal(b, m, deterministic)
}
func (m *MultiplyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplyReply.Merge(m, src)
}
func (m *MultiplyReply) XXX_Size() int {
	return xxx_messageInfo_MultiplyReply.Size(m)
}
func (m *MultiplyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplyReply.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplyReply proto.InternalMessageInfo

func (m *MultiplyReply) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

type DivideRequest struct {
	N1                   int32    `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
	N2                   int32    `protobuf:"varint,2,opt,name=n2,proto3" json:"n2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DivideRequest) Reset()         { *m = DivideRequest{} }
func (m *DivideRequest) String() string { return proto.CompactTextString(m) }
func (*DivideRequest) ProtoMessage()    {}
func (*DivideRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *DivideRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DivideRequest.Unmarshal(m, b)
}
func (m *DivideRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DivideRequest.Marshal(b, m, deterministic)
}
func (m *DivideRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DivideRequest.Merge(m, src)
}
func (m *DivideRequest) XXX_Size() int {
	return xxx_messageInfo_DivideRequest.Size(m)
}
func (m *DivideRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DivideRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DivideRequest proto.InternalMessageInfo

func (m *DivideRequest) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func (m *DivideRequest) GetN2() int32 {
	if m != nil {
		return m.N2
	}
	return 0
}

type DivideReply struct {
	N1                   int32    `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DivideReply) Reset()         { *m = DivideReply{} }
func (m *DivideReply) String() string { return proto.CompactTextString(m) }
func (*DivideReply) ProtoMessage()    {}
func (*DivideReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{7}
}

func (m *DivideReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DivideReply.Unmarshal(m, b)
}
func (m *DivideReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DivideReply.Marshal(b, m, deterministic)
}
func (m *DivideReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DivideReply.Merge(m, src)
}
func (m *DivideReply) XXX_Size() int {
	return xxx_messageInfo_DivideReply.Size(m)
}
func (m *DivideReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DivideReply.DiscardUnknown(m)
}

var xxx_messageInfo_DivideReply proto.InternalMessageInfo

func (m *DivideReply) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "calc.AddRequest")
	proto.RegisterType((*AddReply)(nil), "calc.AddReply")
	proto.RegisterType((*SubtractRequest)(nil), "calc.SubtractRequest")
	proto.RegisterType((*SubtractReply)(nil), "calc.SubtractReply")
	proto.RegisterType((*MultiplyRequest)(nil), "calc.MultiplyRequest")
	proto.RegisterType((*MultiplyReply)(nil), "calc.MultiplyReply")
	proto.RegisterType((*DivideRequest)(nil), "calc.DivideRequest")
	proto.RegisterType((*DivideReply)(nil), "calc.DivideReply")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x47, 0x30, 0x0b, 0x92, 0x90, 0x38, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0x11, 0x25, 0x1d, 0x2e, 0x2e, 0xc7, 0x94, 0x94, 0xa0, 0xd4,
	0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x3e, 0x2e, 0xa6, 0x3c, 0x43, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xd6, 0x20, 0xa6, 0x3c, 0x43, 0x30, 0xdf, 0x48, 0x82, 0x09, 0xca, 0x37, 0x52, 0x92, 0xe2, 0xe2,
	0x00, 0xab, 0x2e, 0xc8, 0xa9, 0x44, 0x57, 0xab, 0x64, 0xc8, 0xc5, 0x1f, 0x5c, 0x9a, 0x54, 0x52,
	0x94, 0x98, 0x5c, 0x42, 0xac, 0x71, 0xf2, 0x5c, 0xbc, 0x08, 0x2d, 0x38, 0xcc, 0xf4, 0x2d, 0xcd,
	0x29, 0xc9, 0x2c, 0xc8, 0xa9, 0x24, 0xc1, 0x4c, 0x84, 0x16, 0x6c, 0x66, 0xea, 0x73, 0xf1, 0xba,
	0x64, 0x96, 0x65, 0xa6, 0xa4, 0x12, 0x6b, 0xa2, 0x2c, 0x17, 0x37, 0x4c, 0x03, 0x16, 0xf3, 0x8c,
	0xee, 0x33, 0x72, 0x71, 0x39, 0xc3, 0x03, 0x57, 0x48, 0x93, 0x8b, 0xd9, 0x31, 0x25, 0x45, 0x48,
	0x40, 0x0f, 0x14, 0xbc, 0x7a, 0x88, 0xb0, 0x95, 0xe2, 0x43, 0x12, 0x29, 0xc8, 0xa9, 0x54, 0x62,
	0x10, 0xb2, 0xe0, 0xe2, 0x80, 0x79, 0x5f, 0x48, 0x14, 0x22, 0x8b, 0x16, 0x82, 0x52, 0xc2, 0xe8,
	0xc2, 0x70, 0x9d, 0x30, 0x4f, 0xc2, 0x74, 0xa2, 0x85, 0x13, 0x4c, 0x27, 0x4a, 0x58, 0x28, 0x31,
	0x08, 0x19, 0x71, 0xb1, 0x41, 0x3c, 0x23, 0x04, 0x55, 0x80, 0x12, 0x16, 0x52, 0x82, 0xa8, 0x82,
	0x60, 0x3d, 0x4e, 0x92, 0x51, 0xe2, 0x38, 0x52, 0x55, 0x12, 0x1b, 0x38, 0x2d, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x5b, 0xb6, 0x1b, 0x4b, 0x77, 0x02, 0x00, 0x00,
}
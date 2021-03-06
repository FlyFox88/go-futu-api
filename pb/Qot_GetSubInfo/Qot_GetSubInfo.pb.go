// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetSubInfo.proto

package Qot_GetSubInfo

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

type C2S struct {
	IsReqAllConn         *bool    `protobuf:"varint,1,opt,name=isReqAllConn" json:"isReqAllConn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5551a2176c8a1a6, []int{0}
}

func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (m *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(m, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetIsReqAllConn() bool {
	if m != nil && m.IsReqAllConn != nil {
		return *m.IsReqAllConn
	}
	return false
}

type S2C struct {
	ConnSubInfoList      []*ConnSubInfo `protobuf:"bytes,1,rep,name=connSubInfoList" json:"connSubInfoList,omitempty"`
	TotalUsedQuota       *int32         `protobuf:"varint,2,req,name=totalUsedQuota" json:"totalUsedQuota,omitempty"`
	RemainQuota          *int32         `protobuf:"varint,3,req,name=remainQuota" json:"remainQuota,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5551a2176c8a1a6, []int{1}
}

func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (m *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(m, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetConnSubInfoList() []*ConnSubInfo {
	if m != nil {
		return m.ConnSubInfoList
	}
	return nil
}

func (m *S2C) GetTotalUsedQuota() int32 {
	if m != nil && m.TotalUsedQuota != nil {
		return *m.TotalUsedQuota
	}
	return 0
}

func (m *S2C) GetRemainQuota() int32 {
	if m != nil && m.RemainQuota != nil {
		return *m.RemainQuota
	}
	return 0
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5551a2176c8a1a6, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5551a2176c8a1a6, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "Qot_GetSubInfo.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetSubInfo.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetSubInfo.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetSubInfo.Response")
}

func init() { proto.RegisterFile("Qot_GetSubInfo.proto", fileDescriptor_b5551a2176c8a1a6) }

var fileDescriptor_b5551a2176c8a1a6 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4b, 0x7b, 0x31,
	0x14, 0x40, 0xc9, 0x4b, 0xfb, 0x6b, 0x7f, 0xb7, 0xa5, 0x4a, 0x14, 0x0d, 0x1d, 0x24, 0x3c, 0x50,
	0xe2, 0x60, 0x29, 0xc1, 0xc9, 0xad, 0x64, 0x10, 0x41, 0x87, 0xe6, 0xe9, 0x2c, 0xb5, 0xbd, 0x4a,
	0xa1, 0xcd, 0x6d, 0x93, 0x74, 0x70, 0xf6, 0x1b, 0xf8, 0x89, 0xe5, 0xf5, 0x0f, 0xb6, 0x0f, 0xc7,
	0x7b, 0xce, 0x49, 0x72, 0x09, 0x9c, 0x0e, 0x29, 0xbd, 0xde, 0x63, 0x2a, 0x56, 0x6f, 0x0f, 0xfe,
	0x9d, 0x7a, 0x8b, 0x40, 0x89, 0x44, 0xe7, 0x90, 0x76, 0xdb, 0x96, 0xe6, 0x73, 0xf2, 0x1b, 0xdb,
	0x3d, 0x2e, 0xed, 0x3e, 0xc9, 0xaf, 0x81, 0x5b, 0x53, 0x88, 0x1c, 0xda, 0xd3, 0xe8, 0x70, 0x39,
	0x98, 0xcd, 0x2c, 0x79, 0x2f, 0x99, 0x62, 0xba, 0xe9, 0x0e, 0x58, 0xfe, 0xcd, 0x80, 0x17, 0xc6,
	0x8a, 0x01, 0x1c, 0x8d, 0xc9, 0xfb, 0xed, 0x0b, 0x8f, 0xd3, 0x98, 0x24, 0x53, 0x5c, 0xb7, 0xcc,
	0x79, 0x6f, 0xef, 0x7a, 0xfb, 0x9b, 0xb8, 0x6a, 0x2f, 0xae, 0xa0, 0x93, 0x28, 0x8d, 0x66, 0x2f,
	0x11, 0x27, 0xc3, 0x15, 0xa5, 0x91, 0xcc, 0x54, 0xa6, 0xeb, 0xae, 0x42, 0x85, 0x82, 0x56, 0xc0,
	0xf9, 0x68, 0xea, 0x37, 0x11, 0x5f, 0x47, 0xfb, 0x28, 0xef, 0x43, 0xc3, 0xe1, 0x72, 0x85, 0x31,
	0x89, 0x4b, 0xe0, 0x63, 0x13, 0x25, 0x53, 0x99, 0x6e, 0x99, 0x93, 0x5e, 0xe5, 0x7b, 0xac, 0x29,
	0x5c, 0xe9, 0xf3, 0x2f, 0x06, 0x4d, 0x87, 0x71, 0x41, 0x3e, 0xa2, 0xb8, 0x80, 0x46, 0xc0, 0xf4,
	0xfc, 0xb9, 0xc0, 0xf5, 0xb9, 0xfa, 0x5d, 0xed, 0xe6, 0xb6, 0xdf, 0x77, 0x3b, 0x28, 0xce, 0xe0,
	0x5f, 0xc0, 0xf4, 0x14, 0x3f, 0x64, 0xa6, 0x98, 0xfe, 0xef, 0xb6, 0x93, 0x90, 0xd0, 0xc0, 0x10,
	0x2c, 0x4d, 0x50, 0x72, 0xc5, 0x74, 0xdd, 0xed, 0xc6, 0x72, 0x8b, 0x68, 0xc6, 0xb2, 0xa6, 0xd8,
	0x5f, 0x5b, 0x14, 0xc6, 0xba, 0xd2, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xed, 0xe7, 0x0e,
	0xbe, 0x01, 0x00, 0x00,
}

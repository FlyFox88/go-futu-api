// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetTicker.proto

package Qot_GetTicker

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
	Security             *Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	MaxRetNum            *int32    `protobuf:"varint,2,req,name=maxRetNum" json:"maxRetNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_712c6e9383ef0c5e, []int{0}
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

func (m *C2S) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *C2S) GetMaxRetNum() int32 {
	if m != nil && m.MaxRetNum != nil {
		return *m.MaxRetNum
	}
	return 0
}

type S2C struct {
	Security             *Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	TickerList           []*Ticker `protobuf:"bytes,2,rep,name=tickerList" json:"tickerList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_712c6e9383ef0c5e, []int{1}
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

func (m *S2C) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *S2C) GetTickerList() []*Ticker {
	if m != nil {
		return m.TickerList
	}
	return nil
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
	return fileDescriptor_712c6e9383ef0c5e, []int{2}
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
	return fileDescriptor_712c6e9383ef0c5e, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetTicker.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetTicker.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetTicker.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetTicker.Response")
}

func init() { proto.RegisterFile("Qot_GetTicker.proto", fileDescriptor_712c6e9383ef0c5e) }

var fileDescriptor_712c6e9383ef0c5e = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x69, 0xbb, 0xd9, 0xed, 0x9b, 0x82, 0x44, 0x91, 0x30, 0x44, 0x42, 0xf1, 0xd0, 0x8b,
	0xb3, 0x04, 0x4f, 0x5e, 0x73, 0xf0, 0xa2, 0x82, 0x5f, 0xe7, 0x59, 0xa4, 0x7e, 0x48, 0x19, 0x5d,
	0x6a, 0xf2, 0x15, 0xdc, 0xd5, 0xbf, 0x5c, 0xba, 0x76, 0xae, 0x82, 0x27, 0x8f, 0xef, 0xfd, 0xde,
	0x4b, 0x5e, 0x02, 0x27, 0x4f, 0x96, 0x5f, 0xee, 0x88, 0x97, 0x65, 0xb1, 0x22, 0xb7, 0xa8, 0x9d,
	0x65, 0x2b, 0x8e, 0x7e, 0x99, 0xf3, 0x43, 0x63, 0xab, 0xca, 0xae, 0x3b, 0x38, 0x3f, 0x6e, 0xe1,
	0xd0, 0x49, 0x9e, 0x21, 0x32, 0x3a, 0x17, 0x19, 0x4c, 0x3c, 0x15, 0x8d, 0x2b, 0x79, 0x23, 0x03,
	0x15, 0xa6, 0x33, 0x7d, 0xba, 0x18, 0x64, 0xf3, 0x9e, 0xe1, 0x4f, 0x4a, 0x9c, 0xc3, 0xb4, 0x7a,
	0xfd, 0x44, 0xe2, 0xc7, 0xa6, 0x92, 0xa1, 0x0a, 0xd3, 0x31, 0xee, 0x8d, 0x64, 0x05, 0x51, 0xae,
	0xcd, 0x3f, 0x8e, 0xd5, 0x00, 0xbc, 0x5d, 0x7e, 0x5f, 0x7a, 0x96, 0xa1, 0x8a, 0xd2, 0x99, 0x16,
	0xc3, 0x4e, 0xf7, 0x2e, 0x1c, 0xa4, 0x92, 0x6b, 0x88, 0x91, 0x3e, 0x1a, 0xf2, 0x2c, 0x2e, 0x21,
	0x2a, 0xb4, 0xef, 0xef, 0xea, 0x7a, 0xfb, 0x0f, 0x32, 0x3a, 0xc7, 0x16, 0x27, 0x5f, 0x01, 0x4c,
	0x90, 0x7c, 0x6d, 0xd7, 0x9e, 0xc4, 0x05, 0xc4, 0x8e, 0x78, 0xb9, 0xa9, 0x69, 0x5b, 0x1b, 0xdf,
	0x8e, 0xae, 0x6e, 0xb2, 0x0c, 0x77, 0xa6, 0x38, 0x83, 0x03, 0x47, 0xfc, 0xe0, 0xdf, 0x65, 0xa8,
	0x82, 0x74, 0x8a, 0xbd, 0x12, 0x12, 0x62, 0x72, 0xce, 0xd8, 0x37, 0x92, 0x91, 0x0a, 0xd2, 0x31,
	0xee, 0x64, 0x3b, 0xc2, 0xeb, 0x42, 0x8e, 0x54, 0xf0, 0xc7, 0x88, 0x5c, 0x1b, 0x6c, 0xf1, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0xac, 0xc1, 0x6c, 0xbe, 0x01, 0x00, 0x00,
}

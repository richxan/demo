// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pay/proto/pay.proto

package proto

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

type PaymentRequest struct {
	Amount               uint64   `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	OutTradeNo           string   `protobuf:"bytes,2,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentRequest) Reset()         { *m = PaymentRequest{} }
func (m *PaymentRequest) String() string { return proto.CompactTextString(m) }
func (*PaymentRequest) ProtoMessage()    {}
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f4653c102df3c6, []int{0}
}

func (m *PaymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentRequest.Unmarshal(m, b)
}
func (m *PaymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentRequest.Marshal(b, m, deterministic)
}
func (m *PaymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentRequest.Merge(m, src)
}
func (m *PaymentRequest) XXX_Size() int {
	return xxx_messageInfo_PaymentRequest.Size(m)
}
func (m *PaymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentRequest proto.InternalMessageInfo

func (m *PaymentRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PaymentRequest) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

type PaymentResponse struct {
	TransactionId        string   `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	OutTradeNo           string   `protobuf:"bytes,2,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentResponse) Reset()         { *m = PaymentResponse{} }
func (m *PaymentResponse) String() string { return proto.CompactTextString(m) }
func (*PaymentResponse) ProtoMessage()    {}
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f4653c102df3c6, []int{1}
}

func (m *PaymentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentResponse.Unmarshal(m, b)
}
func (m *PaymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentResponse.Marshal(b, m, deterministic)
}
func (m *PaymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentResponse.Merge(m, src)
}
func (m *PaymentResponse) XXX_Size() int {
	return xxx_messageInfo_PaymentResponse.Size(m)
}
func (m *PaymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentResponse proto.InternalMessageInfo

func (m *PaymentResponse) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *PaymentResponse) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

func (m *PaymentResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*PaymentRequest)(nil), "pay.PaymentRequest")
	proto.RegisterType((*PaymentResponse)(nil), "pay.PaymentResponse")
}

func init() { proto.RegisterFile("pay/proto/pay.proto", fileDescriptor_17f4653c102df3c6) }

var fileDescriptor_17f4653c102df3c6 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x89, 0x27, 0x07, 0x19, 0x35, 0xc2, 0x9e, 0x48, 0xb0, 0x0a, 0x07, 0xc2, 0x55, 0x77,
	0xa0, 0xa5, 0x9d, 0x9d, 0x82, 0x72, 0xac, 0x56, 0x36, 0x61, 0x4c, 0xa6, 0x48, 0x71, 0x3b, 0xeb,
	0xce, 0xac, 0xb0, 0xff, 0x5e, 0xb2, 0x86, 0xe0, 0x55, 0x56, 0xbb, 0xef, 0x3d, 0x78, 0x7c, 0x6f,
	0x60, 0xe5, 0x31, 0xed, 0x7c, 0x60, 0xe5, 0x9d, 0xc7, 0xb4, 0xcd, 0x3f, 0xb3, 0xf0, 0x98, 0xd6,
	0xcf, 0x50, 0xed, 0x31, 0x1d, 0xc8, 0xa9, 0xa5, 0xaf, 0x48, 0xa2, 0xe6, 0x1a, 0x96, 0x78, 0xe0,
	0xe8, 0xb4, 0x2e, 0x9a, 0x62, 0x73, 0x6a, 0x27, 0x65, 0x1a, 0x38, 0xe7, 0xa8, 0xad, 0x06, 0xec,
	0xa9, 0x75, 0x5c, 0x9f, 0x34, 0xc5, 0xa6, 0xb4, 0xc0, 0x51, 0xdf, 0x47, 0xeb, 0x95, 0xd7, 0x01,
	0x2e, 0xe7, 0x2e, 0xf1, 0xec, 0x84, 0xcc, 0x2d, 0x54, 0x1a, 0xd0, 0x09, 0x76, 0x3a, 0xb0, 0x6b,
	0x87, 0x3e, 0x97, 0x96, 0xf6, 0xe2, 0x8f, 0xfb, 0xd4, 0xff, 0xdf, 0x3d, 0x52, 0x89, 0xa2, 0x46,
	0xa9, 0x17, 0x39, 0x9b, 0xd4, 0xdd, 0xcb, 0xcc, 0xff, 0x46, 0xe1, 0x7b, 0xe8, 0xc8, 0x3c, 0x40,
	0xb5, 0x0f, 0xdc, 0x91, 0xc8, 0x14, 0x98, 0xd5, 0x76, 0x1c, 0x7d, 0x3c, 0xf3, 0xe6, 0xea, 0xd8,
	0xfc, 0xe5, 0x7d, 0x3c, 0xfb, 0x28, 0xe7, 0x53, 0x7d, 0x2e, 0xf3, 0x73, 0xff, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0x9f, 0xfa, 0xfa, 0x3e, 0x01, 0x00, 0x00,
}

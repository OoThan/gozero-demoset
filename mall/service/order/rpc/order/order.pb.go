// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order

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

// order create
type CreateRequest struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Pid                  uint64   `protobuf:"varint,2,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Amount               uint64   `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status               uint64   `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CreateRequest) GetPid() uint64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *CreateRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CreateRequest) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type CreateResponse struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// order update
type UpdateRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Uid                  uint64   `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Pid                  uint64   `protobuf:"varint,3,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Amount               uint64   `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status               uint64   `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateRequest) GetPid() uint64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *UpdateRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *UpdateRequest) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

// order remove
type RemoveRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{4}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RemoveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{5}
}

func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResponse.Unmarshal(m, b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
}
func (m *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(m, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveResponse.Size(m)
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

// order detail
type DetailRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailRequest) Reset()         { *m = DetailRequest{} }
func (m *DetailRequest) String() string { return proto.CompactTextString(m) }
func (*DetailRequest) ProtoMessage()    {}
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{6}
}

func (m *DetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailRequest.Unmarshal(m, b)
}
func (m *DetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailRequest.Marshal(b, m, deterministic)
}
func (m *DetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailRequest.Merge(m, src)
}
func (m *DetailRequest) XXX_Size() int {
	return xxx_messageInfo_DetailRequest.Size(m)
}
func (m *DetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailRequest proto.InternalMessageInfo

func (m *DetailRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DetailResponse struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Uid                  uint64   `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Pid                  uint64   `protobuf:"varint,3,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Amount               uint64   `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status               uint64   `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailResponse) Reset()         { *m = DetailResponse{} }
func (m *DetailResponse) String() string { return proto.CompactTextString(m) }
func (*DetailResponse) ProtoMessage()    {}
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{7}
}

func (m *DetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailResponse.Unmarshal(m, b)
}
func (m *DetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailResponse.Marshal(b, m, deterministic)
}
func (m *DetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailResponse.Merge(m, src)
}
func (m *DetailResponse) XXX_Size() int {
	return xxx_messageInfo_DetailResponse.Size(m)
}
func (m *DetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetailResponse proto.InternalMessageInfo

func (m *DetailResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DetailResponse) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *DetailResponse) GetPid() uint64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *DetailResponse) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DetailResponse) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

// order list
type ListRequest struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{8}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type ListResponse struct {
	Data                 []*DetailResponse `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{9}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetData() []*DetailResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

// order paid
type PaidRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaidRequest) Reset()         { *m = PaidRequest{} }
func (m *PaidRequest) String() string { return proto.CompactTextString(m) }
func (*PaidRequest) ProtoMessage()    {}
func (*PaidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{10}
}

func (m *PaidRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaidRequest.Unmarshal(m, b)
}
func (m *PaidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaidRequest.Marshal(b, m, deterministic)
}
func (m *PaidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaidRequest.Merge(m, src)
}
func (m *PaidRequest) XXX_Size() int {
	return xxx_messageInfo_PaidRequest.Size(m)
}
func (m *PaidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaidRequest proto.InternalMessageInfo

func (m *PaidRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PaidResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaidResponse) Reset()         { *m = PaidResponse{} }
func (m *PaidResponse) String() string { return proto.CompactTextString(m) }
func (*PaidResponse) ProtoMessage()    {}
func (*PaidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{11}
}

func (m *PaidResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaidResponse.Unmarshal(m, b)
}
func (m *PaidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaidResponse.Marshal(b, m, deterministic)
}
func (m *PaidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaidResponse.Merge(m, src)
}
func (m *PaidResponse) XXX_Size() int {
	return xxx_messageInfo_PaidResponse.Size(m)
}
func (m *PaidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PaidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PaidResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateRequest)(nil), "order.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "order.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "order.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "order.UpdateResponse")
	proto.RegisterType((*RemoveRequest)(nil), "order.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "order.RemoveResponse")
	proto.RegisterType((*DetailRequest)(nil), "order.DetailRequest")
	proto.RegisterType((*DetailResponse)(nil), "order.DetailResponse")
	proto.RegisterType((*ListRequest)(nil), "order.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "order.ListResponse")
	proto.RegisterType((*PaidRequest)(nil), "order.PaidRequest")
	proto.RegisterType((*PaidResponse)(nil), "order.PaidResponse")
}

func init() {
	proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077)
}

var fileDescriptor_cd01338c35d87077 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x31, 0x6f, 0xfa, 0x30,
	0x10, 0xc5, 0x45, 0x12, 0xf2, 0xd7, 0xff, 0x02, 0x11, 0x72, 0x4b, 0x15, 0x21, 0x55, 0xa0, 0x4c,
	0x74, 0x01, 0x89, 0x8a, 0xa1, 0x63, 0x5b, 0x16, 0xa4, 0x4a, 0x45, 0xa9, 0x58, 0xba, 0xa5, 0xb5,
	0x07, 0x4b, 0x05, 0xd3, 0xd8, 0xf4, 0x1b, 0xf6, 0x7b, 0x55, 0xce, 0xc5, 0x26, 0xa6, 0xc9, 0xd8,
	0x0d, 0xbf, 0xdc, 0xcf, 0x7e, 0x77, 0xef, 0x80, 0x48, 0x14, 0x94, 0x15, 0xb3, 0x43, 0x21, 0x94,
	0x20, 0xdd, 0xf2, 0x90, 0xbe, 0x43, 0xff, 0xb1, 0x60, 0xb9, 0x62, 0x19, 0xfb, 0x3c, 0x32, 0xa9,
	0xc8, 0x00, 0xfc, 0x2d, 0xa7, 0x49, 0x67, 0xd2, 0x99, 0x06, 0x99, 0xfe, 0xa9, 0x95, 0x0d, 0xa7,
	0x89, 0x87, 0xca, 0x86, 0x53, 0x72, 0x05, 0xe1, 0xfd, 0x4e, 0x1c, 0xf7, 0x2a, 0xf1, 0x4b, 0xb1,
	0x3a, 0x69, 0xfd, 0x45, 0xe5, 0xea, 0x28, 0x93, 0x00, 0x75, 0x3c, 0xa5, 0x13, 0x88, 0xcd, 0x23,
	0xf2, 0x20, 0xf6, 0x92, 0x91, 0x18, 0xbc, 0xb5, 0x79, 0xc4, 0x5b, 0xd3, 0x54, 0x42, 0x7f, 0x7b,
	0xa0, 0x35, 0x1b, 0x67, 0x05, 0xc6, 0x96, 0xf7, 0xcb, 0x96, 0xdf, 0x64, 0x2b, 0x68, 0xb1, 0xd5,
	0x75, 0x6c, 0x0d, 0x20, 0x36, 0x8f, 0xa2, 0xad, 0x74, 0x0c, 0xfd, 0x8c, 0xed, 0xc4, 0x57, 0x9b,
	0x0d, 0x8d, 0x98, 0x82, 0x13, 0xb2, 0x62, 0x2a, 0xe7, 0x1f, 0x6d, 0x88, 0x82, 0xd8, 0x14, 0x34,
	0x37, 0xff, 0x27, 0xbd, 0x8d, 0x21, 0x7a, 0xe2, 0x52, 0xb5, 0xa6, 0x9a, 0xde, 0x41, 0x0f, 0x0b,
	0x2a, 0x53, 0x37, 0x10, 0xac, 0x72, 0x95, 0x27, 0x9d, 0x89, 0x3f, 0x8d, 0x16, 0xc3, 0x19, 0xee,
	0x8a, 0xeb, 0x3c, 0x2b, 0x4b, 0xd2, 0x6b, 0x88, 0x36, 0x39, 0xa7, 0x6d, 0x0d, 0xc7, 0xd0, 0xc3,
	0xcf, 0x08, 0x2d, 0xbe, 0x3d, 0xe8, 0x3e, 0xeb, 0xdb, 0xc8, 0x12, 0x42, 0xdc, 0x03, 0x72, 0x59,
	0xdd, 0xef, 0xec, 0xde, 0x68, 0x78, 0xa6, 0x56, 0xd6, 0x96, 0x10, 0x62, 0x4e, 0x16, 0x73, 0x76,
	0xc5, 0x62, 0x6e, 0x98, 0x1a, 0xc3, 0xac, 0x2c, 0xe6, 0x64, 0x6b, 0x31, 0x37, 0x50, 0x8d, 0x61,
	0xd7, 0x16, 0x73, 0xf2, 0x1d, 0x35, 0x8f, 0x86, 0xcc, 0x21, 0xd0, 0xf3, 0x24, 0xa4, 0xfa, 0x5c,
	0x9b, 0xfe, 0xe8, 0xc2, 0xd1, 0x4e, 0x80, 0x1e, 0x93, 0x05, 0x6a, 0x23, 0xb5, 0x40, 0x7d, 0x8e,
	0x0f, 0xff, 0x5f, 0xff, 0xcd, 0xe6, 0xa5, 0xfe, 0x16, 0x96, 0xff, 0xe1, 0xdb, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x83, 0x67, 0x05, 0x06, 0xd2, 0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment.proto

package consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_514ebfffef9f6c62, []int{0}
}
func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (dst *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(dst, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_514ebfffef9f6c62, []int{1}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_514ebfffef9f6c62, []int{2}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_514ebfffef9f6c62, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "consignment.Consignment")
	proto.RegisterType((*Container)(nil), "consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "consignment.Response")
}

func init() {
	proto.RegisterFile("proto/consignment.proto", fileDescriptor_consignment_514ebfffef9f6c62)
}

var fileDescriptor_consignment_514ebfffef9f6c62 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x6e, 0xe2, 0x30,
	0x14, 0x85, 0x27, 0xfc, 0xe7, 0x06, 0x0d, 0x1a, 0x4b, 0x03, 0xd6, 0xcc, 0xa2, 0x51, 0x56, 0xac,
	0xa8, 0x44, 0xa5, 0x2e, 0xaa, 0xee, 0x22, 0x15, 0xb1, 0x35, 0x0f, 0x50, 0xd1, 0xf8, 0x2a, 0x5c,
	0xa9, 0xd8, 0xa9, 0x6d, 0xe8, 0xdb, 0xf0, 0x04, 0x7d, 0xc8, 0x0a, 0x87, 0x14, 0xd3, 0x8a, 0x5d,
	0xce, 0xb9, 0xe7, 0x38, 0x9f, 0x6f, 0x02, 0x93, 0xca, 0x68, 0xa7, 0x6f, 0x0b, 0xad, 0x2c, 0x95,
	0x6a, 0x8b, 0xca, 0xcd, 0xbc, 0xc3, 0x92, 0xc0, 0xca, 0x3e, 0x22, 0x48, 0xf2, 0xb3, 0x66, 0xbf,
	0xa1, 0x45, 0x92, 0x47, 0x69, 0x34, 0x8d, 0x45, 0x8b, 0x24, 0x4b, 0x21, 0x91, 0x68, 0x0b, 0x43,
	0x95, 0x23, 0xad, 0x78, 0xcb, 0x0f, 0x42, 0x8b, 0x8d, 0xa1, 0xf7, 0x8e, 0x54, 0x6e, 0x1c, 0x6f,
	0xa7, 0xd1, 0xb4, 0x2b, 0x4e, 0x8a, 0xdd, 0x03, 0x14, 0x5a, 0xb9, 0x35, 0x29, 0x34, 0x96, 0x77,
	0xd2, 0xf6, 0x34, 0x99, 0x8f, 0x67, 0x21, 0x4e, 0xde, 0x8c, 0x45, 0x90, 0x64, 0xff, 0x21, 0xde,
	0xa3, 0xb5, 0xf8, 0xfa, 0x4c, 0x92, 0x77, 0xfd, 0xfb, 0x06, 0xb5, 0xb1, 0x94, 0xd9, 0x16, 0xe2,
	0xaf, 0xd6, 0x0f, 0xd6, 0x1b, 0x48, 0x8a, 0x9d, 0x75, 0x7a, 0x8b, 0xe6, 0xd8, 0xad, 0x59, 0xa1,
	0xb1, 0x96, 0xf2, 0x88, 0xaa, 0x0d, 0x95, 0xa4, 0x3c, 0x6a, 0x2c, 0x4e, 0x8a, 0x4d, 0xa0, 0xbf,
	0xb3, 0x75, 0xa9, 0x53, 0x0f, 0x8e, 0x72, 0x29, 0xb3, 0x21, 0xc0, 0x02, 0x9d, 0xc0, 0xb7, 0x1d,
	0x5a, 0x97, 0x1d, 0x22, 0x18, 0x08, 0xb4, 0x95, 0x56, 0x16, 0x19, 0x87, 0x7e, 0x61, 0x70, 0xed,
	0xb0, 0x26, 0x18, 0x88, 0x46, 0xb2, 0x07, 0x08, 0x37, 0xec, 0x31, 0x92, 0x39, 0xff, 0x7e, 0xf3,
	0xe6, 0x59, 0x84, 0x61, 0xf6, 0x08, 0xc3, 0x40, 0x5a, 0xde, 0xf6, 0x6b, 0xbb, 0x5e, 0xbe, 0x48,
	0xcf, 0x0f, 0x11, 0x8c, 0x56, 0x1b, 0xaa, 0x2a, 0x52, 0xe5, 0x0a, 0xcd, 0x9e, 0x0a, 0x64, 0x4f,
	0xf0, 0x27, 0xf7, 0x60, 0xe1, 0x57, 0xbe, 0x7a, 0xe0, 0xbf, 0xbf, 0x17, 0x93, 0xe6, 0xb6, 0xd9,
	0x2f, 0x96, 0xc3, 0x68, 0x81, 0x2e, 0x88, 0x5a, 0x36, 0xb9, 0xc8, 0x9e, 0x17, 0x75, 0xf5, 0x90,
	0x97, 0x9e, 0xff, 0x03, 0xef, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x8b, 0x83, 0x3e, 0x9c,
	0x02, 0x00, 0x00,
}

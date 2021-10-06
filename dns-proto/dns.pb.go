// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dns-proto/dns.proto

package dnsProto

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

type NewDomainAddress struct {
	Domain               string   `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewDomainAddress) Reset()         { *m = NewDomainAddress{} }
func (m *NewDomainAddress) String() string { return proto.CompactTextString(m) }
func (*NewDomainAddress) ProtoMessage()    {}
func (*NewDomainAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_d460590a80dac11d, []int{0}
}

func (m *NewDomainAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewDomainAddress.Unmarshal(m, b)
}
func (m *NewDomainAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewDomainAddress.Marshal(b, m, deterministic)
}
func (m *NewDomainAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewDomainAddress.Merge(m, src)
}
func (m *NewDomainAddress) XXX_Size() int {
	return xxx_messageInfo_NewDomainAddress.Size(m)
}
func (m *NewDomainAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NewDomainAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NewDomainAddress proto.InternalMessageInfo

func (m *NewDomainAddress) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *NewDomainAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Address struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_d460590a80dac11d, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Domain struct {
	Domain               string   `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Domain) Reset()         { *m = Domain{} }
func (m *Domain) String() string { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()    {}
func (*Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_d460590a80dac11d, []int{2}
}

func (m *Domain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Domain.Unmarshal(m, b)
}
func (m *Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Domain.Marshal(b, m, deterministic)
}
func (m *Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain.Merge(m, src)
}
func (m *Domain) XXX_Size() int {
	return xxx_messageInfo_Domain.Size(m)
}
func (m *Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Domain proto.InternalMessageInfo

func (m *Domain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func init() {
	proto.RegisterType((*NewDomainAddress)(nil), "dnsProto.NewDomainAddress")
	proto.RegisterType((*Address)(nil), "dnsProto.Address")
	proto.RegisterType((*Domain)(nil), "dnsProto.Domain")
}

func init() { proto.RegisterFile("dns-proto/dns.proto", fileDescriptor_d460590a80dac11d) }

var fileDescriptor_d460590a80dac11d = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xc9, 0x2b, 0xd6,
	0x2d, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xc9, 0x2b, 0xd6, 0x03, 0xb3, 0x84, 0x38, 0x52, 0xf2,
	0x8a, 0x03, 0x40, 0x2c, 0x25, 0x17, 0x2e, 0x01, 0xbf, 0xd4, 0x72, 0x97, 0xfc, 0xdc, 0xc4, 0xcc,
	0x3c, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x21, 0x31, 0x2e, 0xb6, 0x14, 0xb0, 0x80, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x94, 0x27, 0x24, 0xc1, 0xc5, 0x9e, 0x08, 0x51, 0x22, 0xc1,
	0x04, 0x96, 0x80, 0x71, 0x95, 0x94, 0xb9, 0xd8, 0x61, 0x9a, 0x91, 0x14, 0x31, 0xa2, 0x2a, 0x52,
	0xe0, 0x62, 0x83, 0xd8, 0x83, 0xcb, 0x02, 0xa3, 0x5e, 0x46, 0x2e, 0x2e, 0x97, 0xbc, 0xe2, 0xe0,
	0x92, 0xfc, 0xa2, 0xc4, 0xf4, 0x54, 0x21, 0x17, 0x2e, 0xc1, 0xe0, 0xc4, 0xb2, 0x54, 0x54, 0xc7,
	0x49, 0xe9, 0xc1, 0xdc, 0xae, 0x87, 0xee, 0x70, 0x29, 0x41, 0x84, 0x1c, 0x54, 0x48, 0x89, 0x41,
	0xc8, 0x92, 0x4b, 0xc0, 0x3d, 0xb5, 0x04, 0xd5, 0x10, 0x01, 0x84, 0x42, 0x88, 0x04, 0x56, 0xad,
	0x4e, 0x3a, 0x51, 0x5a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x89,
	0x39, 0x39, 0xa9, 0x15, 0x95, 0xc5, 0x39, 0xa9, 0xa9, 0x05, 0xc5, 0xa0, 0xb0, 0xd4, 0x07, 0x87,
	0xa5, 0x35, 0x4c, 0x5f, 0x12, 0x1b, 0x98, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x11, 0xc5,
	0x86, 0x41, 0x72, 0x01, 0x00, 0x00,
}

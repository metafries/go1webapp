// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Club.proto

package protobuff

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

type Club struct {
	Clubname             string           `protobuf:"bytes,1,opt,name=clubname,proto3" json:"clubname,omitempty"`
	LeagueName           string           `protobuf:"bytes,2,opt,name=LeagueName,proto3" json:"LeagueName,omitempty"`
	CI                   []*Club_ClubInfo `protobuf:"bytes,3,rep,name=CI,proto3" json:"CI,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Club) Reset()         { *m = Club{} }
func (m *Club) String() string { return proto.CompactTextString(m) }
func (*Club) ProtoMessage()    {}
func (*Club) Descriptor() ([]byte, []int) {
	return fileDescriptor_88c26d7e3588db28, []int{0}
}
func (m *Club) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Club.Unmarshal(m, b)
}
func (m *Club) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Club.Marshal(b, m, deterministic)
}
func (m *Club) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Club.Merge(m, src)
}
func (m *Club) XXX_Size() int {
	return xxx_messageInfo_Club.Size(m)
}
func (m *Club) XXX_DiscardUnknown() {
	xxx_messageInfo_Club.DiscardUnknown(m)
}

var xxx_messageInfo_Club proto.InternalMessageInfo

func (m *Club) GetClubname() string {
	if m != nil {
		return m.Clubname
	}
	return ""
}

func (m *Club) GetLeagueName() string {
	if m != nil {
		return m.LeagueName
	}
	return ""
}

func (m *Club) GetCI() []*Club_ClubInfo {
	if m != nil {
		return m.CI
	}
	return nil
}

type Club_ClubInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ranking              string   `protobuf:"bytes,3,opt,name=ranking,proto3" json:"ranking,omitempty"`
	League               string   `protobuf:"bytes,4,opt,name=league,proto3" json:"league,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Club_ClubInfo) Reset()         { *m = Club_ClubInfo{} }
func (m *Club_ClubInfo) String() string { return proto.CompactTextString(m) }
func (*Club_ClubInfo) ProtoMessage()    {}
func (*Club_ClubInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_88c26d7e3588db28, []int{0, 0}
}
func (m *Club_ClubInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Club_ClubInfo.Unmarshal(m, b)
}
func (m *Club_ClubInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Club_ClubInfo.Marshal(b, m, deterministic)
}
func (m *Club_ClubInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Club_ClubInfo.Merge(m, src)
}
func (m *Club_ClubInfo) XXX_Size() int {
	return xxx_messageInfo_Club_ClubInfo.Size(m)
}
func (m *Club_ClubInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Club_ClubInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Club_ClubInfo proto.InternalMessageInfo

func (m *Club_ClubInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Club_ClubInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Club_ClubInfo) GetRanking() string {
	if m != nil {
		return m.Ranking
	}
	return ""
}

func (m *Club_ClubInfo) GetLeague() string {
	if m != nil {
		return m.League
	}
	return ""
}

func init() {
	proto.RegisterType((*Club)(nil), "protobuff.Club")
	proto.RegisterType((*Club_ClubInfo)(nil), "protobuff.Club.ClubInfo")
}

func init() { proto.RegisterFile("Club.proto", fileDescriptor_88c26d7e3588db28) }

var fileDescriptor_88c26d7e3588db28 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x72, 0xce, 0x29, 0x4d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x53, 0x49, 0xa5, 0x69, 0x69, 0x4a, 0xe7,
	0x18, 0xb9, 0x58, 0x40, 0x32, 0x42, 0x52, 0x5c, 0x1c, 0xc9, 0x39, 0xa5, 0x49, 0x79, 0x89, 0xb9,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x1c, 0x17, 0x97, 0x4f, 0x6a,
	0x62, 0x7a, 0x69, 0xaa, 0x1f, 0x48, 0x96, 0x09, 0x2c, 0x8b, 0x24, 0x22, 0xa4, 0xc1, 0xc5, 0xe4,
	0xec, 0x29, 0xc1, 0xac, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xa1, 0x07, 0x37, 0x5c, 0x0f, 0x6c, 0x25,
	0x88, 0xf0, 0xcc, 0x4b, 0xcb, 0x0f, 0x62, 0x72, 0xf6, 0x94, 0x4a, 0xe0, 0xe2, 0x80, 0xf1, 0x85,
	0xf8, 0xb8, 0x98, 0x32, 0x53, 0xc0, 0x76, 0xb1, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1,
	0xe4, 0x21, 0xcc, 0x07, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0x8b, 0x12, 0xf3, 0xb2, 0x33, 0xf3, 0xd2,
	0x25, 0x98, 0xc1, 0xc2, 0x30, 0xae, 0x90, 0x18, 0x17, 0x5b, 0x0e, 0xd8, 0x05, 0x12, 0x2c, 0x60,
	0x09, 0x28, 0x2f, 0x89, 0x0d, 0x6c, 0xbd, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x09, 0x7e, 0x3e,
	0x3a, 0xf0, 0x00, 0x00, 0x00,
}

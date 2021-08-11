// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/gov/actor.proto

package gov

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type ActorStatus int32

const (
	// Undefined status
	ActorStatus_UNDEFINED ActorStatus = 0
	// Unclaimed status
	ActorStatus_UNCLAIMED ActorStatus = 1
	// Active status
	ActorStatus_ACTIVE ActorStatus = 2
	// Paused status
	ActorStatus_PAUSED ActorStatus = 3
	// Inactive status
	ActorStatus_INACTIVE ActorStatus = 4
	// Jailed status
	ActorStatus_JAILED ActorStatus = 5
	// Removed status
	ActorStatus_REMOVED ActorStatus = 6
)

var ActorStatus_name = map[int32]string{
	0: "UNDEFINED",
	1: "UNCLAIMED",
	2: "ACTIVE",
	3: "PAUSED",
	4: "INACTIVE",
	5: "JAILED",
	6: "REMOVED",
}

var ActorStatus_value = map[string]int32{
	"UNDEFINED": 0,
	"UNCLAIMED": 1,
	"ACTIVE":    2,
	"PAUSED":    3,
	"INACTIVE":  4,
	"JAILED":    5,
	"REMOVED":   6,
}

func (x ActorStatus) String() string {
	return proto.EnumName(ActorStatus_name, int32(x))
}

func (ActorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{0}
}

type Permissions struct {
	Blacklist            []uint32 `protobuf:"varint,1,rep,packed,name=blacklist,proto3" json:"blacklist,omitempty"`
	Whitelist            []uint32 `protobuf:"varint,2,rep,packed,name=whitelist,proto3" json:"whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{0}
}

func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (m *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(m, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetBlacklist() []uint32 {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *Permissions) GetWhitelist() []uint32 {
	if m != nil {
		return m.Whitelist
	}
	return nil
}

type NetworkActor struct {
	Address              []byte       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Roles                []uint64     `protobuf:"varint,2,rep,packed,name=roles,proto3" json:"roles,omitempty"`
	Status               ActorStatus  `protobuf:"varint,3,opt,name=status,proto3,enum=kira.gov.ActorStatus" json:"status,omitempty"`
	Votes                []VoteOption `protobuf:"varint,4,rep,packed,name=votes,proto3,enum=kira.gov.VoteOption" json:"votes,omitempty"`
	Permissions          *Permissions `protobuf:"bytes,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Skin                 uint64       `protobuf:"varint,6,opt,name=skin,proto3" json:"skin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NetworkActor) Reset()         { *m = NetworkActor{} }
func (m *NetworkActor) String() string { return proto.CompactTextString(m) }
func (*NetworkActor) ProtoMessage()    {}
func (*NetworkActor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{1}
}

func (m *NetworkActor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkActor.Unmarshal(m, b)
}
func (m *NetworkActor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkActor.Marshal(b, m, deterministic)
}
func (m *NetworkActor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkActor.Merge(m, src)
}
func (m *NetworkActor) XXX_Size() int {
	return xxx_messageInfo_NetworkActor.Size(m)
}
func (m *NetworkActor) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkActor.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkActor proto.InternalMessageInfo

func (m *NetworkActor) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *NetworkActor) GetRoles() []uint64 {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *NetworkActor) GetStatus() ActorStatus {
	if m != nil {
		return m.Status
	}
	return ActorStatus_UNDEFINED
}

func (m *NetworkActor) GetVotes() []VoteOption {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *NetworkActor) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *NetworkActor) GetSkin() uint64 {
	if m != nil {
		return m.Skin
	}
	return 0
}

type MsgWhitelistPermissions struct {
	Proposer             []byte   `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Permission           uint32   `protobuf:"varint,3,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgWhitelistPermissions) Reset()         { *m = MsgWhitelistPermissions{} }
func (m *MsgWhitelistPermissions) String() string { return proto.CompactTextString(m) }
func (*MsgWhitelistPermissions) ProtoMessage()    {}
func (*MsgWhitelistPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{2}
}

func (m *MsgWhitelistPermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWhitelistPermissions.Unmarshal(m, b)
}
func (m *MsgWhitelistPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWhitelistPermissions.Marshal(b, m, deterministic)
}
func (m *MsgWhitelistPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWhitelistPermissions.Merge(m, src)
}
func (m *MsgWhitelistPermissions) XXX_Size() int {
	return xxx_messageInfo_MsgWhitelistPermissions.Size(m)
}
func (m *MsgWhitelistPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWhitelistPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWhitelistPermissions proto.InternalMessageInfo

func (m *MsgWhitelistPermissions) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *MsgWhitelistPermissions) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgWhitelistPermissions) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

type MsgBlacklistPermissions struct {
	Proposer             []byte   `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Permission           uint32   `protobuf:"varint,3,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBlacklistPermissions) Reset()         { *m = MsgBlacklistPermissions{} }
func (m *MsgBlacklistPermissions) String() string { return proto.CompactTextString(m) }
func (*MsgBlacklistPermissions) ProtoMessage()    {}
func (*MsgBlacklistPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{3}
}

func (m *MsgBlacklistPermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBlacklistPermissions.Unmarshal(m, b)
}
func (m *MsgBlacklistPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBlacklistPermissions.Marshal(b, m, deterministic)
}
func (m *MsgBlacklistPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBlacklistPermissions.Merge(m, src)
}
func (m *MsgBlacklistPermissions) XXX_Size() int {
	return xxx_messageInfo_MsgBlacklistPermissions.Size(m)
}
func (m *MsgBlacklistPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBlacklistPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBlacklistPermissions proto.InternalMessageInfo

func (m *MsgBlacklistPermissions) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *MsgBlacklistPermissions) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgBlacklistPermissions) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

func init() {
	proto.RegisterEnum("kira.gov.ActorStatus", ActorStatus_name, ActorStatus_value)
	proto.RegisterType((*Permissions)(nil), "kira.gov.Permissions")
	proto.RegisterType((*NetworkActor)(nil), "kira.gov.NetworkActor")
	proto.RegisterType((*MsgWhitelistPermissions)(nil), "kira.gov.MsgWhitelistPermissions")
	proto.RegisterType((*MsgBlacklistPermissions)(nil), "kira.gov.MsgBlacklistPermissions")
}

func init() {
	proto.RegisterFile("kira/gov/actor.proto", fileDescriptor_c4584595efc3386e)
}

var fileDescriptor_c4584595efc3386e = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0x8d, 0x6c, 0xd9, 0x71, 0xc6, 0x49, 0x10, 0x83, 0xbf, 0x2f, 0x42, 0x94, 0x44, 0x68, 0x25,
	0x02, 0x96, 0x68, 0x42, 0x29, 0x64, 0xa7, 0x58, 0x2a, 0x28, 0x8d, 0x9d, 0xa0, 0xfc, 0xb4, 0x14,
	0xba, 0x98, 0x48, 0x53, 0x65, 0xd0, 0xcf, 0x08, 0xcd, 0xd8, 0x21, 0x6f, 0x50, 0xfc, 0x0e, 0x86,
	0x42, 0x97, 0x7d, 0x96, 0xbe, 0x42, 0xb6, 0xed, 0xa6, 0x9b, 0x2e, 0xbb, 0x2a, 0x92, 0x2c, 0x5b,
	0xdb, 0x2e, 0xba, 0xe9, 0x4a, 0xa3, 0x73, 0xee, 0xb9, 0xdc, 0x7b, 0xee, 0xcc, 0x05, 0x83, 0x88,
	0xe4, 0xc8, 0x0c, 0xe9, 0xcc, 0x44, 0x3e, 0xa7, 0xb9, 0x91, 0xe5, 0x94, 0x53, 0xd8, 0x2b, 0x50,
	0x23, 0xa4, 0x33, 0x65, 0x10, 0xd2, 0x90, 0x96, 0xa0, 0x59, 0x9c, 0x2a, 0x5e, 0xd9, 0x5b, 0xa9,
	0xb2, 0x9c, 0x66, 0x94, 0xa1, 0xb8, 0x22, 0x34, 0x17, 0xf4, 0x2f, 0x71, 0x9e, 0x10, 0xc6, 0x08,
	0x4d, 0x19, 0x7c, 0x06, 0xb6, 0xee, 0x62, 0xe4, 0x47, 0x31, 0x61, 0x5c, 0x16, 0xd4, 0xb6, 0xbe,
	0xe3, 0xad, 0x81, 0x82, 0x7d, 0xb8, 0x27, 0x1c, 0x97, 0x6c, 0xab, 0x62, 0x57, 0x80, 0xf6, 0xa5,
	0x05, 0xb6, 0x27, 0x98, 0x3f, 0xd0, 0x3c, 0xb2, 0x8a, 0xd2, 0xe0, 0x7b, 0xb0, 0x89, 0x82, 0x20,
	0xc7, 0x8c, 0xc9, 0x82, 0x2a, 0xe8, 0xdb, 0xa7, 0xa3, 0x9f, 0x4f, 0x07, 0xbb, 0x8f, 0x28, 0x89,
	0x4f, 0xb4, 0x25, 0xa1, 0xfd, 0x7a, 0x3a, 0x18, 0x86, 0x84, 0xdf, 0x4f, 0xef, 0x0c, 0x9f, 0x26,
	0xa6, 0x4f, 0x59, 0x42, 0xd9, 0xf2, 0x33, 0x64, 0x41, 0x64, 0xf2, 0xc7, 0x0c, 0x33, 0xc3, 0xf2,
	0x7d, 0xab, 0x52, 0x78, 0x75, 0x4e, 0x38, 0x00, 0x9d, 0x9c, 0xc6, 0x98, 0x95, 0x95, 0x88, 0x5e,
	0xf5, 0x03, 0x87, 0xa0, 0xcb, 0x38, 0xe2, 0x53, 0x26, 0xb7, 0x55, 0x41, 0xdf, 0x3d, 0xfa, 0xcf,
	0xa8, 0xad, 0x31, 0xca, 0xaa, 0xae, 0x4a, 0xd2, 0x5b, 0x06, 0xc1, 0x43, 0xd0, 0x99, 0x51, 0x8e,
	0x99, 0x2c, 0xaa, 0x6d, 0x7d, 0xf7, 0x68, 0xb0, 0x8e, 0xbe, 0xa5, 0x1c, 0x5f, 0x64, 0x9c, 0xd0,
	0xd4, 0xab, 0x42, 0xe0, 0x4b, 0xd0, 0xcf, 0xd6, 0x5e, 0xc9, 0x1d, 0x55, 0xd0, 0xfb, 0xcd, 0xfc,
	0x0d, 0x23, 0xbd, 0x66, 0x24, 0x84, 0x40, 0x64, 0x11, 0x49, 0xe5, 0xae, 0x2a, 0xe8, 0xa2, 0x57,
	0x9e, 0xb5, 0x1f, 0x02, 0xd8, 0x1b, 0xb3, 0xf0, 0x4d, 0x6d, 0x5f, 0x73, 0x0a, 0x63, 0xd0, 0xab,
	0xc6, 0x84, 0xf3, 0xa5, 0x73, 0xcf, 0xff, 0xdc, 0xa7, 0x55, 0x8a, 0xe6, 0x1c, 0x5a, 0x7f, 0x61,
	0x0e, 0xfb, 0x00, 0xac, 0x9b, 0x2d, 0x5d, 0xdf, 0xf1, 0x1a, 0xc8, 0x89, 0xf8, 0xfd, 0xd3, 0x81,
	0xa0, 0x7d, 0xab, 0xfa, 0x3d, 0xad, 0x2f, 0xd3, 0x3f, 0xdb, 0xef, 0xe1, 0x57, 0x01, 0xf4, 0x1b,
	0x57, 0xad, 0x78, 0x35, 0x37, 0x13, 0xdb, 0x79, 0xe5, 0x4e, 0x1c, 0x5b, 0xda, 0x50, 0x76, 0xe6,
	0x0b, 0x75, 0xeb, 0x26, 0x0d, 0xf0, 0x07, 0x92, 0xe2, 0xa0, 0x62, 0x47, 0xe7, 0x96, 0x3b, 0x76,
	0x6c, 0x49, 0xa8, 0x59, 0x3f, 0x46, 0x24, 0xc1, 0x01, 0xfc, 0x1f, 0x74, 0xad, 0xd1, 0xb5, 0x7b,
	0xeb, 0x48, 0x2d, 0x05, 0xcc, 0x17, 0x6a, 0xd7, 0xf2, 0x39, 0x99, 0xe1, 0x02, 0xbf, 0xb4, 0x6e,
	0xae, 0x1c, 0x5b, 0x6a, 0x57, 0xf8, 0x25, 0x9a, 0x32, 0x1c, 0x40, 0x05, 0xf4, 0xdc, 0xc9, 0x52,
	0x21, 0x2a, 0xdb, 0xf3, 0x85, 0xda, 0x73, 0x53, 0xb4, 0xd2, 0x9c, 0x59, 0xee, 0xb9, 0x63, 0x4b,
	0x9d, 0x4a, 0x73, 0x86, 0x48, 0x8c, 0x03, 0x28, 0x83, 0x4d, 0xcf, 0x19, 0x5f, 0xdc, 0x3a, 0xb6,
	0xd4, 0x55, 0xfa, 0xf3, 0x85, 0xba, 0xe9, 0xe1, 0x84, 0xce, 0x70, 0xa0, 0x88, 0x1f, 0x3f, 0xef,
	0x6f, 0x9c, 0xbe, 0x78, 0x77, 0xdc, 0x70, 0xea, 0x35, 0xc9, 0xd1, 0x88, 0xe6, 0xd8, 0x64, 0x38,
	0x42, 0xc4, 0x74, 0x27, 0xd7, 0x8e, 0xf7, 0xd6, 0x2c, 0x97, 0xc9, 0x30, 0xc4, 0xa9, 0x59, 0x2f,
	0x9a, 0xbb, 0x6e, 0x89, 0x1d, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x65, 0xd1, 0x8b, 0x32, 0xb1,
	0x04, 0x00, 0x00,
}

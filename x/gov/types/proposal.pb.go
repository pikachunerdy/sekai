// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proposal.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// VoteOption enumerates the valid vote options for a given governance proposal.
type VoteOption int32

const (
	// VOTE_OPTION_UNSPECIFIED defines a no-op vote option.
	OptionEmpty VoteOption = 0
	// VOTE_OPTION_YES defines a yes vote option.
	OptionYes VoteOption = 1
	// VOTE_OPTION_ABSTAIN defines an abstain vote option.
	OptionAbstain VoteOption = 2
	// VOTE_OPTION_NO defines a no vote option.
	OptionNo VoteOption = 3
	// VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
	OptionNoWithVeto VoteOption = 4
)

var VoteOption_name = map[int32]string{
	0: "VOTE_OPTION_UNSPECIFIED",
	1: "VOTE_OPTION_YES",
	2: "VOTE_OPTION_ABSTAIN",
	3: "VOTE_OPTION_NO",
	4: "VOTE_OPTION_NO_WITH_VETO",
}

var VoteOption_value = map[string]int32{
	"VOTE_OPTION_UNSPECIFIED":  0,
	"VOTE_OPTION_YES":          1,
	"VOTE_OPTION_ABSTAIN":      2,
	"VOTE_OPTION_NO":           3,
	"VOTE_OPTION_NO_WITH_VETO": 4,
}

func (x VoteOption) String() string {
	return proto.EnumName(VoteOption_name, int32(x))
}

func (VoteOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{0}
}

type Vote struct {
	ProposalId uint64                                        `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Voter      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=voter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"voter,omitempty" yaml:"address"`
	Option     VoteOption                                    `protobuf:"varint,3,opt,name=option,proto3,enum=kira.gov.VoteOption" json:"option,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{0}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *Vote) GetVoter() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *Vote) GetOption() VoteOption {
	if m != nil {
		return m.Option
	}
	return OptionEmpty
}

type MsgVoteProposal struct {
	ProposalId uint64                                        `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Voter      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=voter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"voter,omitempty" yaml:"address"`
	Option     VoteOption                                    `protobuf:"varint,3,opt,name=option,proto3,enum=kira.gov.VoteOption" json:"option,omitempty"`
}

func (m *MsgVoteProposal) Reset()         { *m = MsgVoteProposal{} }
func (m *MsgVoteProposal) String() string { return proto.CompactTextString(m) }
func (*MsgVoteProposal) ProtoMessage()    {}
func (*MsgVoteProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{1}
}
func (m *MsgVoteProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVoteProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVoteProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVoteProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVoteProposal.Merge(m, src)
}
func (m *MsgVoteProposal) XXX_Size() int {
	return m.Size()
}
func (m *MsgVoteProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVoteProposal.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVoteProposal proto.InternalMessageInfo

func (m *MsgVoteProposal) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *MsgVoteProposal) GetVoter() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *MsgVoteProposal) GetOption() VoteOption {
	if m != nil {
		return m.Option
	}
	return OptionEmpty
}

type MsgProposalAssignPermission struct {
	Proposer   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty" yaml:"address"`
	Address    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" yaml:"address"`
	Permission uint32                                        `protobuf:"varint,3,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (m *MsgProposalAssignPermission) Reset()         { *m = MsgProposalAssignPermission{} }
func (m *MsgProposalAssignPermission) String() string { return proto.CompactTextString(m) }
func (*MsgProposalAssignPermission) ProtoMessage()    {}
func (*MsgProposalAssignPermission) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{2}
}
func (m *MsgProposalAssignPermission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposalAssignPermission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposalAssignPermission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposalAssignPermission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposalAssignPermission.Merge(m, src)
}
func (m *MsgProposalAssignPermission) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposalAssignPermission) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposalAssignPermission.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposalAssignPermission proto.InternalMessageInfo

func (m *MsgProposalAssignPermission) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *MsgProposalAssignPermission) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgProposalAssignPermission) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

type ProposalAssignPermission struct {
	Address    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" yaml:"address"`
	Permission uint32                                        `protobuf:"varint,2,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (m *ProposalAssignPermission) Reset()         { *m = ProposalAssignPermission{} }
func (m *ProposalAssignPermission) String() string { return proto.CompactTextString(m) }
func (*ProposalAssignPermission) ProtoMessage()    {}
func (*ProposalAssignPermission) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{3}
}
func (m *ProposalAssignPermission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalAssignPermission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalAssignPermission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalAssignPermission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalAssignPermission.Merge(m, src)
}
func (m *ProposalAssignPermission) XXX_Size() int {
	return m.Size()
}
func (m *ProposalAssignPermission) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalAssignPermission.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalAssignPermission proto.InternalMessageInfo

func (m *ProposalAssignPermission) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ProposalAssignPermission) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

func init() {
	proto.RegisterEnum("kira.gov.VoteOption", VoteOption_name, VoteOption_value)
	proto.RegisterType((*Vote)(nil), "kira.gov.Vote")
	proto.RegisterType((*MsgVoteProposal)(nil), "kira.gov.MsgVoteProposal")
	proto.RegisterType((*MsgProposalAssignPermission)(nil), "kira.gov.MsgProposalAssignPermission")
	proto.RegisterType((*ProposalAssignPermission)(nil), "kira.gov.ProposalAssignPermission")
}

func init() { proto.RegisterFile("proposal.proto", fileDescriptor_c3ac5ce23bf32d05) }

var fileDescriptor_c3ac5ce23bf32d05 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0xdd, 0xba, 0xd6, 0xb7, 0xdb, 0x36, 0x8e, 0x05, 0x4b, 0x84, 0x34, 0x04, 0x84,
	0xb2, 0xec, 0x26, 0xb0, 0xde, 0xbc, 0x48, 0x5a, 0x23, 0x06, 0xd9, 0xa6, 0xb4, 0xb5, 0x4b, 0x05,
	0x09, 0x69, 0x3b, 0x64, 0x87, 0x6e, 0x3b, 0x21, 0x33, 0x16, 0xfb, 0x0d, 0xa4, 0x27, 0xaf, 0x1e,
	0x7a, 0xf2, 0x0b, 0x78, 0xf7, 0x0b, 0x78, 0xdc, 0xa3, 0xa7, 0x45, 0x5a, 0xf0, 0x03, 0x78, 0x14,
	0x04, 0x69, 0xd2, 0xac, 0x51, 0xf0, 0xb6, 0x1e, 0x3c, 0x25, 0xbc, 0xf7, 0x7b, 0x8f, 0xdf, 0xfb,
	0x07, 0x02, 0xc5, 0x20, 0x64, 0x01, 0xe3, 0xde, 0xb9, 0x1e, 0x84, 0x4c, 0x30, 0x9c, 0x1f, 0xd3,
	0xd0, 0xd3, 0x7d, 0x36, 0x93, 0xcb, 0x3e, 0xf3, 0x59, 0x54, 0x34, 0x36, 0x6f, 0x71, 0x5f, 0xfb,
	0x80, 0x20, 0xd7, 0x63, 0x82, 0xe0, 0x2a, 0xec, 0x25, 0xa3, 0x2e, 0x1d, 0x55, 0x90, 0x8a, 0x6a,
	0xb9, 0x36, 0x24, 0x25, 0x7b, 0x84, 0xfb, 0x70, 0x63, 0xc6, 0x04, 0x09, 0x2b, 0x59, 0x15, 0xd5,
	0xf6, 0xeb, 0x8d, 0x6f, 0x97, 0xd5, 0xe2, 0xdc, 0x9b, 0x9c, 0x3f, 0xd4, 0xbc, 0xd1, 0x28, 0x24,
	0x9c, 0x6b, 0xdf, 0x2f, 0xab, 0x47, 0x3e, 0x15, 0x67, 0xaf, 0x06, 0xfa, 0x90, 0x4d, 0x8c, 0x21,
	0xe3, 0x13, 0xc6, 0xb7, 0x8f, 0x23, 0x3e, 0x1a, 0x1b, 0x62, 0x1e, 0x10, 0xae, 0x9b, 0xc3, 0xa1,
	0x19, 0x4f, 0xb4, 0xe3, 0x8d, 0xf8, 0x10, 0x76, 0x59, 0x20, 0x28, 0x9b, 0x56, 0x76, 0x54, 0x54,
	0x2b, 0x1e, 0x97, 0xf5, 0xc4, 0x5a, 0xdf, 0xb8, 0x39, 0x51, 0xaf, 0xbd, 0x65, 0xb4, 0x8f, 0x08,
	0x4a, 0x27, 0xdc, 0xdf, 0x74, 0x5a, 0x5b, 0xbd, 0xff, 0xc8, 0xfe, 0x07, 0x82, 0x7b, 0x27, 0xdc,
	0x4f, 0xcc, 0x4d, 0xce, 0xa9, 0x3f, 0x6d, 0x91, 0x70, 0x42, 0x39, 0xa7, 0x6c, 0x8a, 0x5d, 0xc8,
	0xc7, 0xda, 0x24, 0x8c, 0xce, 0xb8, 0x26, 0xd7, 0xab, 0xa5, 0xf8, 0x25, 0xdc, 0xdc, 0xae, 0xb9,
	0xce, 0x2c, 0x92, 0x9d, 0x58, 0x01, 0x08, 0xae, 0xae, 0x89, 0x12, 0x29, 0xb4, 0x53, 0x15, 0xed,
	0x1d, 0x82, 0xca, 0x5f, 0x8f, 0x4f, 0xb9, 0xa1, 0x7f, 0xee, 0x96, 0xfd, 0xd3, 0xed, 0xe0, 0x2b,
	0x02, 0xf8, 0xf5, 0xc9, 0xf0, 0x21, 0xdc, 0xed, 0x39, 0x5d, 0xcb, 0x75, 0x5a, 0x5d, 0xdb, 0x69,
	0xba, 0xcf, 0x9b, 0x9d, 0x96, 0xd5, 0xb0, 0x9f, 0xd8, 0xd6, 0x63, 0x29, 0x23, 0x97, 0x16, 0x4b,
	0x75, 0x2f, 0x06, 0xad, 0x49, 0x20, 0xe6, 0x58, 0x83, 0x52, 0x9a, 0xee, 0x5b, 0x1d, 0x09, 0xc9,
	0x85, 0xc5, 0x52, 0xbd, 0x15, 0x53, 0x7d, 0xc2, 0xf1, 0x01, 0xdc, 0x49, 0x33, 0x66, 0xbd, 0xd3,
	0x35, 0xed, 0xa6, 0x94, 0x95, 0x6f, 0x2f, 0x96, 0x6a, 0x21, 0xe6, 0xcc, 0x01, 0x17, 0x1e, 0x9d,
	0x62, 0x15, 0x8a, 0x69, 0xb6, 0xe9, 0x48, 0x3b, 0xf2, 0xfe, 0x62, 0xa9, 0xe6, 0x63, 0xac, 0xc9,
	0xf0, 0x31, 0x54, 0x7e, 0x27, 0xdc, 0x53, 0xbb, 0xfb, 0xd4, 0xed, 0x59, 0x5d, 0x47, 0xca, 0xc9,
	0xe5, 0xc5, 0x52, 0x95, 0x12, 0xf6, 0x94, 0x8a, 0xb3, 0x1e, 0x11, 0x4c, 0xce, 0xbd, 0x79, 0xaf,
	0x64, 0xea, 0x8f, 0x3e, 0xad, 0x14, 0x74, 0xb1, 0x52, 0xd0, 0x97, 0x95, 0x82, 0xde, 0xae, 0x95,
	0xcc, 0xc5, 0x5a, 0xc9, 0x7c, 0x5e, 0x2b, 0x99, 0x17, 0xf7, 0x53, 0xd1, 0x3e, 0xa3, 0xa1, 0xd7,
	0x60, 0x21, 0x31, 0x38, 0x19, 0x7b, 0xd4, 0x78, 0x6d, 0xf8, 0x6c, 0x16, 0xa7, 0x3b, 0xd8, 0x8d,
	0xfe, 0x1e, 0x0f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x08, 0x7a, 0xfc, 0xf1, 0x6f, 0x04, 0x00,
	0x00,
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Option != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Option))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgVoteProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVoteProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVoteProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Option != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Option))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgProposalAssignPermission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposalAssignPermission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposalAssignPermission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Permission != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Permission))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProposalAssignPermission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalAssignPermission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalAssignPermission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Permission != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Permission))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovProposal(uint64(m.ProposalId))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Option != 0 {
		n += 1 + sovProposal(uint64(m.Option))
	}
	return n
}

func (m *MsgVoteProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovProposal(uint64(m.ProposalId))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Option != 0 {
		n += 1 + sovProposal(uint64(m.Option))
	}
	return n
}

func (m *MsgProposalAssignPermission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Permission != 0 {
		n += 1 + sovProposal(uint64(m.Permission))
	}
	return n
}

func (m *ProposalAssignPermission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Permission != 0 {
		n += 1 + sovProposal(uint64(m.Permission))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = append(m.Voter[:0], dAtA[iNdEx:postIndex]...)
			if m.Voter == nil {
				m.Voter = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
			}
			m.Option = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Option |= VoteOption(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgVoteProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgVoteProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVoteProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = append(m.Voter[:0], dAtA[iNdEx:postIndex]...)
			if m.Voter == nil {
				m.Voter = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
			}
			m.Option = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Option |= VoteOption(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgProposalAssignPermission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgProposalAssignPermission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposalAssignPermission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			m.Permission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permission |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProposalAssignPermission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProposalAssignPermission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalAssignPermission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			m.Permission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permission |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)

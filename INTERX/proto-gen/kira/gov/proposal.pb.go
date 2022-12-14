// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/gov/proposal.proto

package gov

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/regen-network/cosmos-proto"
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

// VoteOption enumerates the valid vote options for a given governance proposal.
type VoteOption int32

const (
	// VOTE_OPTION_UNSPECIFIED defines a no-op vote option.
	VoteOption_VOTE_OPTION_UNSPECIFIED VoteOption = 0
	// VOTE_OPTION_YES defines a yes vote option.
	VoteOption_VOTE_OPTION_YES VoteOption = 1
	// VOTE_OPTION_ABSTAIN defines an abstain vote option.
	VoteOption_VOTE_OPTION_ABSTAIN VoteOption = 2
	// VOTE_OPTION_NO defines a no vote option.
	VoteOption_VOTE_OPTION_NO VoteOption = 3
	// VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
	VoteOption_VOTE_OPTION_NO_WITH_VETO VoteOption = 4
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
	return fileDescriptor_e108f9b9b2d46b19, []int{0}
}

type VoteResult int32

const (
	VoteResult_VOTE_RESULT_UNKNOWN               VoteResult = 0
	VoteResult_VOTE_RESULT_PASSED                VoteResult = 1
	VoteResult_VOTE_RESULT_REJECTED              VoteResult = 2
	VoteResult_VOTE_RESULT_REJECTED_WITH_VETO    VoteResult = 3
	VoteResult_VOTE_PENDING                      VoteResult = 4
	VoteResult_VOTE_RESULT_QUORUM_NOT_REACHED    VoteResult = 5
	VoteResult_VOTE_RESULT_ENACTMENT             VoteResult = 6
	VoteResult_VOTE_RESULT_PASSED_WITH_EXEC_FAIL VoteResult = 7
)

var VoteResult_name = map[int32]string{
	0: "VOTE_RESULT_UNKNOWN",
	1: "VOTE_RESULT_PASSED",
	2: "VOTE_RESULT_REJECTED",
	3: "VOTE_RESULT_REJECTED_WITH_VETO",
	4: "VOTE_PENDING",
	5: "VOTE_RESULT_QUORUM_NOT_REACHED",
	6: "VOTE_RESULT_ENACTMENT",
	7: "VOTE_RESULT_PASSED_WITH_EXEC_FAIL",
}

var VoteResult_value = map[string]int32{
	"VOTE_RESULT_UNKNOWN":               0,
	"VOTE_RESULT_PASSED":                1,
	"VOTE_RESULT_REJECTED":              2,
	"VOTE_RESULT_REJECTED_WITH_VETO":    3,
	"VOTE_PENDING":                      4,
	"VOTE_RESULT_QUORUM_NOT_REACHED":    5,
	"VOTE_RESULT_ENACTMENT":             6,
	"VOTE_RESULT_PASSED_WITH_EXEC_FAIL": 7,
}

func (x VoteResult) String() string {
	return proto.EnumName(VoteResult_name, int32(x))
}

func (VoteResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e108f9b9b2d46b19, []int{1}
}

type Vote struct {
	ProposalId           uint64     `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Voter                []byte     `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	Option               VoteOption `protobuf:"varint,3,opt,name=option,proto3,enum=kira.gov.VoteOption" json:"option,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_e108f9b9b2d46b19, []int{0}
}

func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
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

func (m *Vote) GetVoter() []byte {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *Vote) GetOption() VoteOption {
	if m != nil {
		return m.Option
	}
	return VoteOption_VOTE_OPTION_UNSPECIFIED
}

type MsgVoteProposal struct {
	ProposalId           uint64     `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Voter                []byte     `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	Option               VoteOption `protobuf:"varint,3,opt,name=option,proto3,enum=kira.gov.VoteOption" json:"option,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgVoteProposal) Reset()         { *m = MsgVoteProposal{} }
func (m *MsgVoteProposal) String() string { return proto.CompactTextString(m) }
func (*MsgVoteProposal) ProtoMessage()    {}
func (*MsgVoteProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e108f9b9b2d46b19, []int{1}
}

func (m *MsgVoteProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgVoteProposal.Unmarshal(m, b)
}
func (m *MsgVoteProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgVoteProposal.Marshal(b, m, deterministic)
}
func (m *MsgVoteProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVoteProposal.Merge(m, src)
}
func (m *MsgVoteProposal) XXX_Size() int {
	return xxx_messageInfo_MsgVoteProposal.Size(m)
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

func (m *MsgVoteProposal) GetVoter() []byte {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *MsgVoteProposal) GetOption() VoteOption {
	if m != nil {
		return m.Option
	}
	return VoteOption_VOTE_OPTION_UNSPECIFIED
}

type Proposal struct {
	ProposalId                 uint64               `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Title                      string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description                string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Content                    *any.Any             `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	SubmitTime                 *timestamp.Timestamp `protobuf:"bytes,5,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	VotingEndTime              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=voting_end_time,json=votingEndTime,proto3" json:"voting_end_time,omitempty"`
	EnactmentEndTime           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=enactment_end_time,json=enactmentEndTime,proto3" json:"enactment_end_time,omitempty"`
	MinVotingEndBlockHeight    int64                `protobuf:"varint,8,opt,name=min_voting_end_block_height,json=minVotingEndBlockHeight,proto3" json:"min_voting_end_block_height,omitempty"`
	MinEnactmentEndBlockHeight int64                `protobuf:"varint,9,opt,name=min_enactment_end_block_height,json=minEnactmentEndBlockHeight,proto3" json:"min_enactment_end_block_height,omitempty"`
	Result                     VoteResult           `protobuf:"varint,10,opt,name=result,proto3,enum=kira.gov.VoteResult" json:"result,omitempty"`
	ExecResult                 string               `protobuf:"bytes,11,opt,name=exec_result,json=execResult,proto3" json:"exec_result,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}             `json:"-"`
	XXX_unrecognized           []byte               `json:"-"`
	XXX_sizecache              int32                `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e108f9b9b2d46b19, []int{2}
}

func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *Proposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Proposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Proposal) GetContent() *any.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Proposal) GetSubmitTime() *timestamp.Timestamp {
	if m != nil {
		return m.SubmitTime
	}
	return nil
}

func (m *Proposal) GetVotingEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.VotingEndTime
	}
	return nil
}

func (m *Proposal) GetEnactmentEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EnactmentEndTime
	}
	return nil
}

func (m *Proposal) GetMinVotingEndBlockHeight() int64 {
	if m != nil {
		return m.MinVotingEndBlockHeight
	}
	return 0
}

func (m *Proposal) GetMinEnactmentEndBlockHeight() int64 {
	if m != nil {
		return m.MinEnactmentEndBlockHeight
	}
	return 0
}

func (m *Proposal) GetResult() VoteResult {
	if m != nil {
		return m.Result
	}
	return VoteResult_VOTE_RESULT_UNKNOWN
}

func (m *Proposal) GetExecResult() string {
	if m != nil {
		return m.ExecResult
	}
	return ""
}

type AssignPermissionProposal struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Permission           uint32   `protobuf:"varint,2,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignPermissionProposal) Reset()         { *m = AssignPermissionProposal{} }
func (m *AssignPermissionProposal) String() string { return proto.CompactTextString(m) }
func (*AssignPermissionProposal) ProtoMessage()    {}
func (*AssignPermissionProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e108f9b9b2d46b19, []int{3}
}

func (m *AssignPermissionProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignPermissionProposal.Unmarshal(m, b)
}
func (m *AssignPermissionProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignPermissionProposal.Marshal(b, m, deterministic)
}
func (m *AssignPermissionProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignPermissionProposal.Merge(m, src)
}
func (m *AssignPermissionProposal) XXX_Size() int {
	return xxx_messageInfo_AssignPermissionProposal.Size(m)
}
func (m *AssignPermissionProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignPermissionProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AssignPermissionProposal proto.InternalMessageInfo

func (m *AssignPermissionProposal) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AssignPermissionProposal) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

type SetNetworkPropertyProposal struct {
	NetworkProperty      NetworkProperty       `protobuf:"varint,1,opt,name=network_property,json=networkProperty,proto3,enum=kira.gov.NetworkProperty" json:"network_property,omitempty"`
	Value                *NetworkPropertyValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SetNetworkPropertyProposal) Reset()         { *m = SetNetworkPropertyProposal{} }
func (m *SetNetworkPropertyProposal) String() string { return proto.CompactTextString(m) }
func (*SetNetworkPropertyProposal) ProtoMessage()    {}
func (*SetNetworkPropertyProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e108f9b9b2d46b19, []int{4}
}

func (m *SetNetworkPropertyProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetNetworkPropertyProposal.Unmarshal(m, b)
}
func (m *SetNetworkPropertyProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetNetworkPropertyProposal.Marshal(b, m, deterministic)
}
func (m *SetNetworkPropertyProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetNetworkPropertyProposal.Merge(m, src)
}
func (m *SetNetworkPropertyProposal) XXX_Size() int {
	return xxx_messageInfo_SetNetworkPropertyProposal.Size(m)
}
func (m *SetNetworkPropertyProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetNetworkPropertyProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetNetworkPropertyProposal proto.InternalMessageInfo

func (m *SetNetworkPropertyProposal) GetNetworkProperty() NetworkProperty {
	if m != nil {
		return m.NetworkProperty
	}
	return NetworkProperty_MIN_TX_FEE
}

func (m *SetNetworkPropertyProposal) GetValue() *NetworkPropertyValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type UpsertDataRegistryProposal struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Reference            string   `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	Encoding             string   `protobuf:"bytes,4,opt,name=encoding,proto3" json:"encoding,omitempty"`
	Size                 uint64   `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertDataRegistryProposal) Reset()         { *m = UpsertDataRegistryProposal{} }
func (m *UpsertDataRegistryProposal) String() string { return proto.CompactTextString(m) }
func (*UpsertDataRegistryProposal) ProtoMessage()    {}
func (*UpsertDataRegistryProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e108f9b9b2d46b19, []int{5}
}

func (m *UpsertDataRegistryProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertDataRegistryProposal.Unmarshal(m, b)
}
func (m *UpsertDataRegistryProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertDataRegistryProposal.Marshal(b, m, deterministic)
}
func (m *UpsertDataRegistryProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertDataRegistryProposal.Merge(m, src)
}
func (m *UpsertDataRegistryProposal) XXX_Size() int {
	return xxx_messageInfo_UpsertDataRegistryProposal.Size(m)
}
func (m *UpsertDataRegistryProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertDataRegistryProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertDataRegistryProposal proto.InternalMessageInfo

func (m *UpsertDataRegistryProposal) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UpsertDataRegistryProposal) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *UpsertDataRegistryProposal) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *UpsertDataRegistryProposal) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *UpsertDataRegistryProposal) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type SetPoorNetworkMessagesProposal struct {
	Messages             []string `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetPoorNetworkMessagesProposal) Reset()         { *m = SetPoorNetworkMessagesProposal{} }
func (m *SetPoorNetworkMessagesProposal) String() string { return proto.CompactTextString(m) }
func (*SetPoorNetworkMessagesProposal) ProtoMessage()    {}
func (*SetPoorNetworkMessagesProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e108f9b9b2d46b19, []int{6}
}

func (m *SetPoorNetworkMessagesProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetPoorNetworkMessagesProposal.Unmarshal(m, b)
}
func (m *SetPoorNetworkMessagesProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetPoorNetworkMessagesProposal.Marshal(b, m, deterministic)
}
func (m *SetPoorNetworkMessagesProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPoorNetworkMessagesProposal.Merge(m, src)
}
func (m *SetPoorNetworkMessagesProposal) XXX_Size() int {
	return xxx_messageInfo_SetPoorNetworkMessagesProposal.Size(m)
}
func (m *SetPoorNetworkMessagesProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPoorNetworkMessagesProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetPoorNetworkMessagesProposal proto.InternalMessageInfo

func (m *SetPoorNetworkMessagesProposal) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

type CreateRoleProposal struct {
	RoleSid                string      `protobuf:"bytes,1,opt,name=role_sid,json=roleSid,proto3" json:"role_sid,omitempty"`
	RoleDescription        string      `protobuf:"bytes,2,opt,name=role_description,json=roleDescription,proto3" json:"role_description,omitempty"`
	WhitelistedPermissions []PermValue `protobuf:"varint,3,rep,packed,name=whitelisted_permissions,json=whitelistedPermissions,proto3,enum=kira.gov.PermValue" json:"whitelisted_permissions,omitempty"`
	BlacklistedPermissions []PermValue `protobuf:"varint,4,rep,packed,name=blacklisted_permissions,json=blacklistedPermissions,proto3,enum=kira.gov.PermValue" json:"blacklisted_permissions,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}    `json:"-"`
	XXX_unrecognized       []byte      `json:"-"`
	XXX_sizecache          int32       `json:"-"`
}

func (m *CreateRoleProposal) Reset()         { *m = CreateRoleProposal{} }
func (m *CreateRoleProposal) String() string { return proto.CompactTextString(m) }
func (*CreateRoleProposal) ProtoMessage()    {}
func (*CreateRoleProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e108f9b9b2d46b19, []int{7}
}

func (m *CreateRoleProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleProposal.Unmarshal(m, b)
}
func (m *CreateRoleProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleProposal.Marshal(b, m, deterministic)
}
func (m *CreateRoleProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleProposal.Merge(m, src)
}
func (m *CreateRoleProposal) XXX_Size() int {
	return xxx_messageInfo_CreateRoleProposal.Size(m)
}
func (m *CreateRoleProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleProposal proto.InternalMessageInfo

func (m *CreateRoleProposal) GetRoleSid() string {
	if m != nil {
		return m.RoleSid
	}
	return ""
}

func (m *CreateRoleProposal) GetRoleDescription() string {
	if m != nil {
		return m.RoleDescription
	}
	return ""
}

func (m *CreateRoleProposal) GetWhitelistedPermissions() []PermValue {
	if m != nil {
		return m.WhitelistedPermissions
	}
	return nil
}

func (m *CreateRoleProposal) GetBlacklistedPermissions() []PermValue {
	if m != nil {
		return m.BlacklistedPermissions
	}
	return nil
}

type SetProposalDurationsProposal struct {
	TypeofProposals      []string `protobuf:"bytes,1,rep,name=typeof_proposals,json=typeofProposals,proto3" json:"typeof_proposals,omitempty"`
	ProposalDurations    []uint64 `protobuf:"varint,2,rep,packed,name=proposal_durations,json=proposalDurations,proto3" json:"proposal_durations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetProposalDurationsProposal) Reset()         { *m = SetProposalDurationsProposal{} }
func (m *SetProposalDurationsProposal) String() string { return proto.CompactTextString(m) }
func (*SetProposalDurationsProposal) ProtoMessage()    {}
func (*SetProposalDurationsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e108f9b9b2d46b19, []int{8}
}

func (m *SetProposalDurationsProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetProposalDurationsProposal.Unmarshal(m, b)
}
func (m *SetProposalDurationsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetProposalDurationsProposal.Marshal(b, m, deterministic)
}
func (m *SetProposalDurationsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetProposalDurationsProposal.Merge(m, src)
}
func (m *SetProposalDurationsProposal) XXX_Size() int {
	return xxx_messageInfo_SetProposalDurationsProposal.Size(m)
}
func (m *SetProposalDurationsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetProposalDurationsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetProposalDurationsProposal proto.InternalMessageInfo

func (m *SetProposalDurationsProposal) GetTypeofProposals() []string {
	if m != nil {
		return m.TypeofProposals
	}
	return nil
}

func (m *SetProposalDurationsProposal) GetProposalDurations() []uint64 {
	if m != nil {
		return m.ProposalDurations
	}
	return nil
}

func init() {
	proto.RegisterEnum("kira.gov.VoteOption", VoteOption_name, VoteOption_value)
	proto.RegisterEnum("kira.gov.VoteResult", VoteResult_name, VoteResult_value)
	proto.RegisterType((*Vote)(nil), "kira.gov.Vote")
	proto.RegisterType((*MsgVoteProposal)(nil), "kira.gov.MsgVoteProposal")
	proto.RegisterType((*Proposal)(nil), "kira.gov.Proposal")
	proto.RegisterType((*AssignPermissionProposal)(nil), "kira.gov.AssignPermissionProposal")
	proto.RegisterType((*SetNetworkPropertyProposal)(nil), "kira.gov.SetNetworkPropertyProposal")
	proto.RegisterType((*UpsertDataRegistryProposal)(nil), "kira.gov.UpsertDataRegistryProposal")
	proto.RegisterType((*SetPoorNetworkMessagesProposal)(nil), "kira.gov.SetPoorNetworkMessagesProposal")
	proto.RegisterType((*CreateRoleProposal)(nil), "kira.gov.CreateRoleProposal")
	proto.RegisterType((*SetProposalDurationsProposal)(nil), "kira.gov.SetProposalDurationsProposal")
}

func init() {
	proto.RegisterFile("kira/gov/proposal.proto", fileDescriptor_e108f9b9b2d46b19)
}

var fileDescriptor_e108f9b9b2d46b19 = []byte{
	// 1340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcf, 0x73, 0xdb, 0xc4,
	0x17, 0x8f, 0x62, 0xe7, 0x87, 0xd7, 0x49, 0xac, 0x6e, 0xf3, 0x6d, 0x1c, 0x7d, 0x4b, 0xac, 0x7a,
	0x80, 0x71, 0x3b, 0x8d, 0x3d, 0xa4, 0xc3, 0x0c, 0x13, 0xca, 0xc1, 0x3f, 0xd4, 0xd6, 0xb4, 0x91,
	0x5d, 0xd9, 0x49, 0x5b, 0x18, 0x46, 0x23, 0x4b, 0x1b, 0x79, 0xb1, 0xa5, 0xf5, 0x68, 0xd7, 0x69,
	0xcd, 0x89, 0x1b, 0x1d, 0x9f, 0xf8, 0x07, 0xcc, 0x94, 0xe1, 0x04, 0x33, 0x0c, 0x17, 0xfe, 0x08,
	0x86, 0x23, 0x7f, 0x40, 0xb8, 0x31, 0x9c, 0x7b, 0xe4, 0xc4, 0x48, 0x2b, 0x59, 0x4a, 0x52, 0x28,
	0xdc, 0x38, 0x59, 0xfb, 0xde, 0xe7, 0xbd, 0xcf, 0xe7, 0x3d, 0xbd, 0x7d, 0x32, 0xd8, 0x1a, 0x60,
	0xcf, 0xa8, 0xd8, 0xe4, 0xa4, 0x32, 0xf2, 0xc8, 0x88, 0x50, 0x63, 0x58, 0x1e, 0x79, 0x84, 0x11,
	0xb8, 0xea, 0x3b, 0xca, 0x36, 0x39, 0x91, 0x36, 0x6d, 0x62, 0x93, 0xc0, 0x58, 0xf1, 0x9f, 0xb8,
	0x5f, 0x2a, 0xd8, 0x84, 0xd8, 0x43, 0x54, 0x09, 0x4e, 0xbd, 0xf1, 0x71, 0x85, 0x61, 0x07, 0x51,
	0x66, 0x38, 0xa3, 0x10, 0xb0, 0x7d, 0x1e, 0x60, 0xb8, 0x93, 0xc8, 0x65, 0x12, 0xea, 0x10, 0xaa,
	0xf3, 0xa4, 0xfc, 0x10, 0xba, 0xae, 0xcd, 0xf5, 0xb8, 0x88, 0x3d, 0x25, 0xde, 0xc0, 0x07, 0x8d,
	0x90, 0xc7, 0x30, 0x8a, 0x20, 0xdb, 0xb1, 0x64, 0xe4, 0x39, 0x98, 0x52, 0x4c, 0x5c, 0xee, 0x2a,
	0x7e, 0x25, 0x80, 0xf4, 0x11, 0x61, 0x08, 0x16, 0x40, 0x36, 0xaa, 0x47, 0xc7, 0x56, 0x5e, 0x90,
	0x85, 0x52, 0x5a, 0x03, 0x91, 0xa9, 0x69, 0xc1, 0xbb, 0x60, 0xe9, 0x84, 0x30, 0xe4, 0xe5, 0x17,
	0x65, 0xa1, 0xb4, 0x56, 0x7b, 0xe7, 0x8f, 0xd3, 0xc2, 0xae, 0x8d, 0x59, 0x7f, 0xdc, 0x2b, 0x9b,
	0xc4, 0x09, 0x35, 0x85, 0x3f, 0xbb, 0xd4, 0x1a, 0x54, 0xd8, 0x64, 0x84, 0x68, 0xb9, 0x6a, 0x9a,
	0x55, 0xcb, 0xf2, 0x10, 0xa5, 0x1a, 0x8f, 0x87, 0x37, 0xc1, 0x32, 0x19, 0x31, 0x4c, 0xdc, 0x7c,
	0x4a, 0x16, 0x4a, 0x1b, 0x7b, 0x9b, 0xe5, 0xa8, 0x71, 0x65, 0x5f, 0x49, 0x2b, 0xf0, 0x69, 0x21,
	0xa6, 0xf8, 0xad, 0x00, 0x72, 0x07, 0xd4, 0xf6, 0x3d, 0xed, 0x50, 0xcc, 0x7f, 0x56, 0xeb, 0xf7,
	0x4b, 0x60, 0xf5, 0x9f, 0x8b, 0xdc, 0x04, 0x4b, 0x0c, 0xb3, 0x21, 0x0a, 0x44, 0x66, 0x34, 0x7e,
	0x80, 0x32, 0xc8, 0x5a, 0x88, 0x9a, 0x1e, 0x8e, 0x69, 0x33, 0x5a, 0xd2, 0x04, 0xdf, 0x07, 0x2b,
	0x26, 0x71, 0x19, 0x72, 0x59, 0x3e, 0x2d, 0x0b, 0xa5, 0xec, 0xde, 0x66, 0x99, 0x0f, 0x4e, 0x39,
	0x1a, 0x9c, 0x72, 0xd5, 0x9d, 0xd4, 0xb2, 0x3f, 0xff, 0xb8, 0xbb, 0x52, 0xe7, 0x40, 0x2d, 0x8a,
	0x80, 0x1f, 0x83, 0x2c, 0x1d, 0xf7, 0x1c, 0xcc, 0x74, 0x7f, 0xfa, 0xf2, 0x4b, 0x41, 0x02, 0xe9,
	0x42, 0x82, 0x6e, 0x34, 0x9a, 0xb5, 0x9d, 0x9f, 0x4e, 0x0b, 0x0b, 0x2f, 0x4f, 0x0b, 0x70, 0x62,
	0x38, 0xc3, 0xfd, 0x62, 0x22, 0xb8, 0xf8, 0xe5, 0xaf, 0x05, 0x41, 0x03, 0xdc, 0xe2, 0x07, 0xc0,
	0x63, 0x90, 0x3b, 0x21, 0x0c, 0xbb, 0xb6, 0x8e, 0x5c, 0x8b, 0x13, 0x2c, 0xbf, 0x96, 0xa0, 0x18,
	0x12, 0x5c, 0xe1, 0x04, 0xe7, 0x12, 0x70, 0x92, 0x75, 0x6e, 0x55, 0x5c, 0x2b, 0xe0, 0x21, 0x00,
	0x22, 0xd7, 0x30, 0x99, 0x83, 0x5c, 0x16, 0x53, 0xad, 0xbc, 0x96, 0xea, 0xad, 0x90, 0x6a, 0x9b,
	0x53, 0x5d, 0xcc, 0xc1, 0xd9, 0xc4, 0xb9, 0x23, 0x22, 0xbc, 0x0d, 0xfe, 0xef, 0x60, 0x57, 0x4f,
	0x68, 0xeb, 0x0d, 0x89, 0x39, 0xd0, 0xfb, 0x08, 0xdb, 0x7d, 0x96, 0x5f, 0x95, 0x85, 0x52, 0x4a,
	0xdb, 0x72, 0xb0, 0x7b, 0x14, 0xe9, 0xac, 0xf9, 0xfe, 0x7b, 0x81, 0x1b, 0xd6, 0xc0, 0x8e, 0x1f,
	0x7d, 0x96, 0xee, 0x4c, 0x82, 0x4c, 0x90, 0x40, 0x72, 0xb0, 0xab, 0x24, 0xa8, 0x93, 0x39, 0x6e,
	0x82, 0x65, 0x0f, 0xd1, 0xf1, 0x90, 0xe5, 0xc1, 0xab, 0x06, 0x51, 0x0b, 0x7c, 0x5a, 0x88, 0xf1,
	0x67, 0x0f, 0x3d, 0x43, 0xa6, 0x1e, 0x86, 0x64, 0x83, 0x21, 0x02, 0xbe, 0x89, 0x03, 0xf7, 0xd3,
	0xcf, 0x5f, 0x14, 0x16, 0x8a, 0xdf, 0x09, 0x20, 0x5f, 0xa5, 0x14, 0xdb, 0x6e, 0x7b, 0xbe, 0x17,
	0xe6, 0xf3, 0xfb, 0x09, 0x58, 0x31, 0xf8, 0x65, 0x08, 0x66, 0x77, 0xad, 0x56, 0x7f, 0x79, 0x5a,
	0xd8, 0xe0, 0x9d, 0x0b, 0x1d, 0xc5, 0x7f, 0x7f, 0xaf, 0xa2, 0x9c, 0x70, 0x07, 0x80, 0x78, 0x19,
	0x05, 0x57, 0x60, 0x5d, 0x4b, 0x58, 0xf6, 0x73, 0xbf, 0xbf, 0x28, 0x08, 0xbf, 0xc4, 0x23, 0x5c,
	0xfc, 0x41, 0x00, 0x52, 0x07, 0x31, 0x95, 0x2f, 0xb9, 0x36, 0xdf, 0x71, 0x93, 0xb9, 0xdc, 0x06,
	0x10, 0xcf, 0xed, 0xbf, 0x49, 0xa0, 0x7b, 0x63, 0x6f, 0x3b, 0x6e, 0xd5, 0xb9, 0x60, 0x2d, 0xe7,
	0x9e, 0x35, 0xc0, 0x7d, 0xb0, 0x74, 0x62, 0x0c, 0xc7, 0xfc, 0x4e, 0x66, 0xf7, 0x76, 0xfe, 0x32,
	0xf4, 0xc8, 0x47, 0xd5, 0xd2, 0xfe, 0x40, 0x69, 0x3c, 0xe4, 0xa2, 0xe2, 0xaf, 0x05, 0x20, 0x1d,
	0x8e, 0x28, 0xf2, 0x58, 0xc3, 0x60, 0x86, 0x86, 0x6c, 0x4c, 0x99, 0x17, 0x2b, 0x16, 0x41, 0x6a,
	0x80, 0xb8, 0xc8, 0x8c, 0xe6, 0x3f, 0x42, 0x08, 0xd2, 0x7d, 0x83, 0xf6, 0xc3, 0x85, 0x10, 0x3c,
	0xc3, 0xab, 0x20, 0xe3, 0xa1, 0x63, 0xe4, 0x21, 0xd7, 0x44, 0xe1, 0x36, 0x88, 0x0d, 0x50, 0x02,
	0xab, 0xc8, 0x35, 0x89, 0x85, 0x5d, 0x3b, 0x58, 0x06, 0x19, 0x6d, 0x7e, 0xf6, 0xb3, 0x51, 0xfc,
	0x19, 0xbf, 0xe3, 0x69, 0x2d, 0x78, 0xbe, 0xa8, 0xf1, 0x36, 0xd8, 0xe9, 0x20, 0xd6, 0x26, 0xc4,
	0x0b, 0x0b, 0x3c, 0x40, 0x94, 0x1a, 0x36, 0xa2, 0x73, 0x99, 0x12, 0x58, 0x75, 0x42, 0x5b, 0x5e,
	0x90, 0x53, 0x3e, 0x45, 0x74, 0x2e, 0x7e, 0xb1, 0x08, 0x60, 0xdd, 0x43, 0x06, 0x43, 0x1a, 0x19,
	0xc6, 0xfb, 0x79, 0x1b, 0xac, 0x7a, 0x64, 0x88, 0x74, 0x1a, 0xee, 0xbd, 0x8c, 0xb6, 0xe2, 0x9f,
	0x3b, 0xd8, 0x82, 0xd7, 0x81, 0x18, 0xb8, 0x92, 0x3b, 0x8e, 0x97, 0x9b, 0xf3, 0xed, 0x8d, 0xc4,
	0x9e, 0x7b, 0x00, 0xb6, 0x9e, 0xf6, 0x31, 0x43, 0x43, 0x4c, 0x19, 0xb2, 0xf4, 0x78, 0x36, 0x68,
	0x3e, 0x25, 0xa7, 0x4a, 0x1b, 0x7b, 0x97, 0xe3, 0xb7, 0xe3, 0xcf, 0x6f, 0xf0, 0x4a, 0xb4, 0x2b,
	0x89, 0x98, 0x78, 0xaa, 0xa9, 0x9f, 0xad, 0x37, 0x34, 0xcc, 0xc1, 0x2b, 0xb2, 0xa5, 0xff, 0x26,
	0x5b, 0x22, 0x26, 0x91, 0xed, 0x62, 0x1f, 0x9f, 0x81, 0xab, 0x7e, 0x1f, 0xc3, 0x0e, 0x34, 0xc6,
	0x9e, 0xe1, 0xd7, 0x10, 0x77, 0xf1, 0x3a, 0x10, 0xfd, 0xbb, 0x40, 0x8e, 0xf5, 0xe8, 0x0b, 0x10,
	0x75, 0x33, 0xc7, 0xed, 0x11, 0x92, 0xc2, 0x5d, 0x00, 0xe7, 0x1f, 0x0e, 0x2b, 0x4a, 0x94, 0x5f,
	0x94, 0x53, 0xa5, 0xb4, 0x76, 0x69, 0x74, 0x9e, 0xe1, 0xc6, 0x6f, 0x02, 0x00, 0xf1, 0xb7, 0x08,
	0xde, 0x04, 0x5b, 0x47, 0xad, 0xae, 0xa2, 0xb7, 0xda, 0xdd, 0x66, 0x4b, 0xd5, 0x0f, 0xd5, 0x4e,
	0x5b, 0xa9, 0x37, 0xef, 0x34, 0x95, 0x86, 0xb8, 0x20, 0xe5, 0xa6, 0x33, 0x39, 0xcb, 0x81, 0x8a,
	0x33, 0x62, 0x13, 0x58, 0x04, 0xb9, 0x24, 0xfa, 0x89, 0xd2, 0x11, 0x05, 0x69, 0x7d, 0x3a, 0x93,
	0x33, 0x1c, 0xf5, 0x04, 0x51, 0x78, 0x03, 0x5c, 0x4e, 0x62, 0xaa, 0xb5, 0x4e, 0xb7, 0xda, 0x54,
	0xc5, 0x45, 0xe9, 0xd2, 0x74, 0x26, 0xaf, 0x73, 0x5c, 0xb5, 0x47, 0x99, 0x81, 0x5d, 0x28, 0x83,
	0x8d, 0x24, 0x56, 0x6d, 0x89, 0x29, 0x69, 0x6d, 0x3a, 0x93, 0x57, 0x39, 0x4c, 0x25, 0x70, 0x0f,
	0xe4, 0xcf, 0x22, 0xf4, 0x47, 0xcd, 0xee, 0x3d, 0xfd, 0x48, 0xe9, 0xb6, 0xc4, 0xb4, 0xb4, 0x39,
	0x9d, 0xc9, 0x62, 0x84, 0x7d, 0x84, 0x59, 0xff, 0x08, 0x31, 0x22, 0xa5, 0x9f, 0x7f, 0xb3, 0xb3,
	0x70, 0xe3, 0xf3, 0x14, 0x2f, 0x94, 0xaf, 0x30, 0xf8, 0x66, 0x28, 0x4b, 0x53, 0x3a, 0x87, 0x0f,
	0xba, 0xfa, 0xa1, 0x7a, 0x5f, 0x6d, 0x3d, 0x52, 0xc5, 0x05, 0x29, 0x3b, 0x9d, 0xc9, 0x2b, 0x87,
	0xee, 0xc0, 0x25, 0x4f, 0x5d, 0x58, 0x04, 0x30, 0x89, 0x6a, 0x57, 0x3b, 0x1d, 0xa5, 0x21, 0x0a,
	0x12, 0x98, 0xce, 0xe4, 0xe5, 0xb6, 0x41, 0x29, 0xb2, 0xe0, 0xdb, 0x60, 0x33, 0x89, 0xd1, 0x94,
	0x0f, 0x95, 0x7a, 0x57, 0x69, 0x88, 0x8b, 0x5c, 0xba, 0x86, 0x3e, 0x45, 0x26, 0x43, 0x16, 0x7c,
	0x0f, 0xec, 0xbc, 0x0a, 0x97, 0x28, 0x20, 0xc5, 0x0b, 0x88, 0x22, 0xa2, 0x02, 0xe0, 0x1b, 0x60,
	0x2d, 0x88, 0x6c, 0x2b, 0x6a, 0xa3, 0xa9, 0xde, 0x15, 0xd3, 0x5c, 0x64, 0x1b, 0xb9, 0xc1, 0x4d,
	0x3d, 0x97, 0xf8, 0xe1, 0x61, 0x4b, 0x3b, 0x3c, 0xd0, 0xd5, 0x96, 0xcf, 0x51, 0xad, 0xdf, 0x53,
	0x1a, 0xe2, 0x12, 0x4f, 0xfc, 0x70, 0x4c, 0xbc, 0xb1, 0xa3, 0x12, 0xa6, 0x21, 0xc3, 0xec, 0x23,
	0x0b, 0x96, 0xc0, 0xff, 0x92, 0x91, 0x8a, 0x5a, 0xad, 0x77, 0x0f, 0x14, 0xb5, 0x2b, 0x2e, 0xf3,
	0xb7, 0x38, 0xff, 0x9c, 0xc0, 0x0f, 0xc0, 0xb5, 0x8b, 0x8d, 0xe0, 0xd2, 0x95, 0xc7, 0x4a, 0x5d,
	0xbf, 0x53, 0x6d, 0x3e, 0x10, 0x57, 0xa4, 0x2b, 0xd3, 0x99, 0x0c, 0x79, 0x5f, 0x7c, 0xf5, 0xca,
	0x33, 0x64, 0xde, 0x31, 0xf0, 0x90, 0xbf, 0x82, 0xda, 0xbb, 0x1f, 0xdd, 0x4a, 0xac, 0xfb, 0xfb,
	0xd8, 0x33, 0xea, 0xc4, 0x43, 0x15, 0x8a, 0x06, 0x06, 0xae, 0x34, 0xd5, 0xae, 0xa2, 0x3d, 0xe6,
	0xff, 0x5e, 0x77, 0x6d, 0xe4, 0x56, 0xa2, 0x7f, 0x9d, 0xbd, 0xe5, 0xc0, 0x76, 0xeb, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0x9b, 0xb1, 0x2f, 0x3b, 0x0b, 0x00, 0x00,
}

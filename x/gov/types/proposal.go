package types

import (
	"fmt"
	"time"

	kiratypes "github.com/KiraCore/sekai/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
)

var _ Content = &AssignPermissionProposal{}

// NewProposal creates a new proposal
func NewProposal(
	proposalID uint64,
	title, description string,
	content Content,
	submitTime time.Time,
	votingEndTime time.Time,
	enactmentEndTime time.Time,
	minVotingEndBlockHeight int64,
	minEnactmentEndBlockHeight int64,
) (Proposal, error) {
	msg, ok := content.(proto.Message)
	if !ok {
		return Proposal{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%T does not implement proto.Message", content))
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return Proposal{}, err
	}

	return Proposal{
		ProposalId:                 proposalID,
		Title:                      title,
		Description:                description,
		SubmitTime:                 submitTime,
		VotingEndTime:              votingEndTime,
		EnactmentEndTime:           enactmentEndTime,
		MinVotingEndBlockHeight:    minVotingEndBlockHeight,
		MinEnactmentEndBlockHeight: minEnactmentEndBlockHeight,
		Content:                    any,
		Result:                     Pending,
	}, nil
}

// GetContent returns the proposal Content
func (p Proposal) GetContent() Content {
	content, ok := p.Content.GetCachedValue().(Content)
	if !ok {
		return nil
	}
	return content
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var content Content
	return unpacker.UnpackAny(p.Content, &content)
}

// NewAssignPermissionProposal creates a new assign permission proposal
func NewAssignPermissionProposal(
	address types.AccAddress,
	permission PermValue,
) Content {
	return &AssignPermissionProposal{
		Address:    address,
		Permission: uint32(permission),
	}
}

// ProposalType returns proposal's type
func (m *AssignPermissionProposal) ProposalType() string {
	return kiratypes.AssignPermissionProposalType
}

// ValidateBasic returns basic validation
func (m *AssignPermissionProposal) ValidateBasic() error {
	if m.Address.Empty() {
		return ErrEmptyPermissionsAccAddress
	}
	return nil
}

func (m *AssignPermissionProposal) ProposalPermission() PermValue {
	return PermCreateSetPermissionsProposal
}

func (m *AssignPermissionProposal) VotePermission() PermValue {
	return PermVoteSetPermissionProposal
}

// NewSetNetworkPropertyProposal creates a new set network property proposal
func NewSetNetworkPropertyProposal(
	property NetworkProperty,
	value NetworkPropertyValue,
) Content {
	return &SetNetworkPropertyProposal{
		NetworkProperty: property,
		Value:           value,
	}
}

// ProposalType returns proposal's type
func (m *SetNetworkPropertyProposal) ProposalType() string {
	return kiratypes.SetNetworkPropertyProposalType
}

func (m *SetNetworkPropertyProposal) ProposalPermission() PermValue {
	return PermCreateSetNetworkPropertyProposal
}

// VotePermission returns permission to vote on this proposal
func (m *SetNetworkPropertyProposal) VotePermission() PermValue {
	return PermVoteSetNetworkPropertyProposal
}

// ValidateBasic returns basic validation
func (m *SetNetworkPropertyProposal) ValidateBasic() error {
	switch m.NetworkProperty {
	case MinTxFee,
		MaxTxFee,
		VoteQuorum,
		MinimumProposalEndTime,
		ProposalEnactmentTime,
		EnableForeignFeePayments,
		MischanceRankDecreaseAmount,
		MischanceConfidence,
		MaxMischance,
		InactiveRankDecreasePercent,
		PoorNetworkMaxBankSend,
		MinValidators,
		UnjailMaxTime,
		EnableTokenWhitelist,
		EnableTokenBlacklist,
		MinIdentityApprovalTip,
		UniqueIdentityKeys:
		return nil
	default:
		return ErrInvalidNetworkProperty
	}
}

func NewUpsertDataRegistryProposal(key, hash, reference, encoding string, size uint64) Content {
	return &UpsertDataRegistryProposal{
		Key:       key,
		Hash:      hash,
		Reference: reference,
		Encoding:  encoding,
		Size_:     size,
	}
}

func (m *UpsertDataRegistryProposal) ProposalType() string {
	return kiratypes.UpsertDataRegistryProposalType
}

func (m *UpsertDataRegistryProposal) ProposalPermission() PermValue {
	return PermCreateUpsertDataRegistryProposal
}

func (m *UpsertDataRegistryProposal) VotePermission() PermValue {
	return PermVoteUpsertDataRegistryProposal
}

// ValidateBasic returns basic validation
func (m *UpsertDataRegistryProposal) ValidateBasic() error {
	return nil
}

func NewSetPoorNetworkMessagesProposal(msgs []string) Content {
	return &SetPoorNetworkMessagesProposal{
		Messages: msgs,
	}
}

func (m *SetPoorNetworkMessagesProposal) ProposalType() string {
	return kiratypes.SetPoorNetworkMessagesProposalType
}

func (m *SetPoorNetworkMessagesProposal) ProposalPermission() PermValue {
	return PermCreateSetPoorNetworkMessagesProposal
}

func (m *SetPoorNetworkMessagesProposal) VotePermission() PermValue {
	return PermVoteSetPoorNetworkMessagesProposal
}

// ValidateBasic returns basic validation
func (m *SetPoorNetworkMessagesProposal) ValidateBasic() error {
	return nil
}

func NewCreateRoleProposal(sid, description string, whitelist []PermValue, blacklist []PermValue) Content {
	return &CreateRoleProposal{
		RoleSid:                sid,
		RoleDescription:        description,
		WhitelistedPermissions: whitelist,
		BlacklistedPermissions: blacklist,
	}
}

func (m *CreateRoleProposal) ProposalType() string {
	return kiratypes.CreateRoleProposalType
}

func (m *CreateRoleProposal) ProposalPermission() PermValue {
	return PermCreateRoleProposal
}

func (m *CreateRoleProposal) VotePermission() PermValue {
	return PermVoteCreateRoleProposal
}

// ValidateBasic returns basic validation
func (m *CreateRoleProposal) ValidateBasic() error {
	if len(m.WhitelistedPermissions) == 0 && len(m.BlacklistedPermissions) == 0 {
		return ErrEmptyPermissions
	}

	return nil
}

func NewSetProposalDurationsProposal(typeofProposals []string, durations []uint64) Content {
	return &SetProposalDurationsProposal{
		TypeofProposals:   typeofProposals,
		ProposalDurations: durations,
	}
}

func (m *SetProposalDurationsProposal) ProposalType() string {
	return kiratypes.SetProposalDurationsProposalType
}

func (m *SetProposalDurationsProposal) ProposalPermission() PermValue {
	return PermCreateSetProposalDurationProposal
}

func (m *SetProposalDurationsProposal) VotePermission() PermValue {
	return PermVoteSetProposalDurationProposal
}

// ValidateBasic returns basic validation
func (m *SetProposalDurationsProposal) ValidateBasic() error {
	if len(m.TypeofProposals) == 0 {
		return fmt.Errorf("at least one proposal type should be set")
	}
	if len(m.TypeofProposals) != len(m.ProposalDurations) {
		return fmt.Errorf("the length of proposal types and durations should be equal")
	}
	for _, pt := range m.TypeofProposals {
		if pt == "" {
			return fmt.Errorf("empty proposal type is not allowed")
		}
	}
	for _, pd := range m.ProposalDurations {
		if pd == 0 {
			return fmt.Errorf("zero proposal duration is not allowed")
		}
	}
	return nil
}

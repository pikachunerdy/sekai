package gov

import (
	"fmt"

	kiratypes "github.com/KiraCore/sekai/types"
	"github.com/KiraCore/sekai/x/gov/keeper"
	"github.com/KiraCore/sekai/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
)

type ApplyAssignPermissionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyAssignPermissionProposalHandler(keeper keeper.Keeper) *ApplyAssignPermissionProposalHandler {
	return &ApplyAssignPermissionProposalHandler{keeper: keeper}
}

func (a ApplyAssignPermissionProposalHandler) ProposalType() string {
	return kiratypes.AssignPermissionProposalType
}

func (a ApplyAssignPermissionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.AssignPermissionProposal)

	actor, found := a.keeper.GetNetworkActorByAddress(ctx, p.Address)
	if found {
		if actor.Permissions.IsWhitelisted(types.PermValue(p.Permission)) {
			return sdkerrors.Wrap(types.ErrWhitelisting, "permission already whitelisted")
		}
	} else {
		actor = types.NewDefaultActor(p.Address)
	}

	return a.keeper.AddWhitelistPermission(ctx, actor, types.PermValue(p.Permission))
}

type ApplySetNetworkPropertyProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplySetNetworkPropertyProposalHandler(keeper keeper.Keeper) *ApplySetNetworkPropertyProposalHandler {
	return &ApplySetNetworkPropertyProposalHandler{keeper: keeper}
}

func (a ApplySetNetworkPropertyProposalHandler) ProposalType() string {
	return kiratypes.SetNetworkPropertyProposalType
}

func (a ApplySetNetworkPropertyProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.SetNetworkPropertyProposal)

	property, err := a.keeper.GetNetworkProperty(ctx, p.NetworkProperty)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if property == p.Value {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "network property already set as proposed value")
	}

	return a.keeper.SetNetworkProperty(ctx, p.NetworkProperty, p.Value)
}

type ApplyUpsertDataRegistryProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUpsertDataRegistryProposalHandler(keeper keeper.Keeper) *ApplyUpsertDataRegistryProposalHandler {
	return &ApplyUpsertDataRegistryProposalHandler{keeper: keeper}
}

func (a ApplyUpsertDataRegistryProposalHandler) ProposalType() string {
	return kiratypes.UpsertDataRegistryProposalType
}

func (a ApplyUpsertDataRegistryProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.UpsertDataRegistryProposal)
	entry := types.NewDataRegistryEntry(p.Hash, p.Reference, p.Encoding, p.Size_)
	a.keeper.UpsertDataRegistryEntry(ctx, p.Key, entry)
	return nil
}

type ApplySetPoorNetworkMessagesProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplySetPoorNetworkMessagesProposalHandler(keeper keeper.Keeper) *ApplySetPoorNetworkMessagesProposalHandler {
	return &ApplySetPoorNetworkMessagesProposalHandler{keeper: keeper}
}

func (a ApplySetPoorNetworkMessagesProposalHandler) ProposalType() string {
	return kiratypes.SetPoorNetworkMessagesProposalType
}

func (a ApplySetPoorNetworkMessagesProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.SetPoorNetworkMessagesProposal)
	msgs := types.AllowedMessages{Messages: p.Messages}
	a.keeper.SavePoorNetworkMessages(ctx, &msgs)
	return nil
}

type CreateRoleProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyCreateRoleProposalHandler(keeper keeper.Keeper) *CreateRoleProposalHandler {
	return &CreateRoleProposalHandler{keeper: keeper}
}

func (c CreateRoleProposalHandler) ProposalType() string {
	return kiratypes.CreateRoleProposalType
}

func (c CreateRoleProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.CreateRoleProposal)

	// check sid is good variable naming form
	if !keeper.ValidateRoleSidKey(p.RoleSid) {
		return errors.Wrap(types.ErrInvalidRoleSid, fmt.Sprintf("invalid role sid configuration: sid=%s", p.RoleSid))
	}

	_, err := c.keeper.GetRoleBySid(ctx, p.RoleSid)
	if err == nil {
		return types.ErrRoleExist
	}

	roleId := c.keeper.CreateRole(ctx, p.RoleSid, p.RoleDescription)

	for _, w := range p.WhitelistedPermissions {
		err := c.keeper.WhitelistRolePermission(ctx, roleId, w)
		if err != nil {
			return err
		}
	}

	for _, b := range p.BlacklistedPermissions {
		err := c.keeper.BlacklistRolePermission(ctx, roleId, b)
		if err != nil {
			return err
		}
	}
	return nil
}

type SetProposalDurationsProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplySetProposalDurationsProposalHandler(keeper keeper.Keeper) *SetProposalDurationsProposalHandler {
	return &SetProposalDurationsProposalHandler{keeper: keeper}
}

func (c SetProposalDurationsProposalHandler) ProposalType() string {
	return kiratypes.SetProposalDurationsProposalType
}

func (c SetProposalDurationsProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*types.SetProposalDurationsProposal)
	for i, pt := range p.TypeofProposals {
		err := c.keeper.SetProposalDuration(ctx, pt, p.ProposalDurations[i])
		if err != nil {
			return nil
		}
	}
	return nil
}

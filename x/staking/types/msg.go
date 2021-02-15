package types

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/KiraCore/sekai/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgClaimValidator{}
	_ sdk.Msg = &MsgProposalUnjailValidator{}
)

func NewMsgClaimValidator(
	moniker string,
	website string,
	social string,
	identity string,
	comission sdk.Dec,
	valKey sdk.ValAddress,
	pubKey cryptotypes.PubKey,
) (*MsgClaimValidator, error) {
	if valKey == nil {
		return nil, fmt.Errorf("validator not set")
	}

	if pubKey == nil {
		return nil, fmt.Errorf("public key not set")
	}

	var pkAny *codectypes.Any
	var err error
	if pkAny, err = codectypes.NewAnyWithValue(pubKey); err != nil {
		return nil, err
	}

	return &MsgClaimValidator{
		Moniker:    moniker,
		Website:    website,
		Social:     social,
		Identity:   identity,
		Commission: comission,
		ValKey:     valKey,
		PubKey:     pkAny,
	}, nil
}

func (m *MsgClaimValidator) Route() string {
	return ModuleName
}

func (m *MsgClaimValidator) Type() string {
	return types.MsgTypeClaimValidator
}

func (m *MsgClaimValidator) ValidateBasic() error {
	if m.ValKey.Empty() {
		return fmt.Errorf("validator not set")
	}

	return nil
}

func (m *MsgClaimValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgClaimValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		sdk.AccAddress(m.ValKey),
	}
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m *MsgClaimValidator) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pubKey cryptotypes.PubKey
	return unpacker.UnpackAny(m.PubKey, &pubKey)
}

func NewMsgProposalUnjailValidator(proposer sdk.AccAddress, hash, reference string) *MsgProposalUnjailValidator {
	return &MsgProposalUnjailValidator{
		Proposer:  proposer,
		Hash:      hash,
		Reference: reference,
	}
}

func (m *MsgProposalUnjailValidator) Route() string {
	return ModuleName
}

func (m *MsgProposalUnjailValidator) Type() string {
	return types.MsgTypeProposalUnjailValidator
}

func (m *MsgProposalUnjailValidator) ValidateBasic() error {
	if m.Proposer.Empty() {
		return fmt.Errorf("proposer not set")
	}

	return nil
}

func (m *MsgProposalUnjailValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgProposalUnjailValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Proposer}
}

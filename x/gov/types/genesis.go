package types

import (
	kiratypes "github.com/KiraCore/sekai/types"
)

// DefaultGenesis returns the default CustomGo genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NextRoleId: 3,
		Roles: []Role{
			{
				Id:          uint32(RoleSudo),
				Sid:         "sudo",
				Description: "Sudo role",
			},
			{
				Id:          uint32(RoleValidator),
				Sid:         "validator",
				Description: "Validator role",
			},
		},
		RolePermissions: map[uint64]*Permissions{
			RoleSudo: NewPermissions([]PermValue{
				PermSetPermissions,
				PermClaimValidator,
				PermClaimCouncilor,
				PermUpsertTokenAlias,
				// PermChangeTxFee, // do not give this permission to sudo account - test does not pass
				PermUpsertTokenRate,
				PermUpsertRole,
				PermCreateSetPermissionsProposal,
				PermVoteSetPermissionProposal,
				PermCreateSetNetworkPropertyProposal,
				PermVoteSetNetworkPropertyProposal,
				PermCreateUpsertDataRegistryProposal,
				PermVoteUpsertDataRegistryProposal,
				PermCreateUpsertTokenAliasProposal,
				PermVoteUpsertTokenAliasProposal,
				PermCreateUpsertTokenRateProposal,
				PermVoteUpsertTokenRateProposal,
				PermCreateUnjailValidatorProposal,
				PermVoteUnjailValidatorProposal,
				PermCreateRoleProposal,
				PermVoteCreateRoleProposal,
				PermCreateSetProposalDurationProposal,
				PermVoteSetProposalDurationProposal,
				PermCreateTokensWhiteBlackChangeProposal,
				PermVoteTokensWhiteBlackChangeProposal,
				PermCreateSetPoorNetworkMessagesProposal,
				PermVoteSetPoorNetworkMessagesProposal,
			}, nil),
			uint64(RoleValidator): NewPermissions([]PermValue{PermClaimValidator}, nil),
		},
		StartingProposalId: 1,
		NetworkProperties: &NetworkProperties{
			MinTxFee:                    100,
			MaxTxFee:                    1000000,
			VoteQuorum:                  33,
			MinimumProposalEndTime:      300, // 300 seconds / 5 mins
			ProposalEnactmentTime:       300, // 300 seconds / 5 mins
			MinProposalEndBlocks:        2,
			MinProposalEnactmentBlocks:  1,
			EnableForeignFeePayments:    true,
			MischanceRankDecreaseAmount: 10,
			MischanceConfidence:         10,
			MaxMischance:                110,
			InactiveRankDecreasePercent: 50,      // 50%
			PoorNetworkMaxBankSend:      1000000, // 1M ukex
			MinValidators:               1,
			UnjailMaxTime:               600, // 600  seconds / 10 mins
			EnableTokenWhitelist:        false,
			EnableTokenBlacklist:        true,
			MinIdentityApprovalTip:      200,
			UniqueIdentityKeys:          "moniker,username",
		},
		ExecutionFees: []*ExecutionFee{
			{
				Name:              "Claim Validator Seat",
				TransactionType:   "claim-validator-seat",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Claim Governance Seat",
				TransactionType:   "claim-governance-seat",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Claim Proposal Type X",
				TransactionType:   "claim-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Vote Proposal Type X",
				TransactionType:   "vote-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Submit Proposal Type X",
				TransactionType:   "submit-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Veto Proposal Type X",
				TransactionType:   "veto-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Upsert Token Alias Execution Fee",
				TransactionType:   kiratypes.MsgTypeUpsertTokenAlias,
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Activate a validator",
				TransactionType:   kiratypes.MsgTypeActivate,
				ExecutionFee:      100,
				FailureFee:        1000,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Pause a validator",
				TransactionType:   kiratypes.MsgTypePause,
				ExecutionFee:      10,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Unpause a validator",
				TransactionType:   kiratypes.MsgTypeUnpause,
				ExecutionFee:      10,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
		},
		PoorNetworkMessages: &AllowedMessages{
			Messages: []string{
				kiratypes.MsgTypeSubmitProposal,
				kiratypes.MsgTypeSetNetworkProperties,
				kiratypes.MsgTypeVoteProposal,
				kiratypes.MsgTypeClaimCouncilor,
				kiratypes.MsgTypeWhitelistPermissions,
				kiratypes.MsgTypeBlacklistPermissions,
				kiratypes.MsgTypeCreateRole,
				kiratypes.MsgTypeAssignRole,
				kiratypes.MsgTypeRemoveRole,
				kiratypes.MsgTypeWhitelistRolePermission,
				kiratypes.MsgTypeBlacklistRolePermission,
				kiratypes.MsgTypeRemoveWhitelistRolePermission,
				kiratypes.MsgTypeRemoveBlacklistRolePermission,
				kiratypes.MsgTypeClaimValidator,
				kiratypes.MsgTypeActivate,
				kiratypes.MsgTypePause,
				kiratypes.MsgTypeUnpause,
				kiratypes.MsgTypeRegisterIdentityRecords,
				kiratypes.MsgTypeEditIdentityRecord,
				kiratypes.MsgTypeRequestIdentityRecordsVerify,
				kiratypes.MsgTypeHandleIdentityRecordsVerifyRequest,
				kiratypes.MsgTypeCancelIdentityRecordsVerifyRequest,
			},
		},
		LastIdentityRecordId:        0,
		LastIdRecordVerifyRequestId: 0,
	}
}

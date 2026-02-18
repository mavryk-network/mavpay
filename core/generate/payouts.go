package generate

import (
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/gomavryk/mavryk"
)

type PayoutCandidate struct {
	Source           mavryk.Address              `json:"source,omitempty"`
	Recipient        mavryk.Address              `json:"recipient,omitempty"`
	FeeRate          float64                    `json:"fee_rate,omitempty"`
	StakedBalance    mavryk.Z                    `json:"staked_balance,omitempty"`
	DelegatedBalance mavryk.Z                    `json:"delegated_balance,omitempty"`
	IsInvalid        bool                       `json:"is_invalid,omitempty"`
	IsEmptied        bool                       `json:"is_emptied,omitempty"`
	InvalidBecause   enums.EPayoutInvalidReason `json:"invalid_because,omitempty"`
}

func (candidate *PayoutCandidate) GetDelegatedBalance() mavryk.Z {
	return candidate.DelegatedBalance
}

func (candidate *PayoutCandidate) ToValidationContext(ctx *PayoutGenerationContext) PayoutValidationContext {
	pkh, _ := candidate.Recipient.MarshalText()
	var overrides *configuration.RuntimeDelegatorOverride
	if delegatorOverride, found := ctx.configuration.Delegators.Overrides[string(pkh)]; found {
		overrides = &delegatorOverride
	}
	return PayoutValidationContext{
		Configuration: ctx.configuration,
		Overrides:     overrides,
		Payout:        candidate,
		Ctx:           ctx,
	}
}

type PayoutCandidateWithBondAmount struct {
	PayoutCandidate
	BondsAmount mavryk.Z                     `json:"bonds_amount,omitempty"`
	TxKind      enums.EPayoutTransactionKind `json:"tx_kind,omitempty"`
	FATokenId   mavryk.Z                      `json:"fa_token_id,omitempty"` // required only if fa12 or fa2
	FAContract  mavryk.Address                `json:"fa_contract"`           // required only if fa12 or fa2
	FAAlias     string                       `json:"fa_alias,omitempty"`
	FADecimals  int                          `json:"fa_decimals,omitempty"`
}

func (candidate *PayoutCandidateWithBondAmount) GetDestination() mavryk.Address {
	return candidate.Recipient
}

func (candidate *PayoutCandidateWithBondAmount) GetTxKind() enums.EPayoutTransactionKind {
	return candidate.TxKind
}

func (candidate *PayoutCandidateWithBondAmount) GetFATokenId() mavryk.Z {
	return candidate.FATokenId
}

func (candidate *PayoutCandidateWithBondAmount) GetFAContract() mavryk.Address {
	return candidate.FAContract
}

func (candidate *PayoutCandidateWithBondAmount) GetAmount() mavryk.Z {
	return candidate.BondsAmount
}

func (candidate *PayoutCandidateWithBondAmount) GetFeeRate() float64 {
	return candidate.FeeRate
}

type PayoutCandidateWithBondAmountAndFee struct {
	PayoutCandidateWithBondAmount
	Fee mavryk.Z `json:"fee,omitempty"`
}

func (candidate *PayoutCandidateWithBondAmountAndFee) ToValidationContext(ctx *PayoutGenerationContext) PresimPayoutCandidateValidationContext {
	pkh, _ := candidate.Recipient.MarshalText()
	var overrides *configuration.RuntimeDelegatorOverride
	if delegatorOverride, found := ctx.configuration.Delegators.Overrides[string(pkh)]; found {
		overrides = &delegatorOverride
	}

	return PresimPayoutCandidateValidationContext{
		Configuration: ctx.configuration,
		Overrides:     overrides,
		Payout:        candidate,
	}
}

func (payout *PayoutCandidateWithBondAmountAndFee) ToPayoutRecipe(baker mavryk.Address, cycle int64, kind enums.EPayoutKind) common.PayoutRecipe {
	note := ""
	if payout.IsInvalid {
		note = string(payout.InvalidBecause)
	}

	return common.PayoutRecipe{
		Baker:            baker,
		Cycle:            cycle,
		Kind:             kind,
		TxKind:           payout.TxKind,
		Delegator:        payout.Source,
		Recipient:        payout.Recipient,
		DelegatedBalance: payout.DelegatedBalance,
		StakedBalance:    payout.StakedBalance,
		FATokenId:        payout.FATokenId,
		FAContract:       payout.FAContract,
		FAAlias:          payout.FAAlias,
		FADecimals:       payout.FADecimals,
		Amount:           payout.BondsAmount,
		FeeRate:          payout.FeeRate,
		Fee:              payout.Fee,
		Note:             note,
		IsValid:          !payout.IsInvalid,
	}
}

func DelegatorToPayoutCandidate(delegator common.Delegator, configuration *configuration.RuntimeConfiguration) PayoutCandidate {
	pkh, _ := delegator.Address.MarshalText()
	delegatorOverrides := configuration.Delegators.Overrides
	payoutFeeRate := configuration.PayoutConfiguration.Fee
	payoutRecipient := delegator.Address

	if delegatorOverride, ok := delegatorOverrides[string(pkh)]; ok {
		if !delegatorOverride.Recipient.Equal(mavryk.InvalidAddress) {
			payoutRecipient = delegatorOverride.Recipient
		}
		if delegatorOverride.Fee != nil {
			payoutFeeRate = *delegatorOverride.Fee
		}
		if delegatorOverride.MaximumBalance != nil && delegatorOverride.MaximumBalance.IsLess(delegator.DelegatedBalance) {
			delegator.DelegatedBalance = *delegatorOverride.MaximumBalance
		}
	}

	return PayoutCandidate{
		Source:           delegator.Address,
		Recipient:        payoutRecipient,
		FeeRate:          payoutFeeRate,
		DelegatedBalance: delegator.DelegatedBalance,
		StakedBalance:    delegator.StakedBalance,
		IsEmptied:        delegator.Emptied,
	}
}

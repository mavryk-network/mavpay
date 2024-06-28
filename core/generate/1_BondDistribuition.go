package generate

import (
	"github.com/tez-capital/tezpay/common"
	"github.com/tez-capital/tezpay/configuration"
	"github.com/tez-capital/tezpay/constants"
	"github.com/tez-capital/tezpay/constants/enums"
	"github.com/tez-capital/tezpay/extension"
	"github.com/tez-capital/tezpay/utils"
	"github.com/trilitech/tzgo/tezos"

	"github.com/samber/lo"
)

type AfterBondsDistributedHookData struct {
	Cycle      int64                           `json:"cycle"`
	Candidates []PayoutCandidateWithBondAmount `json:"candidates"`
}

func ExecuteAfterBondsDistributed(data *AfterBondsDistributedHookData) error {
	return extension.ExecuteHook(enums.EXTENSION_HOOK_AFTER_BONDS_DISTRIBUTED, "0.2", data)
}

func getBakerBondsAmount(cycleData *common.BakersCycleData, effectiveDelegatorsStakingBalance tezos.Z, configuration *configuration.RuntimeConfiguration) tezos.Z {
	bakerBalance := cycleData.GetBakerDelegatedBalance()
	totalRewards := cycleData.GetTotalRewards(configuration.PayoutConfiguration.PayoutMode)

	overdelegationLimit := cycleData.FrozenDepositLimit
	if overdelegationLimit.IsZero() {
		overdelegationLimit = bakerBalance
	}
	bakerAmount := totalRewards.Div64(constants.DELEGATION_CAPACITY_FACTOR)
	stakingBalance := effectiveDelegatorsStakingBalance.Add(bakerBalance)

	if !overdelegationLimit.Mul64(constants.DELEGATION_CAPACITY_FACTOR).Sub(stakingBalance).IsNeg() || !configuration.Overdelegation.IsProtectionEnabled { // not overdelegated
		bakerAmount = totalRewards.Mul(bakerBalance).Div(stakingBalance)
	}
	return bakerAmount
}

func DistributeBonds(ctx *PayoutGenerationContext, options *common.GeneratePayoutsOptions) (*PayoutGenerationContext, error) {
	configuration := ctx.GetConfiguration()
	logger := ctx.logger.With("phase", "distribute_bonds")

	logger.Debug("distributing bonds")

	candidates := ctx.StageData.PayoutCandidates
	effectiveStakingBalance := lo.Reduce(candidates, func(total tezos.Z, candidate PayoutCandidate, _ int) tezos.Z {
		// of all delegators, including invalids, except ignored and possibly excluding bellow minimum balance
		if candidate.IsInvalid {
			if candidate.InvalidBecause == enums.INVALID_DELEGATOR_IGNORED {
				return total
			}
			if ctx.configuration.Delegators.Requirements.BellowMinimumBalanceRewardDestination == enums.REWARD_DESTINATION_EVERYONE && candidate.InvalidBecause == enums.INVALID_DELEGATOR_LOW_BAlANCE {
				return total
			}
		}
		return total.Add(candidate.GetEffectiveBalance())
	}, tezos.NewZ(0))

	bakerBonds := getBakerBondsAmount(ctx.StageData.CycleData, effectiveStakingBalance, configuration)
	availableRewards := ctx.StageData.CycleData.GetTotalRewards(configuration.PayoutConfiguration.PayoutMode).Sub(bakerBonds)

	ctx.StageData.PayoutCandidatesWithBondAmount = lo.Map(candidates, func(candidate PayoutCandidate, _ int) PayoutCandidateWithBondAmount {
		if candidate.IsInvalid {
			return PayoutCandidateWithBondAmount{
				PayoutCandidate: candidate,
				BondsAmount:     tezos.Zero,
			}
		}
		return PayoutCandidateWithBondAmount{
			PayoutCandidate: candidate,
			BondsAmount:     availableRewards.Mul(candidate.GetEffectiveBalance()).Div(effectiveStakingBalance),
			TxKind:          enums.PAYOUT_TX_KIND_TEZ,
		}
	})

	bondsDonate := utils.GetZPortion(bakerBonds, configuration.IncomeRecipients.DonateBonds)
	ctx.StageData.BakerBondsAmount = bakerBonds.Sub(bondsDonate)
	ctx.StageData.DonateBondsAmount = bondsDonate

	hookData := &AfterBondsDistributedHookData{
		Cycle:      options.Cycle,
		Candidates: ctx.StageData.PayoutCandidatesWithBondAmount,
	}
	err := ExecuteAfterBondsDistributed(hookData)
	if err != nil {
		return ctx, err
	}
	ctx.StageData.PayoutCandidatesWithBondAmount = hookData.Candidates

	return ctx, nil
}

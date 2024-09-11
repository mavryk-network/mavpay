package generate

import (
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mavpay/extension"
	"github.com/mavryk-network/mavpay/utils"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/samber/lo"
)

type OnFeesCollectionHookData = struct {
	Cycle      int64                                 `json:"cycle"`
	Candidates []PayoutCandidateWithBondAmountAndFee `json:"candidates"`
}

func ExecuteOnFeesCollection(data *OnFeesCollectionHookData) error {
	return extension.ExecuteHook(enums.EXTENSION_HOOK_ON_FEES_COLLECTION, "0.2", data)
}

func CollectBakerFee(ctx *PayoutGenerationContext, options *common.GeneratePayoutsOptions) (*PayoutGenerationContext, error) {
	configuration := ctx.GetConfiguration()
	logger := ctx.logger.With("phase", "collect_baker_fee")
	logger.Debug("collecting baker fee")
	candidates := ctx.StageData.PayoutCandidatesWithBondAmount

	candidatesWithBondsAndFees := lo.Map(candidates, func(candidateWithBondsAmount PayoutCandidateWithBondAmount, _ int) PayoutCandidateWithBondAmountAndFee {
		if candidateWithBondsAmount.IsInvalid {
			return PayoutCandidateWithBondAmountAndFee{
				PayoutCandidateWithBondAmount: candidateWithBondsAmount,
			}
		}

		if candidateWithBondsAmount.TxKind != enums.PAYOUT_TX_KIND_TEZ {
			logger.Debug("skipping fee collection for non mavryk payout", "delegate", candidateWithBondsAmount.Source, "tx_kind", candidateWithBondsAmount.TxKind)
			return PayoutCandidateWithBondAmountAndFee{
				PayoutCandidateWithBondAmount: candidateWithBondsAmount,
			}
		}

		fee := utils.GetZPortion(candidateWithBondsAmount.BondsAmount, candidateWithBondsAmount.FeeRate)
		candidateWithBondsAmount.BondsAmount = candidateWithBondsAmount.BondsAmount.Sub(fee)
		if candidateWithBondsAmount.BondsAmount.IsZero() || candidateWithBondsAmount.BondsAmount.IsNeg() {
			candidateWithBondsAmount.IsInvalid = true
			candidateWithBondsAmount.InvalidBecause = enums.INVALID_PAYOUT_BELLOW_MINIMUM
		}
		return PayoutCandidateWithBondAmountAndFee{
			PayoutCandidateWithBondAmount: candidateWithBondsAmount,
			Fee:                           fee,
		}
	})

	hookData := &OnFeesCollectionHookData{
		Cycle:      options.Cycle,
		Candidates: candidatesWithBondsAndFees,
	}
	err := ExecuteOnFeesCollection(hookData)
	if err != nil {
		return ctx, err
	}
	candidatesWithBondsAndFees = hookData.Candidates

	collectedFees := lo.Reduce(candidatesWithBondsAndFees, func(agg mavryk.Z, candidateWithBondsAmountAndFee PayoutCandidateWithBondAmountAndFee, _ int) mavryk.Z {
		return agg.Add(candidateWithBondsAmountAndFee.Fee)
	}, mavryk.Zero)

	feesDonate := utils.GetZPortion(collectedFees, configuration.IncomeRecipients.DonateFees)
	ctx.StageData.BakerFeesAmount = collectedFees.Sub(feesDonate)
	ctx.StageData.DonateFeesAmount = feesDonate
	ctx.StageData.PayoutCandidatesWithBondAmountAndFees = candidatesWithBondsAndFees

	return ctx, nil
}

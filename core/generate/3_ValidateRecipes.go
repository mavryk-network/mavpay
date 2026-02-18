package generate

import (
	"github.com/samber/lo"
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/utils"
	"github.com/mavryk-network/gomavryk/mavryk"
)

func ValidateRecipe(ctx *PayoutGenerationContext, options *common.GeneratePayoutsOptions) (result *PayoutGenerationContext, err error) {
	logger := ctx.logger.With("phase", "validate_simulated_payouts")
	logger.Info("validating payout candidates")

	ctx.StageData.PayoutCandidatesWithBondAmountAndFees = lo.Map(ctx.StageData.PayoutCandidatesWithBondAmountAndFees, func(candidate PayoutCandidateWithBondAmountAndFee, _ int) PayoutCandidateWithBondAmountAndFee {
		validationContext := candidate.ToValidationContext(ctx)
		result := *validationContext.Validate(
			TxKindValidator,
			MinumumAmountValidator,
		).ToPresimPayoutCandidate()

		utils.AssertZAmountPositiveOrZero(candidate.BondsAmount)
		if candidate.IsInvalid {
			candidate.Fee = candidate.Fee.Add(candidate.BondsAmount) // we need to add because we already collected fees from bonds in 2_CollectBakerFee.go
			candidate.BondsAmount = mavryk.Zero
		}
		return result
	})

	return ctx, nil
}

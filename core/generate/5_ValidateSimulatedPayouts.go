package generate

import (
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/samber/lo"
)

func ValidateSimulatedPayouts(ctx *PayoutGenerationContext, options *common.GeneratePayoutsOptions) (result *PayoutGenerationContext, err error) {
	configuration := ctx.GetConfiguration()
	logger := ctx.logger.With("phase", "validate_simulated_payouts")
	simulated := ctx.StageData.PayoutCandidatesSimulated

	logger.Info("validating simulated payout candidates")

	ctx.StageData.PayoutCandidatesSimulated = lo.Map(simulated, func(candidate PayoutCandidateSimulated, _ int) PayoutCandidateSimulated {
		if candidate.IsInvalid {
			return candidate
		}

		validationContext := candidate.ToValidationContext(configuration)
		result := *validationContext.Validate(
			MinumumAmountSimulatedValidator,
		).ToPayoutCandidateSimulated()

		// collect fees if invalid
		if candidate.IsInvalid {
			ctx.StageData.BakerFeesAmount = ctx.StageData.BakerFeesAmount.Add(candidate.BondsAmount)
			candidate.Fee = candidate.Fee.Add(candidate.BondsAmount)
			candidate.BondsAmount = mavryk.Zero
		}
		return result
	})

	return ctx, nil
}

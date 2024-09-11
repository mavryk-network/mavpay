package core

import (
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/core/generate"
)

const (
	PAYOUT_EXECUTION_FAILURE = iota
	PAYOUT_EXECUTION_SUCCESS
)

func GeneratePayouts(config *configuration.RuntimeConfiguration, engineContext *common.GeneratePayoutsEngineContext, options *common.GeneratePayoutsOptions) (*common.CyclePayoutBlueprint, error) {
	if config == nil {
		return nil, constants.ErrMissingConfiguration
	}

	ctx, err := generate.NewPayoutGenerationContext(config, engineContext)
	if err != nil {
		return nil, err
	}

	ctx, err = WrapContext[*generate.PayoutGenerationContext, *common.GeneratePayoutsOptions](ctx).ExecuteStages(options,
		generate.SendAnalytics,
		generate.CheckConditionsAndPrepare,
		generate.GeneratePayoutCandidates,
		// hooks
		generate.DistributeBonds,
		generate.CollectBakerFee,
		generate.CheckSufficientBalance,
		generate.CollectTransactionFees,
		generate.ValidateSimulatedPayouts,
		generate.FinalizePayouts,
		generate.CreateBlueprint).Unwrap()
	return ctx.StageData.PayoutBlueprint, err
}

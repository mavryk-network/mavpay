package core

import (
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/core/execute"
)

func ExecutePayouts(preparationResult *common.PreparePayoutsResult, config *configuration.RuntimeConfiguration, engineContext *common.ExecutePayoutsEngineContext, options *common.ExecutePayoutsOptions) (*common.ExecutePayoutsResult, error) {
	if config == nil {
		return nil, constants.ErrMissingConfiguration
	}

	ctx, err := execute.NewPayoutExecutionContext(preparationResult, config, engineContext, options)
	if err != nil {
		return nil, err
	}

	ctx, err = WrapContext[*execute.PayoutExecutionContext, *common.ExecutePayoutsOptions](ctx).ExecuteStages(options,
		execute.SplitIntoBatches,
		execute.ExecutePayouts).Unwrap()
	return &common.ExecutePayoutsResult{
		BatchResults:   ctx.StageData.BatchResults,
		PaidDelegators: ctx.StageData.PaidDelegators,
	}, err
}

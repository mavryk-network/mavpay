package core

import (
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/core/prepare"
)

func PreparePayouts(blueprints []*common.CyclePayoutBlueprint, config *configuration.RuntimeConfiguration, engineContext *common.PreparePayoutsEngineContext, options *common.PreparePayoutsOptions) (*common.PreparePayoutsResult, error) {
	if config == nil {
		return nil, constants.ErrMissingConfiguration
	}

	ctx, err := prepare.NewPayoutPreparationContext(blueprints, config, engineContext, options)
	if err != nil {
		return nil, err
	}

	ctx, err = WrapContext[*prepare.PayoutPrepareContext, *common.PreparePayoutsOptions](ctx).ExecuteStages(options,
		prepare.PreparePayouts,
		prepare.AccumulatePayouts).Unwrap()
	return &common.PreparePayoutsResult{
		Blueprints:                    ctx.PayoutBlueprints,
		ValidPayouts:                  ctx.StageData.ValidPayouts,
		AccumulatedPayouts:            ctx.StageData.AccumulatedPayouts,
		InvalidPayouts:                ctx.StageData.InvalidPayouts,
		ReportsOfPastSuccesfulPayouts: ctx.StageData.ReportsOfPastSuccesfulPayouts,
	}, err
}

func PrepareCyclePayouts(blueprint *common.CyclePayoutBlueprint, config *configuration.RuntimeConfiguration, engineContext *common.PreparePayoutsEngineContext, options *common.PreparePayoutsOptions) (*common.PreparePayoutsResult, error) {
	return PreparePayouts([]*common.CyclePayoutBlueprint{blueprint}, config, engineContext, options)
}

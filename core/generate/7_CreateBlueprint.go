package generate

import (
	"time"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mavpay/extension"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/samber/lo"
)

func sumValidPayoutsAmount(payouts []common.PayoutRecipe) mavryk.Z {
	return lo.Reduce(payouts, func(agg mavryk.Z, payout common.PayoutRecipe, _ int) mavryk.Z {
		if !payout.IsValid {
			return agg
		}
		return agg.Add(payout.Amount)
	}, mavryk.Zero)
}

type AfterPayoutsBlueprintGeneratedHookData = common.CyclePayoutBlueprint

// NOTE: do we want to allow rewriting of blueprint?
func ExecuteAfterPayoutsBlueprintGenerated(data AfterPayoutsBlueprintGeneratedHookData) error {
	return extension.ExecuteHook(enums.EXTENSION_HOOK_AFTER_PAYOUTS_BLUEPRINT_GENERATED, "0.1", &data)
}

func CreateBlueprint(ctx *PayoutGenerationContext, options *common.GeneratePayoutsOptions) (result *PayoutGenerationContext, err error) {
	stageData := ctx.StageData
	logger := ctx.logger.With("phase", "create_blueprint")
	logger.Info("creating payout blueprint")

	blueprint := common.CyclePayoutBlueprint{
		Cycle:   options.Cycle,
		Payouts: stageData.Payouts,
		Summary: common.CyclePayoutSummary{
			Cycle:                    options.Cycle,
			Delegators:               len(stageData.CycleData.Delegators),
			PaidDelegators:           stageData.PaidDelegators,
			OwnStakedBalance:         stageData.CycleData.OwnStakedBalance,
			OwnDelegatedBalance:      stageData.CycleData.OwnDelegatedBalance,
			ExternalStakedBalance:    stageData.CycleData.ExternalStakedBalance,
			ExternalDelegatedBalance: stageData.CycleData.ExternalDelegatedBalance,
			EarnedFees:               stageData.CycleData.BlockDelegatedFees,
			EarnedRewards:            stageData.CycleData.GetTotalDelegatedRewards(ctx.configuration.PayoutConfiguration.PayoutMode),
			DistributedRewards:       sumValidPayoutsAmount(stageData.Payouts),
			BondIncome:               stageData.BakerBondsAmount,
			FeeIncome:                stageData.BakerFeesAmount,
			IncomeTotal:              stageData.BakerBondsAmount.Add(stageData.BakerFeesAmount),
			DonatedBonds:             stageData.DonateBondsAmount,
			DonatedFees:              stageData.DonateFeesAmount,
			DonatedTotal:             stageData.DonateFeesAmount.Add(stageData.DonateBondsAmount),
			Timestamp:                time.Now(),
		},
		BatchMetadataDeserializationGasLimit: stageData.BatchMetadataDeserializationGasLimit,
	}

	err = ExecuteAfterPayoutsBlueprintGenerated(blueprint)
	if err != nil {
		return ctx, err
	}

	stageData.PayoutBlueprint = &blueprint
	return ctx, nil
}

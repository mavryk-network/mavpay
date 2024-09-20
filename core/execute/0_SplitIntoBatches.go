package execute

import (
	"errors"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mavpay/utils"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/samber/lo"
)

func splitIntoBatches(payouts []common.PayoutRecipe, limits *common.OperationLimits, metadataDeserializationGasLimit int64) ([]common.RecipeBatch, error) {
	batches := make([]common.RecipeBatch, 0)
	batchBlueprint := common.NewBatch(limits, metadataDeserializationGasLimit)

	for _, payout := range payouts {
		if !batchBlueprint.AddPayout(payout) {
			batches = append(batches, batchBlueprint.ToBatch())
			batchBlueprint = common.NewBatch(limits, metadataDeserializationGasLimit)
			if !batchBlueprint.AddPayout(payout) {
				return nil, constants.ErrPayoutDidNotFitTheBatch
			}
		}
	}
	// append last
	batches = append(batches, batchBlueprint.ToBatch())

	return lo.Filter(batches, func(batch common.RecipeBatch, _ int) bool {
		return len(batch) > 0
	}), nil
}

func SplitIntoBatches(ctx *PayoutExecutionContext, options *common.ExecutePayoutsOptions) (*PayoutExecutionContext, error) {
	logger := ctx.logger.With("phase", "split_into_batches")
	logger.Info("splitting into batches")
	var err error
	ctx.StageData.Limits, err = ctx.GetTransactor().GetLimits()
	if err != nil {
		return nil, errors.Join(constants.ErrGetChainLimitsFailed, err)
	}
	payouts := ctx.ValidPayouts
	payoutsWithoutFa := utils.RejectPayoutsByTxKind(payouts, enums.FA_OPERATION_KINDS)

	faRecipes := utils.FilterPayoutsByTxKind(payouts, enums.FA_OPERATION_KINDS)
	contractMavRecipes := utils.FilterPayoutsByType(payoutsWithoutFa, mavryk.AddressTypeContract)
	classicMavRecipes := utils.RejectPayoutsByType(payoutsWithoutFa, mavryk.AddressTypeContract)

	toBatch := make([][]common.PayoutRecipe, 0, 3)
	if options.MixInFATransfers {
		classicMavRecipes = append(classicMavRecipes, faRecipes...)
	} else {
		toBatch = append(toBatch, faRecipes)
	}
	if options.MixInContractCalls {
		classicMavRecipes = append(classicMavRecipes, contractMavRecipes...)
	} else {
		toBatch = append(toBatch, contractMavRecipes)
	}
	toBatch = append(toBatch, classicMavRecipes)

	batchMetadataDeserializationGasLimit := lo.Reduce(ctx.PayoutBlueprints, func(agg int64, blueprint *common.CyclePayoutBlueprint, _ int) int64 {
		return max(agg, blueprint.BatchMetadataDeserializationGasLimit)
	}, 0)

	stageBatches := make([]common.RecipeBatch, 0)
	for _, batch := range toBatch {
		batches, err := splitIntoBatches(batch, ctx.StageData.Limits, batchMetadataDeserializationGasLimit)
		if err != nil {
			return nil, err
		}
		stageBatches = append(stageBatches, batches...)
	}

	ctx.StageData.Batches = stageBatches
	return ctx, nil
}

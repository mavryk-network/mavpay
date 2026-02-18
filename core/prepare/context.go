package prepare

import (
	"log/slog"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mvgo/mavryk"
)

type StageData struct {
	Payouts                       []common.PayoutRecipe
	AccumulatedPayouts            []*common.AccumulatedPayoutRecipe
	InvalidRecipes                []common.PayoutRecipe
	ReportsOfPastSuccesfulPayouts []common.PayoutReport
	// protocol, signature etc.
	BatchMetadataDeserializationGasLimit int64
}

type PayoutPrepareContext struct {
	common.PreparePayoutsEngineContext
	configuration *configuration.RuntimeConfiguration

	StageData *StageData

	PayoutBlueprints []*common.CyclePayoutBlueprint

	PayoutKey mavryk.Key

	logger *slog.Logger
}

func (ctx *PayoutPrepareContext) GetConfiguration() *configuration.RuntimeConfiguration {
	return ctx.configuration
}

func NewPayoutPreparationContext(blueprints []*common.CyclePayoutBlueprint, configuration *configuration.RuntimeConfiguration, engineContext *common.PreparePayoutsEngineContext, options *common.PreparePayoutsOptions) (*PayoutPrepareContext, error) {
	if err := engineContext.Validate(); err != nil {
		return nil, err
	}

	return &PayoutPrepareContext{
		PreparePayoutsEngineContext: *engineContext,
		configuration:               configuration,

		StageData: &StageData{},

		PayoutBlueprints: blueprints,
		PayoutKey:        engineContext.GetSigner().GetKey(),

		logger: slog.Default().With("stage", "prepare"),
	}, nil
}

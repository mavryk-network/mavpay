package generate

import (
	"log/slog"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mvgo/mavryk"
)

type StageData struct {
	CycleData                             *common.BakersCycleData
	PayoutCandidates                      []PayoutCandidate
	PayoutCandidatesWithBondAmount        []PayoutCandidateWithBondAmount
	PayoutCandidatesWithBondAmountAndFees []PayoutCandidateWithBondAmountAndFee
	PayoutCandidatesSimulated             []PayoutCandidateSimulated
	PayoutBlueprint                       *common.CyclePayoutBlueprint

	Payouts           []common.PayoutRecipe
	BakerBondsAmount  mavryk.Z
	DonateBondsAmount mavryk.Z
	BakerFeesAmount   mavryk.Z
	DonateFeesAmount  mavryk.Z
	PaidDelegators    int

	// protocol, signature etc.
	BatchMetadataDeserializationGasLimit int64
}

type PayoutGenerationContext struct {
	common.GeneratePayoutsEngineContext
	configuration *configuration.RuntimeConfiguration

	StageData *StageData

	PayoutKey mavryk.Key

	logger *slog.Logger
}

func NewPayoutGenerationContext(configuration *configuration.RuntimeConfiguration, engineContext *common.GeneratePayoutsEngineContext) (*PayoutGenerationContext, error) {
	slog.Debug("mavpay payout context initialization")
	if err := engineContext.Validate(); err != nil {
		return nil, err
	}

	ctx := PayoutGenerationContext{
		GeneratePayoutsEngineContext: *engineContext,
		configuration:                configuration,

		StageData: &StageData{},

		PayoutKey: engineContext.GetSigner().GetKey(),

		logger: slog.Default().With("stage", "generate"),
	}

	return &ctx, nil
}

func (ctx *PayoutGenerationContext) GetConfiguration() *configuration.RuntimeConfiguration {
	return ctx.configuration
}

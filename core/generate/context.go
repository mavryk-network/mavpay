package generate

import (
	"log/slog"

	"blockwatch.cc/tzgo/tezos"
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
)

type StageData struct {
	CycleData                             *common.BakersCycleData
	PayoutCandidates                      []PayoutCandidate
	PayoutCandidatesWithBondAmount        []PayoutCandidateWithBondAmount
	PayoutCandidatesWithBondAmountAndFees []PayoutCandidateWithBondAmountAndFee
	PayoutCandidatesSimulated             []PayoutCandidateSimulated
	PayoutBlueprint                       *common.CyclePayoutBlueprint

	Payouts           []common.PayoutRecipe
	BakerBondsAmount  tezos.Z
	DonateBondsAmount tezos.Z
	BakerFeesAmount   tezos.Z
	DonateFeesAmount  tezos.Z
	PaidDelegators    int

	// protocol, signature etc.
	BatchMetadataDeserializationGasLimit int64
}

type PayoutGenerationContext struct {
	common.GeneratePayoutsEngineContext
	configuration *configuration.RuntimeConfiguration

	StageData *StageData

	PayoutKey tezos.Key

	logger *slog.Logger
}

func NewPayoutGenerationContext(configuration *configuration.RuntimeConfiguration, engineContext *common.GeneratePayoutsEngineContext) (*PayoutGenerationContext, error) {
	slog.Debug("tezpay payout context initialization")
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

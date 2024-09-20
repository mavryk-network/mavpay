package common

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/samber/lo"
)

type OpLimits struct {
	TransactionFee          int64 `json:"transaction_fee,omitempty"`
	StorageLimit            int64 `json:"storage_limit,omitempty"`
	GasLimit                int64 `json:"gas_limit,omitempty"`
	DeserializationGasLimit int64 `json:"deserialization_gas_limit,omitempty"`
	AllocationBurn          int64 `json:"allocation_burn,omitempty"`
	StorageBurn             int64 `json:"storage_burn,omitempty"`
}

func (psr *OpLimits) GetOperationTotalFees() int64 {
	return psr.TransactionFee + psr.AllocationBurn + psr.StorageBurn
}

func (psr *OpLimits) GetAllocationFee() int64 {
	return psr.AllocationBurn
}

func (psr *OpLimits) GetOperationFeesWithoutAllocation() int64 {
	return psr.TransactionFee + psr.StorageBurn
}

type PayoutRecipe struct {
	Baker            mavryk.Address               `json:"baker"`
	Delegator        mavryk.Address               `json:"delegator,omitempty"`
	Cycle            int64                        `json:"cycle,omitempty"`
	Recipient        mavryk.Address               `json:"recipient,omitempty"`
	Kind             enums.EPayoutKind            `json:"kind,omitempty"`
	TxKind           enums.EPayoutTransactionKind `json:"tx_kind,omitempty"`
	FATokenId        mavryk.Z                     `json:"fa_token_id,omitempty"`
	FAContract       mavryk.Address               `json:"fa_contract,omitempty"`
	DelegatedBalance mavryk.Z                     `json:"delegator_balance,omitempty"`
	StakedBalance    mavryk.Z                     `json:"-"` // enable in output when relevant (P)
	Amount           mavryk.Z                     `json:"amount,omitempty"`
	FeeRate          float64                      `json:"fee_rate,omitempty"`
	Fee              mavryk.Z                     `json:"fee,omitempty"`
	OpLimits         *OpLimits                    `json:"op_limits,omitempty"`
	Note             string                       `json:"note,omitempty"`
	IsValid          bool                         `json:"valid,omitempty"`
	// mainly for accumulation to be able to check if fee was collected and subtract it from the amount
	TxFeeCollected bool `json:"tx_fee_collected,omitempty"`
	// mainly for accumulation to be able to check if fee was collected and subtract it from the amount
	AllocationFeeCollected bool `json:"allocation_fee_collected,omitempty"`
}

func (candidate *PayoutRecipe) GetDestination() mavryk.Address {
	return candidate.Recipient
}

func (candidate *PayoutRecipe) GetTxKind() enums.EPayoutTransactionKind {
	return candidate.TxKind
}

func (candidate *PayoutRecipe) GetFATokenId() mavryk.Z {
	return candidate.FATokenId
}

func (candidate *PayoutRecipe) GetFAContract() mavryk.Address {
	return candidate.FAContract
}

func (candidate *PayoutRecipe) GetAmount() mavryk.Z {
	return candidate.Amount
}

type PayoutRecipeIdentifier struct {
	Delegator  mavryk.Address               `json:"delegator,omitempty"`
	Recipient  mavryk.Address               `json:"recipient,omitempty"`
	Kind       enums.EPayoutKind            `json:"kind,omitempty"`
	TxKind     enums.EPayoutTransactionKind `json:"tx_kind,omitempty"`
	FATokenId  mavryk.Z                     `json:"fa_token_id,omitempty"`
	FAContract mavryk.Address               `json:"fa_contract,omitempty"`
	IsValid    bool                         `json:"valid,omitempty"`
}

func (identifier *PayoutRecipeIdentifier) ToJSON() ([]byte, error) {
	return json.Marshal(identifier)
}

func (recipe *PayoutRecipe) GetIdentifier() string {
	identifier := PayoutRecipeIdentifier{
		Delegator:  recipe.Delegator,
		Recipient:  recipe.Recipient,
		Kind:       recipe.Kind,
		TxKind:     recipe.TxKind,
		FATokenId:  recipe.FATokenId,
		FAContract: recipe.FAContract,
		IsValid:    recipe.IsValid,
	}
	k, err := identifier.ToJSON()
	if err != nil {
		return ""
	}
	hashBytes := sha256.Sum256(k)
	hash := hex.EncodeToString(hashBytes[:])
	return hash
}

func (recipe *PayoutRecipe) GetShortIdentifier() string {
	return recipe.GetIdentifier()[:16]
}

func (recipe *PayoutRecipe) Combine(otherRecipe *PayoutRecipe) (*PayoutRecipe, error) {
	if !recipe.Recipient.Equal(otherRecipe.Recipient) {
		return nil, errors.New("cannot combine different recipients")
	}
	if !recipe.Delegator.Equal(otherRecipe.Delegator) {
		return nil, errors.New("cannot combine different delegators")
	}
	if recipe.Kind != otherRecipe.Kind {
		return nil, errors.New("cannot combine different kinds")
	}
	if recipe.TxKind != otherRecipe.TxKind {
		return nil, errors.New("cannot combine different tx kinds")
	}
	if !recipe.FATokenId.Equal(otherRecipe.FATokenId) {
		return nil, errors.New("cannot combine different FA token ids")
	}
	if !recipe.FAContract.Equal(otherRecipe.FAContract) {
		return nil, errors.New("cannot combine different FA contracts")
	}
	if recipe.IsValid != otherRecipe.IsValid {
		return nil, errors.New("cannot combine different validities")
	}
	if recipe.OpLimits == nil || otherRecipe.OpLimits == nil {
		return nil, errors.New("cannot combine recipes with missing op limits")
	}

	recipe.DelegatedBalance = recipe.DelegatedBalance.Add(otherRecipe.DelegatedBalance).Div64(2)
	recipe.StakedBalance = recipe.StakedBalance.Add(otherRecipe.StakedBalance).Div64(2)
	recipe.Amount = recipe.Amount.Add(otherRecipe.Amount)
	recipe.Fee = recipe.Fee.Add(otherRecipe.Fee)
	recipe.OpLimits = &OpLimits{
		StorageBurn:             recipe.OpLimits.StorageBurn + otherRecipe.OpLimits.StorageBurn,
		AllocationBurn:          recipe.OpLimits.AllocationBurn + otherRecipe.OpLimits.AllocationBurn,
		TransactionFee:          recipe.OpLimits.TransactionFee + otherRecipe.OpLimits.TransactionFee,
		StorageLimit:            recipe.OpLimits.StorageLimit + otherRecipe.OpLimits.StorageLimit,
		GasLimit:                recipe.OpLimits.GasLimit + otherRecipe.OpLimits.GasLimit,
		DeserializationGasLimit: recipe.OpLimits.DeserializationGasLimit + otherRecipe.OpLimits.DeserializationGasLimit,
	}

	otherRecipe.Kind = enums.PAYOUT_KIND_ACCUMULATED
	otherRecipe.Note = fmt.Sprintf("%s#%d", recipe.GetShortIdentifier(), recipe.Cycle)

	return recipe, nil
}

func (pr *PayoutRecipe) GetAccumulatedIdentifier() string {
	return fmt.Sprintf("%s#%d", pr.GetShortIdentifier(), pr.Cycle)
}

func (pr *PayoutRecipe) GetAccumulatedPayoutDetails() (wasAccumulated bool, id string, cycle int64) {
	if pr.Kind != enums.PAYOUT_KIND_ACCUMULATED {
		return false, "", 0
	}
	if len(pr.Note) > 0 {
		_, err := fmt.Sscanf(pr.Note, "%s#%d", &id, &cycle)
		if err == nil {
			return true, id, cycle
		}
	}

	return false, "", 0
}

func (pr *PayoutRecipe) ToPayoutReport() PayoutReport {
	txFee := int64(0)
	if pr.OpLimits != nil {
		txFee = pr.OpLimits.TransactionFee
	}

	return PayoutReport{
		Id:               pr.GetShortIdentifier(),
		Baker:            pr.Baker,
		Timestamp:        time.Now(),
		Cycle:            pr.Cycle,
		Kind:             pr.Kind,
		TxKind:           pr.TxKind,
		FAContract:       pr.FAContract,
		FATokenId:        pr.FATokenId,
		Delegator:        pr.Delegator,
		DelegatedBalance: pr.DelegatedBalance,
		StakedBalance:    pr.StakedBalance,
		Recipient:        pr.Recipient,
		Amount:           pr.Amount,
		FeeRate:          pr.FeeRate,
		Fee:              pr.Fee,
		TransactionFee:   txFee,
		OpHash:           mavryk.ZeroOpHash,
		IsSuccess:        false,
		Note:             pr.Note,
	}
}

func (pr *PayoutRecipe) GetTransactionFee() int64 {
	if pr.OpLimits != nil {
		return pr.OpLimits.TransactionFee
	}
	return 0
}

func (pr *PayoutRecipe) ToTableRowData() []string {
	return []string{
		ShortenAddress(pr.Delegator),
		ShortenAddress(pr.Recipient),
		MumavToMavS(pr.DelegatedBalance.Int64()),
		string(pr.Kind),
		ShortenAddress(pr.FAContract),
		ToStringEmptyIfZero(pr.FATokenId.Int64()),
		FormatAmount(pr.TxKind, pr.Amount.Int64()),
		FloatToPercentage(pr.FeeRate),
		MumavToMavS(pr.Fee.Int64()),
		MumavToMavS(pr.GetTransactionFee()),
		pr.Note,
	}
}

func (pr *PayoutRecipe) GetTableHeaders() []string {
	return []string{
		"Delegator",
		"Recipient",
		"Delegated Balance",
		"Kind",
		"FA Contract",
		"FA Token Id",
		"Amount",
		"Fee Rate",
		"Fee",
		"Tx Fee",
		"Note",
	}
}

func GetRecipesTotals(recipes []PayoutRecipe) []string {
	totalAmount := int64(0)
	totalFee := int64(0)
	totalTx := int64(0)
	for _, recipe := range recipes {
		if recipe.TxKind == enums.PAYOUT_TX_KIND_MAV {
			totalAmount += recipe.Amount.Int64()
		}
		totalFee += recipe.Fee.Int64()
		totalTx += recipe.GetTransactionFee()
	}
	return []string{
		"",
		"",
		"",
		"",
		"",
		"",
		MumavToMavS(totalAmount),
		"",
		MumavToMavS(totalFee),
		MumavToMavS(totalTx),
		"",
	}
}

// returns totals and number of filtered recipes
func GetRecipesFilteredTotals(recipes []PayoutRecipe, kind enums.EPayoutKind) ([]string, int) {
	r := lo.Filter(recipes, func(recipe PayoutRecipe, _ int) bool {
		return recipe.Kind == kind
	})
	return GetRecipesTotals(r), len(r)
}

type CyclePayoutSummary struct {
	Cycle                    int64     `json:"cycle"`
	Delegators               int       `json:"delegators"`
	PaidDelegators           int       `json:"paid_delegators"`
	OwnStakedBalance         mavryk.Z  `json:"own_staked_balance"`
	OwnDelegatedBalance      mavryk.Z  `json:"own_delegated_balance"`
	ExternalStakedBalance    mavryk.Z  `json:"external_staked_balance"`
	ExternalDelegatedBalance mavryk.Z  `json:"external_delegated_balance"`
	EarnedFees               mavryk.Z  `json:"cycle_fees"`
	EarnedRewards            mavryk.Z  `json:"cycle_rewards"`
	DistributedRewards       mavryk.Z  `json:"distributed_rewards"`
	BondIncome               mavryk.Z  `json:"bond_income"`
	FeeIncome                mavryk.Z  `json:"fee_income"`
	IncomeTotal              mavryk.Z  `json:"total_income"`
	DonatedBonds             mavryk.Z  `json:"donated_bonds"`
	DonatedFees              mavryk.Z  `json:"donated_fees"`
	DonatedTotal             mavryk.Z  `json:"donated_total"`
	Timestamp                time.Time `json:"timestamp"`
}

func (summary *CyclePayoutSummary) GetTotalStakedBalance() mavryk.Z {
	return summary.OwnStakedBalance.Add(summary.ExternalStakedBalance)
}

func (summary *CyclePayoutSummary) GetTotalDelegatedBalance() mavryk.Z {
	return summary.OwnDelegatedBalance.Add(summary.ExternalDelegatedBalance)
}

func (summary *CyclePayoutSummary) CombineNumericData(another *CyclePayoutSummary) *CyclePayoutSummary {
	return &CyclePayoutSummary{
		OwnStakedBalance:         summary.OwnStakedBalance.Add(another.OwnStakedBalance),
		OwnDelegatedBalance:      summary.OwnDelegatedBalance.Add(another.OwnDelegatedBalance),
		ExternalStakedBalance:    summary.ExternalStakedBalance.Add(another.ExternalStakedBalance),
		ExternalDelegatedBalance: summary.ExternalDelegatedBalance.Add(another.ExternalDelegatedBalance),
		EarnedFees:               summary.EarnedFees.Add(another.EarnedFees),
		EarnedRewards:            summary.EarnedRewards.Add(another.EarnedRewards),
		DistributedRewards:       summary.DistributedRewards.Add(another.DistributedRewards),
		BondIncome:               summary.BondIncome.Add(another.BondIncome),
		FeeIncome:                summary.FeeIncome.Add(another.FeeIncome),
		IncomeTotal:              summary.IncomeTotal.Add(another.IncomeTotal),
		DonatedBonds:             summary.DonatedBonds.Add(another.DonatedBonds),
		DonatedFees:              summary.DonatedFees.Add(another.DonatedFees),
		DonatedTotal:             summary.DonatedTotal.Add(another.DonatedTotal),
	}
}

type CyclePayoutBlueprint struct {
	Cycle                                int64              `json:"cycles,omitempty"`
	Payouts                              []PayoutRecipe     `json:"payouts,omitempty"`
	Summary                              CyclePayoutSummary `json:"summary,omitempty"`
	BatchMetadataDeserializationGasLimit int64              `json:"batch_metadata_deserialization_gas_limit,omitempty"`
}

type GeneratePayoutsEngineContext struct {
	collector   CollectorEngine
	signer      SignerEngine
	adminNotify func(msg string)
}

func NewGeneratePayoutsEngines(collector CollectorEngine, signer SignerEngine, adminNotify func(msg string)) *GeneratePayoutsEngineContext {
	return &GeneratePayoutsEngineContext{
		collector:   collector,
		signer:      signer,
		adminNotify: adminNotify,
	}
}

func (engines *GeneratePayoutsEngineContext) GetSigner() SignerEngine {
	return engines.signer
}

func (engines *GeneratePayoutsEngineContext) GetCollector() CollectorEngine {
	return engines.collector
}

func (engines *GeneratePayoutsEngineContext) AdminNotify(msg string) {
	if engines.adminNotify != nil {
		engines.adminNotify(msg)
	}
}

func (engines *GeneratePayoutsEngineContext) Validate() error {
	if engines.signer == nil {
		return errors.Join(constants.ErrMissingEngine, constants.ErrMissingSignerEngine)
	}
	if engines.collector == nil {
		return errors.Join(constants.ErrMissingEngine, constants.ErrMissingCollectorEngine)
	}
	return nil
}

type GeneratePayoutsOptions struct {
	Cycle                    int64 `json:"cycle,omitempty"`
	SkipBalanceCheck         bool  `json:"skip_balance_check,omitempty"`
	WaitForSufficientBalance bool  `json:"wait_for_sufficient_balance,omitempty"`
}

type CyclePayoutBlueprints []*CyclePayoutBlueprint

func (results CyclePayoutBlueprints) GetSummary() *CyclePayoutSummary {
	summary := &CyclePayoutSummary{}
	delegators := 0
	for _, result := range results {
		delegators += result.Summary.Delegators
		summary = summary.CombineNumericData(&result.Summary)
	}
	summary.Delegators = delegators / len(results) // average
	return summary
}

type PreparePayoutsEngineContext struct {
	collector   CollectorEngine
	signer      SignerEngine
	reporter    ReporterEngine
	adminNotify func(msg string)
}

func NewPreparePayoutsEngineContext(collector CollectorEngine, signer SignerEngine, reporter ReporterEngine, adminNotify func(msg string)) *PreparePayoutsEngineContext {
	return &PreparePayoutsEngineContext{
		collector:   collector,
		adminNotify: adminNotify,
		signer:      signer,
		reporter:    reporter,
	}
}

func (engines *PreparePayoutsEngineContext) GetCollector() CollectorEngine {
	return engines.collector
}

func (engines *PreparePayoutsEngineContext) GetSigner() SignerEngine {
	return engines.signer
}

func (engines *PreparePayoutsEngineContext) GetReporter() ReporterEngine {
	return engines.reporter
}

func (engines *PreparePayoutsEngineContext) AdminNotify(msg string) {
	if engines.adminNotify != nil {
		engines.adminNotify(msg)
	}
}

func (engines *PreparePayoutsEngineContext) Validate() error {
	if engines.collector == nil {
		return errors.Join(constants.ErrMissingEngine, constants.ErrMissingCollectorEngine)
	}
	if engines.reporter == nil {
		return errors.Join(constants.ErrMissingEngine, constants.ErrMissingReporterEngine)
	}
	return nil
}

type PreparePayoutsOptions struct {
	Accumulate bool `json:"accumulate,omitempty"`
}

type PreparePayoutsResult struct {
	Blueprints                    []*CyclePayoutBlueprint `json:"blueprint,omitempty"`
	ValidPayouts                  []PayoutRecipe          `json:"payouts,omitempty"`
	AccumulatedPayouts            []PayoutRecipe          `json:"accumulated_payouts,omitempty"`
	InvalidPayouts                []PayoutRecipe          `json:"invalid_payouts,omitempty"`
	ReportsOfPastSuccesfulPayouts []PayoutReport          `json:"reports_of_past_succesful_payouts,omitempty"`
}

type ExecutePayoutsEngineContext struct {
	signer      SignerEngine
	transactor  TransactorEngine
	reporter    ReporterEngine
	adminNotify func(msg string)
}

func NewExecutePayoutsEngineContext(signer SignerEngine, transactor TransactorEngine, reporter ReporterEngine, adminNotify func(msg string)) *ExecutePayoutsEngineContext {
	return &ExecutePayoutsEngineContext{
		signer:      signer,
		transactor:  transactor,
		reporter:    reporter,
		adminNotify: adminNotify,
	}
}

func (engines *ExecutePayoutsEngineContext) GetSigner() SignerEngine {
	return engines.signer
}

func (engines *ExecutePayoutsEngineContext) GetTransactor() TransactorEngine {
	return engines.transactor
}

func (engines *ExecutePayoutsEngineContext) GetReporter() ReporterEngine {
	return engines.reporter
}

func (engines *ExecutePayoutsEngineContext) AdminNotify(msg string) {
	if engines.adminNotify != nil {
		engines.adminNotify(msg)
	}
}

func (engines *ExecutePayoutsEngineContext) Validate() error {
	if engines.signer == nil {
		return errors.Join(constants.ErrMissingEngine, constants.ErrMissingSignerEngine)
	}
	if engines.transactor == nil {
		return errors.Join(constants.ErrMissingEngine, constants.ErrMissingTransactorEngine)
	}
	if engines.reporter == nil {
		return errors.Join(constants.ErrMissingEngine, constants.ErrMissingReporterEngine)
	}
	return nil
}

type ExecutePayoutsOptions struct {
	MixInContractCalls bool `json:"mix_in_contract_calls,omitempty"`
	MixInFATransfers   bool `json:"mix_in_fa_transfers,omitempty"`
	DryRun             bool `json:"dry_run,omitempty"`
}

type ExecutePayoutsResult struct {
	BatchResults   BatchResults `json:"batch_results,omitempty"`
	PaidDelegators int          `json:"paid_delegators,omitempty"`
}

package common

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/gomavryk/mavryk"
	"github.com/samber/lo"
	"github.com/mavryk-network/gomavryk/base58"
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
	FATokenId        mavryk.Z                      `json:"fa_token_id,omitempty"`
	FAContract       mavryk.Address                `json:"fa_contract,omitempty"`
	FAAlias          string                       `json:"fa_alias,omitempty"`
	FADecimals       int                          `json:"fa_decimals,omitempty"`
	DelegatedBalance mavryk.Z                      `json:"delegator_balance,omitempty"`
	StakedBalance    mavryk.Z                      `json:"staked_balance,omitempty"`
	Amount           mavryk.Z                      `json:"amount,omitempty"`
	FeeRate          float64                      `json:"fee_rate,omitempty"`
	Fee              mavryk.Z                      `json:"fee,omitempty"`
	TxFee            int64                        `json:"tx_fee,omitempty"` // calculated during fee estimation
	Note             string                       `json:"note,omitempty"`
	IsValid          bool                         `json:"valid,omitempty"`
}

func (candidate PayoutRecipe) GetKind() enums.EPayoutKind {
	return candidate.Kind
}

func (candidate PayoutRecipe) GetDelegatedBalance() mavryk.Z {
	return candidate.DelegatedBalance
}

func (candidate *PayoutRecipe) GetDestination() mavryk.Address {
	return candidate.Recipient
}

func (candidate PayoutRecipe) GetTxKind() enums.EPayoutTransactionKind {
	return candidate.TxKind
}

func (candidate *PayoutRecipe) GetFATokenId() mavryk.Z {
	return candidate.FATokenId
}

func (candidate *PayoutRecipe) GetFAContract() mavryk.Address {
	return candidate.FAContract
}

func (candidate PayoutRecipe) GetAmount() mavryk.Z {
	return candidate.Amount
}

func (candidate PayoutRecipe) GetFee() mavryk.Z {
	return candidate.Fee
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
	}
	k, err := identifier.ToJSON()
	if err != nil {
		return ""
	}
	hashBytes := sha256.Sum256(k)
	return base58.Encode(hashBytes[:])
}

func (recipe *PayoutRecipe) GetShortIdentifier() string {
	return recipe.GetIdentifier()[:16]
}

func (recipe PayoutRecipe) AsAccumulated() *AccumulatedPayoutRecipe {
	clone := recipe // make a copy to avoid issues with references
	return &AccumulatedPayoutRecipe{
		Delegator:  clone.Delegator,
		Cycle:      clone.Cycle,
		Recipient:  clone.Recipient,
		Kind:       clone.Kind,
		TxKind:     clone.TxKind,
		FATokenId:  clone.FATokenId,
		FAContract: clone.FAContract,
		IsValid:    clone.IsValid,
		Recipes:    []*PayoutRecipe{&clone},
		Note:       clone.Note,
	}
}

func (pr PayoutRecipe) ToPayoutReport() PayoutReport {
	return PayoutReport{
		Id:               pr.GetShortIdentifier(),
		Baker:            pr.Baker,
		Timestamp:        time.Now(),
		Cycle:            pr.Cycle,
		Kind:             pr.Kind,
		TxKind:           pr.TxKind,
		FAContract:       pr.FAContract,
		FATokenId:        pr.FATokenId,
		FAAlias:          pr.FAAlias,
		FADecimals:       pr.FADecimals,
		Delegator:        pr.Delegator,
		DelegatedBalance: pr.DelegatedBalance,
		StakedBalance:    pr.StakedBalance,
		Recipient:        pr.Recipient,
		Amount:           pr.Amount,
		FeeRate:          pr.FeeRate,
		Fee:              pr.Fee,
		TxFee:            pr.TxFee,
		OpHash:           mavryk.ZeroOpHash,
		IsSuccess:        false,
		Note:             pr.Note,
	}
}

func (pr PayoutRecipe) GetTxFee() int64 {
	return pr.TxFee
}

func (pr *PayoutRecipe) ToTableRowData() []string {
	return []string{
		ShortenAddress(pr.Delegator),
		ShortenAddress(pr.Recipient),
		MumavToMavS(pr.DelegatedBalance.Int64()),
		string(pr.Kind),
		ShortenAddress(pr.FAContract),
		ToStringEmptyIfZero(pr.FATokenId.Int64()),
		FormatTokenAmount(pr.TxKind, pr.Amount.Int64(), pr.FAAlias, pr.FADecimals),
		FloatToPercentage(pr.FeeRate),
		MumavToMavS(pr.Fee.Int64()),
		MumavToMavS(pr.GetTxFee()),
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

// returns totals and number of filtered recipes

type foo interface {
	GetKind() enums.EPayoutKind
	GetTxKind() enums.EPayoutTransactionKind
	GetAmount() mavryk.Z
	GetFee() mavryk.Z
	GetTxFee() int64
}

func GetRecipesTotals[T foo](recipes []T, withFee bool) []string {
	totalAmount := int64(0)
	totalFee := int64(0)
	totalTx := int64(0)
	for _, recipe := range recipes {
		if recipe.GetTxKind() == enums.PAYOUT_TX_KIND_MAV {
			totalAmount += recipe.GetAmount().Int64()
		}
		totalFee += recipe.GetFee().Int64()
		totalTx += recipe.GetTxFee()
	}
	fee := ""
	if withFee {
		fee = MumavToMavS(totalFee)
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
		fee,
		MumavToMavS(totalTx),
		"",
	}
}

func GetRecipesFilteredTotals[T foo](recipes []T, kind enums.EPayoutKind, withFee bool) ([]string, int) {
	r := lo.Filter(recipes, func(recipe T, _ int) bool {
		return recipe.GetKind() == kind
	})
	return GetRecipesTotals(r, withFee), len(r)
}

type AccumulatedPayoutRecipe struct {
	// PayoutRecipe
	Delegator  mavryk.Address                `json:"delegator,omitempty"`
	Cycle      int64                        `json:"cycle,omitempty"`
	Recipient  mavryk.Address                `json:"recipient,omitempty"`
	Kind       enums.EPayoutKind            `json:"kind,omitempty"`
	TxKind     enums.EPayoutTransactionKind `json:"tx_kind,omitempty"`
	FATokenId  mavryk.Z                      `json:"fa_token_id,omitempty"`
	FAContract mavryk.Address                `json:"fa_contract,omitempty"`
	IsValid    bool                         `json:"valid,omitempty"`
	Note       string                       `json:"note,omitempty"`

	OpLimits *OpLimits       `json:"op_limits,omitempty"`
	Recipes  []*PayoutRecipe `json:"-"`
}

func (recipe *AccumulatedPayoutRecipe) GetTxFee() int64 {
	return lo.Reduce(recipe.Recipes, func(agg int64, recipe *PayoutRecipe, _ int) int64 {
		return agg + recipe.TxFee
	}, 0)
}

func (recipe *AccumulatedPayoutRecipe) Sum() PayoutRecipe {
	if len(recipe.Recipes) == 0 {
		return PayoutRecipe{}
	}
	if len(recipe.Recipes) == 1 {
		return *recipe.Recipes[0]
	}

	result := *recipe.Recipes[0]
	for _, r := range recipe.Recipes[1:] {
		result.DelegatedBalance = result.DelegatedBalance.Add(r.DelegatedBalance).Div64(2)
		result.StakedBalance = result.StakedBalance.Add(r.StakedBalance).Div64(2)
		result.Amount = result.Amount.Add(r.Amount)
		result.Fee = result.Fee.Add(r.Fee)
		result.TxFee = result.TxFee + r.TxFee
	}
	return result
}

func (r *AccumulatedPayoutRecipe) ToTableRowData() []string {
	recipe := r.Sum()
	return []string{
		ShortenAddress(recipe.Delegator),
		ShortenAddress(recipe.Recipient),
		MumavToMavS(recipe.DelegatedBalance.Int64()),
		string(recipe.Kind),
		ShortenAddress(recipe.FAContract),
		ToStringEmptyIfZero(recipe.FATokenId.Int64()),
		FormatTokenAmount(recipe.TxKind, recipe.Amount.Int64(), recipe.FAAlias, recipe.FADecimals),
		FloatToPercentage(recipe.FeeRate),
		MumavToMavS(recipe.Fee.Int64()),
		MumavToMavS(recipe.GetTxFee()),
		recipe.Note,
	}
}

func (recipe *AccumulatedPayoutRecipe) Add(otherRecipe *PayoutRecipe) (*AccumulatedPayoutRecipe, error) {
	if !recipe.Recipient.Equal(otherRecipe.Recipient) {
		return nil, errors.New("cannot add different recipients")
	}
	if !recipe.Delegator.Equal(otherRecipe.Delegator) {
		return nil, errors.New("cannot add different delegators")
	}
	if recipe.Kind != otherRecipe.Kind {
		return nil, errors.New("cannot add different kinds")
	}
	if recipe.TxKind != otherRecipe.TxKind {
		return nil, errors.New("cannot add different tx kinds")
	}
	if !recipe.FAContract.Equal(otherRecipe.FAContract) {
		return nil, errors.New("cannot add different FA contracts")
	}
	if !recipe.FATokenId.Equal(otherRecipe.FATokenId) {
		return nil, errors.New("cannot add different FA token ids")
	}
	if recipe.IsValid != otherRecipe.IsValid {
		return nil, errors.New("cannot add different validity states")
	}

	otherRecipe.Note = fmt.Sprintf("%s_%d", recipe.GetShortIdentifier(), recipe.Cycle)
	recipe.Recipes = append(recipe.Recipes, otherRecipe)
	return recipe, nil
}

func (recipe *AccumulatedPayoutRecipe) GetAmount() mavryk.Z {
	return lo.Reduce(recipe.Recipes, func(agg mavryk.Z, recipe *PayoutRecipe, _ int) mavryk.Z {
		return agg.Add(recipe.Amount)
	}, mavryk.Zero)
}

func (recipe *AccumulatedPayoutRecipe) AddTxFee(amount mavryk.Z, charge bool) {
	if len(recipe.Recipes) == 0 {
		panic("THIS SHOULD NEVER HAPPEN: cannot add tx fee to empty accumulated payout")
	}

	if !charge {
		recipe.Recipes[0].TxFee = recipe.Recipes[0].TxFee + amount.Int64()
		return
	}
	// charge
	remainder := amount
	for _, r := range recipe.Recipes {
		if r.Amount.IsLessEqual(remainder) {
			remainder = remainder.Sub(r.Amount)
			r.TxFee = r.TxFee + r.Amount.Int64()
			r.Amount = mavryk.Zero
			continue
		}
		r.Amount = r.Amount.Sub(remainder)
		r.TxFee = r.TxFee + remainder.Int64()
		break
	}
}

func (recipe *AccumulatedPayoutRecipe) AddTxFee64(amount int64, charge bool) {
	recipe.AddTxFee(mavryk.NewZ(amount), charge)
}

func (recipe *AccumulatedPayoutRecipe) GetFee() mavryk.Z {
	return lo.Reduce(recipe.Recipes, func(agg mavryk.Z, recipe *PayoutRecipe, _ int) mavryk.Z {
		return agg.Add(recipe.Fee)
	}, mavryk.Zero)
}

func (recipe *AccumulatedPayoutRecipe) GetKind() enums.EPayoutKind {
	return recipe.Kind
}

func (recipe *AccumulatedPayoutRecipe) GetTxKind() enums.EPayoutTransactionKind {
	return recipe.TxKind
}

func (recipe *AccumulatedPayoutRecipe) GetFAContract() mavryk.Address {
	return recipe.FAContract
}

func (recipe *AccumulatedPayoutRecipe) GetFATokenId() mavryk.Z {
	return recipe.FATokenId
}

func (recipe *AccumulatedPayoutRecipe) GetDestination() mavryk.Address {
	return recipe.Recipient
}

func (recipe *AccumulatedPayoutRecipe) GetDelegatedBalance() mavryk.Z {
	if len(recipe.Recipes) == 0 {
		return mavryk.Zero
	}
	return recipe.Recipes[0].DelegatedBalance
}

func (recipe *AccumulatedPayoutRecipe) DisperseToInvalid() []PayoutRecipe {
	if recipe.IsValid {
		panic("THIS SHOULD NEVER HAPPEN: cannot disperse valid accumulated payout")
	}

	return lo.Map(recipe.Recipes, func(r *PayoutRecipe, _ int) PayoutRecipe {
		r.IsValid = false
		r.Note = recipe.Note
		return *r
	})
}
func (recipe *AccumulatedPayoutRecipe) GetIdentifier() string {
	identifier := PayoutRecipeIdentifier{
		Delegator:  recipe.Delegator,
		Recipient:  recipe.Recipient,
		Kind:       recipe.Kind,
		TxKind:     recipe.TxKind,
		FATokenId:  recipe.FATokenId,
		FAContract: recipe.FAContract,
		// IsValid:    recipe.IsValid,
	}
	k, err := identifier.ToJSON()
	if err != nil {
		return ""
	}
	hashBytes := sha256.Sum256(k)
	return base58.Encode(hashBytes[:])
}

func (recipe *AccumulatedPayoutRecipe) GetShortIdentifier() string {
	return recipe.GetIdentifier()[:16]
}

// AsRecipe returns the PayoutRecipe representation of the AccumulatedPayoutRecipe.
// This is useful only for printing and reporting purposes. Do not use it for execution.
func (recipe *AccumulatedPayoutRecipe) AsRecipe() PayoutRecipe {
	return recipe.Sum()
}

func (recipe *AccumulatedPayoutRecipe) DeepClone() *AccumulatedPayoutRecipe {
	clonedRecipes := lo.Map(recipe.Recipes, func(r *PayoutRecipe, _ int) *PayoutRecipe {
		clone := *r
		return &clone
	})
	return &AccumulatedPayoutRecipe{
		Delegator:  recipe.Delegator,
		Cycle:      recipe.Cycle,
		Recipient:  recipe.Recipient,
		Kind:       recipe.Kind,
		TxKind:     recipe.TxKind,
		FATokenId:  recipe.FATokenId,
		FAContract: recipe.FAContract,
		IsValid:    recipe.IsValid,
		Note:       recipe.Note,
		OpLimits:   recipe.OpLimits, // shallow copy is fine
		Recipes:    clonedRecipes,
	}
}

type CyclePayoutSummary struct {
	Delegators               int       `json:"delegators"`
	PaidDelegators           int       `json:"paid_delegators"`
	OwnStakedBalance         mavryk.Z   `json:"own_staked_balance"`
	OwnDelegatedBalance      mavryk.Z   `json:"own_delegated_balance"`
	ExternalStakedBalance    mavryk.Z   `json:"external_staked_balance"`
	ExternalDelegatedBalance mavryk.Z   `json:"external_delegated_balance"`
	EarnedBlockFees          mavryk.Z   `json:"cycle_earned_fees"`
	EarnedRewards            mavryk.Z   `json:"cycle_earned_rewards"`
	EarnedTotal              mavryk.Z   `json:"cycle_earned_total"`
	DistributedRewards       mavryk.Z   `json:"distributed_rewards"`
	NotDistributedRewards    mavryk.Z   `json:"not_distributed_rewards"`
	BondIncome               mavryk.Z   `json:"bond_income"`
	FeeIncome                mavryk.Z   `json:"fee_income"`
	IncomeTotal              mavryk.Z   `json:"total_income"`
	TxFeesPaidForRewards     mavryk.Z   `json:"tx_fees_paid_for_rewards"`
	TxFeesPaid               mavryk.Z   `json:"tx_fees_paid"`
	DonatedBonds             mavryk.Z   `json:"donated_bonds"`
	DonatedFees              mavryk.Z   `json:"donated_fees"`
	DonatedTotal             mavryk.Z   `json:"donated_total"`
	Timestamp                time.Time `json:"timestamp"`
}

type PayoutSummary struct {
	CyclePayoutSummary
	Cycles         []int64                      `json:"cycle"`
	CycleSummaries map[int64]CyclePayoutSummary `json:"cycle_summaries,omitempty"`
}

func (summary *PayoutSummary) GetTotalStakedBalance() mavryk.Z {
	return summary.OwnStakedBalance.Add(summary.ExternalStakedBalance)
}

func (summary *PayoutSummary) GetTotalDelegatedBalance() mavryk.Z {
	return summary.OwnDelegatedBalance.Add(summary.ExternalDelegatedBalance)
}

func (summary *PayoutSummary) AddCycleSummary(cycle int64, another *CyclePayoutSummary) {
	if summary.CycleSummaries == nil {
		summary.CycleSummaries = make(map[int64]CyclePayoutSummary)
	}
	if _, ok := summary.CycleSummaries[cycle]; ok {
		panic("cannot add the same cycle summary twice")
	}
	summary.Cycles = append(summary.Cycles, cycle)
	slices.Sort(summary.Cycles)
	summary.CycleSummaries[cycle] = *another

	summary.OwnStakedBalance = summary.OwnStakedBalance.Add(another.OwnStakedBalance)
	summary.OwnDelegatedBalance = summary.OwnDelegatedBalance.Add(another.OwnDelegatedBalance)
	summary.ExternalStakedBalance = summary.ExternalStakedBalance.Add(another.ExternalStakedBalance)
	summary.ExternalDelegatedBalance = summary.ExternalDelegatedBalance.Add(another.ExternalDelegatedBalance)
	summary.EarnedBlockFees = summary.EarnedBlockFees.Add(another.EarnedBlockFees)
	summary.EarnedRewards = summary.EarnedRewards.Add(another.EarnedRewards)
	summary.EarnedTotal = summary.EarnedTotal.Add(another.EarnedTotal)
	summary.DistributedRewards = summary.DistributedRewards.Add(another.DistributedRewards)
	summary.NotDistributedRewards = summary.NotDistributedRewards.Add(another.NotDistributedRewards)
	summary.BondIncome = summary.BondIncome.Add(another.BondIncome)
	summary.FeeIncome = summary.FeeIncome.Add(another.FeeIncome)
	summary.IncomeTotal = summary.IncomeTotal.Add(another.IncomeTotal)
	summary.TxFeesPaid = summary.TxFeesPaid.Add(another.TxFeesPaid)
	summary.TxFeesPaidForRewards = summary.TxFeesPaidForRewards.Add(another.TxFeesPaidForRewards)
	summary.DonatedBonds = summary.DonatedBonds.Add(another.DonatedBonds)
	summary.DonatedFees = summary.DonatedFees.Add(another.DonatedFees)
	summary.DonatedTotal = summary.DonatedTotal.Add(another.DonatedTotal)
}

type CyclePayoutBlueprint struct {
	Cycle   int64          `json:"cycle,omitempty"`
	Payouts []PayoutRecipe `json:"payouts,omitempty"`

	OwnStakedBalance         mavryk.Z   `json:"own_staked_balance"`
	OwnDelegatedBalance      mavryk.Z   `json:"own_delegated_balance"`
	ExternalStakedBalance    mavryk.Z   `json:"external_staked_balance"`
	ExternalDelegatedBalance mavryk.Z   `json:"external_delegated_balance"`
	EarnedBlockFees          mavryk.Z   `json:"cycle_earned_fees"`
	EarnedRewards            mavryk.Z   `json:"cycle_earned_rewards"`
	EarnedTotal              mavryk.Z   `json:"cycle_earned_total"`
	BondIncome               mavryk.Z   `json:"bond_income"`
	FeeIncome                mavryk.Z   `json:"fee_income"`
	IncomeTotal              mavryk.Z   `json:"total_income"`
	DonatedBonds             mavryk.Z   `json:"donated_bonds"`
	DonatedFees              mavryk.Z   `json:"donated_fees"`
	DonatedTotal             mavryk.Z   `json:"donated_total"`
	Timestamp                time.Time `json:"timestamp"`
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
	Cycle int64 `json:"cycle,omitempty"`
}

type CyclePayoutBlueprints []*CyclePayoutBlueprint

func (results CyclePayoutBlueprints) GetCycles() []int64 {
	return lo.Reduce(results, func(acc []int64, result *CyclePayoutBlueprint, _ int) []int64 {
		for _, p := range result.Payouts {
			if !slices.Contains(acc, p.Cycle) {
				acc = append(acc, p.Cycle)
			}
		}
		return acc
	}, []int64{})
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
	Accumulate               bool `json:"accumulate,omitempty"`
	SkipBalanceCheck         bool `json:"skip_balance_check,omitempty"`
	WaitForSufficientBalance bool `json:"wait_for_sufficient_balance,omitempty"`
}

type PreparePayoutsResult struct {
	Blueprints                           []*CyclePayoutBlueprint    `json:"blueprint,omitempty"`
	ValidPayouts                         []*AccumulatedPayoutRecipe `json:"payouts,omitempty"`
	InvalidPayouts                       []PayoutRecipe             `json:"invalid_payouts,omitempty"`
	ReportsOfPastSuccessfulPayouts       []PayoutReport             `json:"reports_of_past_successful_payouts,omitempty"`
	BatchMetadataDeserializationGasLimit int64                      `json:"batch_metadata_deserialization_gas_limit,omitempty"`
}

func (result *PreparePayoutsResult) GetCycles() []int64 {
	return lo.Reduce(result.Blueprints, func(acc []int64, blueprint *CyclePayoutBlueprint, _ int) []int64 {
		if !slices.Contains(acc, blueprint.Cycle) {
			acc = append(acc, blueprint.Cycle)
		}
		return acc
	}, []int64{})
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
	BatchResults   BatchResults  `json:"batch_results,omitempty"`
	PaidDelegators int           `json:"paid_delegators,omitempty"`
	Summary        PayoutSummary `json:"cycle_payout_summary,omitempty"`
}

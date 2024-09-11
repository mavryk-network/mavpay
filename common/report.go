package common

import (
	"time"

	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/samber/lo"
)

type PayoutReport struct {
	Id               string                       `json:"id" csv:"id"`
	Baker            mavryk.Address               `json:"baker" csv:"baker"`
	Timestamp        time.Time                    `json:"timestamp" csv:"timestamp"`
	Cycle            int64                        `json:"cycle" csv:"cycle"`
	Kind             enums.EPayoutKind            `json:"kind,omitempty" csv:"kind"`
	TxKind           enums.EPayoutTransactionKind `json:"tx_kind,omitempty" csv:"op_kind"`
	FAContract       mavryk.Address               `json:"contract,omitempty" csv:"contract"`
	FATokenId        mavryk.Z                     `json:"token_id,omitempty" csv:"token_id"`
	Delegator        mavryk.Address               `json:"delegator,omitempty" csv:"delegator"`
	DelegatedBalance mavryk.Z                     `json:"delegator_balance,omitempty" csv:"delegator_balance"`
	StakedBalance    mavryk.Z                     `json:"-" csv:"-"` // enable when relevant
	Recipient        mavryk.Address               `json:"recipient,omitempty" csv:"recipient"`
	Amount           mavryk.Z                     `json:"amount,omitempty" csv:"amount"`
	FeeRate          float64                      `json:"fee_rate,omitempty" csv:"fee_rate"`
	Fee              mavryk.Z                     `json:"fee,omitempty" csv:"fee"`
	TransactionFee   int64                        `json:"tx_fee,omitempty" csv:"tx_fee"`
	OpHash           mavryk.OpHash                `json:"op_hash,omitempty" csv:"op_hash"`
	IsSuccess        bool                         `json:"success" csv:"success"`
	Note             string                       `json:"note,omitempty" csv:"note"`
}

func (pr *PayoutReport) GetTransactionFee() int64 {
	return pr.TransactionFee
}

func (pr *PayoutReport) ToTableRowData() []string {
	return []string{
		ShortenAddress(pr.Delegator),
		ShortenAddress(pr.Recipient),
		MutezToTezS(pr.DelegatedBalance.Int64()),
		string(pr.Kind),
		ShortenAddress(pr.FAContract),
		ToStringEmptyIfZero(pr.FATokenId.Int64()),
		FormatAmount(pr.TxKind, pr.Amount.Int64()),
		FloatToPercentage(pr.FeeRate),
		MutezToTezS(pr.Fee.Int64()),
		MutezToTezS(pr.GetTransactionFee()),
		pr.OpHash.String(),
		pr.Note,
	}
}

func (pr *PayoutReport) GetTableHeaders() []string {
	return []string{
		"Delegator",
		"Recipient",
		"Delegated Balance",
		"Kind",
		"FA Contract",
		"FA Token ID",
		"Amount",
		"Fee Rate",
		"Fee",
		"Transaction Fee",
		"Op Hash",
		"Note",
	}
}

func GetReportsTotals(reports []PayoutReport) []string {
	var totalAmount, totalFee, totalTxFee int64
	for _, report := range reports {
		if report.TxKind == enums.PAYOUT_TX_KIND_TEZ {
			totalAmount += report.Amount.Int64()
		}
		totalFee += report.Fee.Int64()
		totalTxFee += report.GetTransactionFee()
	}
	return []string{
		"",
		"",
		"",
		"",
		"",
		"",
		MutezToTezS(totalAmount),
		"",
		MutezToTezS(totalFee),
		MutezToTezS(totalTxFee),
		"",
		"",
	}
}

// returns total amounts and count of filtered reports
func GetFilteredReportsTotals(reports []PayoutReport, kind enums.EPayoutKind) ([]string, int) {
	r := lo.Filter(reports, func(report PayoutReport, _ int) bool {
		return report.Kind == kind
	})
	return GetReportsTotals(r), len(r)
}

type PayoutCycleReport struct {
	Cycle   int64               `json:"cycle"`
	Invalid []PayoutRecipe      `json:"invalid,omitempty"`
	Payouts []PayoutReport      `json:"payouts"`
	Sumary  *CyclePayoutSummary `json:"summary"`
}

package utils

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/samber/lo"
)

const (
	TOTAL_PAYOUTS = "Rewards"
	TOTAL         = "Total"
)

func getColumnsByIndexes[T any](row []T, indexes []int) []T {
	return lo.Filter(row, func(_ T, i int) bool {
		return lo.Contains(indexes, i)
	})
}

func columnsAsInterfaces[T any](row []T) []any {
	return lo.Map(row, func(c T, _ int) any {
		return c
	})
}

func replaceZeroFields[T comparable](items []T, value T, stopOnNonEmpty bool) []T {
	var zero T
	for i, item := range items {
		if item == zero {
			items[i] = value
		} else if stopOnNonEmpty {
			break
		}
	}
	return items
}

func getNonEmptyIndexes[T comparable](headers []string, data [][]T) []int {
	var zero T
	return lo.Filter(lo.Range(len(headers)), func(c int, i int) bool {
		return lo.SomeBy(data, func(d []T) bool {
			return d[i] != zero
		})
	})
}

func sortPayouts(payouts []common.PayoutRecipe) {
	slices.SortFunc(payouts, func(a, b common.PayoutRecipe) int {
		if a.Kind == b.Kind {
			if a.Amount.IsLess(b.Amount) {
				return 1
			} else if b.Amount.IsLess(a.Amount) {
				return -1
			} else {
				return 0
			}
		}
		if a.Kind.ToPriority() < b.Kind.ToPriority() {
			return 1
		} else if a.Kind.ToPriority() > b.Kind.ToPriority() {
			return -1
		}
		return 0
	})
}

func PrintPayouts(payouts []common.PayoutRecipe, header string, printTotals bool) {
	if len(payouts) == 0 {
		return
	}

	sortPayouts(payouts)

	payoutTable := table.NewWriter()
	payoutTable.SetStyle(table.StyleLight)
	payoutTable.SetColumnConfigs([]table.ColumnConfig{{Number: 1, Align: text.AlignLeft}, {Number: 2, Align: text.AlignLeft}})
	payoutTable.SetOutputMirror(os.Stdout)
	payoutTable.SetTitle(header)
	payoutTable.Style().Title.Align = text.AlignCenter

	headers := payouts[0].GetTableHeaders()
	data := lo.Map(payouts, func(p common.PayoutRecipe, _ int) []string {
		return p.ToTableRowData()
	})
	validIndexes := getNonEmptyIndexes(headers, data)

	payoutTable.AppendHeader(columnsAsInterfaces(getColumnsByIndexes(headers, validIndexes)), table.RowConfig{AutoMerge: true})

	for _, payout := range data {
		row := replaceZeroFields(payout, "-", false)
		payoutTable.AppendRow(columnsAsInterfaces(getColumnsByIndexes(row, validIndexes)), table.RowConfig{AutoMerge: false})
	}
	if printTotals {
		payoutTable.AppendSeparator()
		rewardsTotals, countOfRwards := common.GetRecipesFilteredTotals(payouts, enums.PAYOUT_KIND_DELEGATOR_REWARD)
		totalsRewards := replaceZeroFields(rewardsTotals, fmt.Sprintf("%s (%d)", TOTAL_PAYOUTS, countOfRwards), true)
		totalsRewards = replaceZeroFields(totalsRewards, "-", false)
		payoutTable.AppendRow(columnsAsInterfaces(getColumnsByIndexes(totalsRewards, validIndexes)), table.RowConfig{AutoMerge: true})

		payoutTable.AppendSeparator()
		totals := replaceZeroFields(common.GetRecipesTotals(payouts), fmt.Sprintf("%s (%d)", TOTAL, len(payouts)), true)
		totals = replaceZeroFields(totals, "-", false)

		payoutTable.AppendRow(columnsAsInterfaces(getColumnsByIndexes(totals, validIndexes)), table.RowConfig{AutoMerge: true})
	}
	payoutTable.Render()
}

func FormatCycleNumbers(cycles ...int64) string {
	conscutive := false
	if len(cycles) > 1 {
		conscutive = true
		for i := 1; i < len(cycles); i++ {
			if cycles[i] != cycles[i-1]+1 {
				conscutive = false
				break
			}
		}
	}
	if conscutive {
		return fmt.Sprintf("#%d-%d", cycles[0], cycles[len(cycles)-1])
	} else {
		return fmt.Sprintf("#%s", strings.Join(lo.Map(cycles, func(c int64, _ int) string {
			return fmt.Sprintf("%d", c)
		}), ", "))
	}
}

// // print invalid payouts
// func PrintInvalidPayoutRecipes(payouts []common.PayoutRecipe, cycles []int64) {
// 	printPayouts(OnlyInvalidPayouts(payouts), fmt.Sprintf("Invalid - %s", FormatCycleNumbers(cycles)), false)
// }

// // print payable payouts
// func PrintValidPayoutRecipes(payouts []common.PayoutRecipe, cycles []int64) {
// 	printPayouts(OnlyValidPayouts(payouts), fmt.Sprintf("Valid - %s", FormatCycleNumbers(cycles)), true)
// }

func IsTty() bool {
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		return true
	} else {
		return false
	}
}

func PrintReports(payouts []common.PayoutReport, header string, printTotals bool) {
	if len(payouts) == 0 {
		return
	}
	payoutTable := table.NewWriter()
	payoutTable.SetStyle(table.StyleLight)
	payoutTable.SetColumnConfigs([]table.ColumnConfig{{Number: 1, Align: text.AlignLeft}, {Number: 2, Align: text.AlignLeft}})
	payoutTable.SetOutputMirror(os.Stdout)
	payoutTable.SetTitle(header)
	payoutTable.Style().Title.Align = text.AlignCenter

	headers := payouts[0].GetTableHeaders()
	data := lo.Map(payouts, func(p common.PayoutReport, _ int) []string {
		return p.ToTableRowData()
	})
	validIndexes := getNonEmptyIndexes(headers, data)

	payoutTable.AppendHeader(columnsAsInterfaces(getColumnsByIndexes(headers, validIndexes)), table.RowConfig{AutoMerge: true})
	for _, payout := range data {
		row := replaceZeroFields(payout, "-", false)
		payoutTable.AppendRow(columnsAsInterfaces(getColumnsByIndexes(row, validIndexes)), table.RowConfig{AutoMerge: false})
	}
	if printTotals {
		payoutTable.AppendSeparator()
		rewardsTotals, countOfRwards := common.GetFilteredReportsTotals(payouts, enums.PAYOUT_KIND_DELEGATOR_REWARD)
		totalsRewards := replaceZeroFields(rewardsTotals, fmt.Sprintf("%s (%d)", TOTAL_PAYOUTS, countOfRwards), true)
		totalsRewards = replaceZeroFields(totalsRewards, "-", false)
		payoutTable.AppendRow(columnsAsInterfaces(getColumnsByIndexes(totalsRewards, validIndexes)), table.RowConfig{AutoMerge: true})

		payoutTable.AppendSeparator()
		totals := replaceZeroFields(common.GetReportsTotals(payouts), fmt.Sprintf("%s (%d)", TOTAL, len(payouts)), true)
		totals = replaceZeroFields(totals, "-", false)
		payoutTable.AppendRow(columnsAsInterfaces(getColumnsByIndexes(totals, validIndexes)), table.RowConfig{AutoMerge: true})
	}
	payoutTable.Render()
}

func PrintCycleSummary(summary common.CyclePayoutSummary, header string) {
	summaryTable := table.NewWriter()
	summaryTable.SetStyle(table.StyleLight)
	summaryTable.SetColumnConfigs([]table.ColumnConfig{{Number: 1, Align: text.AlignLeft}, {Number: 2, Align: text.AlignRight}})
	summaryTable.SetOutputMirror(os.Stdout)
	summaryTable.SetTitle(header)
	summaryTable.Style().Title.Align = text.AlignCenter
	summaryTable.AppendRow(table.Row{"Earned Fees", common.MutezToTezS(summary.EarnedFees.Int64())}, table.RowConfig{AutoMerge: false})
	summaryTable.AppendRow(table.Row{"Earned Rewards", common.MutezToTezS(summary.EarnedRewards.Int64())}, table.RowConfig{AutoMerge: false})
	summaryTable.AppendRow(table.Row{"Distributed Rewards", common.MutezToTezS(summary.DistributedRewards.Int64())}, table.RowConfig{AutoMerge: false})
	summaryTable.AppendSeparator()
	summaryTable.AppendRow(table.Row{"Donated Bonds", common.MutezToTezS(summary.DonatedBonds.Int64())}, table.RowConfig{AutoMerge: false})
	summaryTable.AppendRow(table.Row{"Donated Fees", common.MutezToTezS(summary.DonatedFees.Int64())}, table.RowConfig{AutoMerge: false})
	summaryTable.AppendRow(table.Row{"Donated Total", common.MutezToTezS(summary.DonatedTotal.Int64())}, table.RowConfig{AutoMerge: false})
	summaryTable.AppendSeparator()
	summaryTable.AppendRow(table.Row{"Bond Income", common.MutezToTezS(summary.BondIncome.Int64())}, table.RowConfig{AutoMerge: false})
	summaryTable.AppendRow(table.Row{"Fee Income", common.MutezToTezS(summary.FeeIncome.Int64())}, table.RowConfig{AutoMerge: false})
	summaryTable.AppendRow(table.Row{"Income Total", common.MutezToTezS(summary.IncomeTotal.Int64())}, table.RowConfig{AutoMerge: false})
	summaryTable.Render()
}

func PrintBatchResults(results []common.BatchResult, header string, explorerUrl string) {
	if len(results) == 0 {
		return
	}
	resultsTable := table.NewWriter()
	resultsTable.SetStyle(table.StyleLight)
	resultsTable.SetColumnConfigs([]table.ColumnConfig{{Number: 1, Align: text.AlignLeft}, {Number: 2, Align: text.AlignLeft}})
	resultsTable.SetOutputMirror(os.Stdout)
	resultsTable.SetTitle(header)
	resultsTable.Style().Title.Align = text.AlignCenter
	resultsTable.AppendHeader(table.Row{"n.", "Transactions", "Success", "Reference"}, table.RowConfig{AutoMerge: true})
	for i, result := range results {
		resultsTable.AppendRow(table.Row{i + 1, len(result.Payouts), result.IsSuccess, GetOpReference(result.OpHash, explorerUrl)}, table.RowConfig{AutoMerge: false})
	}
	resultsTable.Render()
}

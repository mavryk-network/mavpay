package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mavpay/core/generate"
	"github.com/mavryk-network/mavpay/core/prepare"
	"github.com/mavryk-network/mvgo/mavryk"
)

func GenerateHookSampleData() {
	payoutCandidate := generate.PayoutCandidateWithBondAmountAndFee{
		PayoutCandidateWithBondAmount: generate.PayoutCandidateWithBondAmount{
			PayoutCandidate: generate.PayoutCandidate{
				Source:                       mavryk.ZeroAddress,
				Recipient:                    mavryk.ZeroAddress,
				FeeRate:                      5.0,
				DelegatedBalance:             mavryk.NewZ(1000000000),
				StakedBalance:                mavryk.NewZ(1000000000),
				IsInvalid:                    true,
				IsEmptied:                    true,
				IsBakerPayingTxFee:           true,
				IsBakerPayingAllocationTxFee: true,
				InvalidBecause:               "reason",
			},
			BondsAmount: mavryk.NewZ(1000000000),
			TxKind:      "fa1",
			FATokenId:   mavryk.NewZ(10),
			FAContract:  mavryk.ZeroContract,
		},
		Fee: mavryk.NewZ(1000000000),
	}

	acg := generate.AfterCandidateGeneratedHookData{
		Cycle:      580,
		Candidates: []generate.PayoutCandidate{payoutCandidate.PayoutCandidate},
	}

	abd := generate.AfterBondsDistributedHookData{
		Cycle:      580,
		Candidates: []generate.PayoutCandidateWithBondAmount{payoutCandidate.PayoutCandidateWithBondAmount},
	}
	acb := generate.CheckBalanceHookData{
		SkipMavCheck: true,
		Message:      "This message is used to carry errors from hook to the caller.",
		IsSufficient: true,
		Payouts:      []generate.PayoutCandidateWithBondAmountAndFee{payoutCandidate},
	}
	ofc := generate.OnFeesCollectionHookData{
		580,
		[]generate.PayoutCandidateWithBondAmountAndFee{payoutCandidate},
	}

	simulatedCandidate := generate.PayoutCandidateSimulated{
		PayoutCandidateWithBondAmountAndFee: payoutCandidate,
		SimulationResult: &common.OpLimits{
			AllocationBurn:          1,
			StorageBurn:             1,
			TransactionFee:          1,
			StorageLimit:            1,
			GasLimit:                1,
			DeserializationGasLimit: 1,
		},
	}

	t, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00+00:00")
	apg := generate.AfterPayoutsBlueprintGeneratedHookData{
		Cycle: 1,
		Payouts: []common.PayoutRecipe{
			simulatedCandidate.ToPayoutRecipe(mavryk.ZeroAddress, 1, enums.PAYOUT_KIND_DELEGATOR_REWARD),
		},
		Summary: common.CyclePayoutSummary{
			Cycle:              1,
			Delegators:         2,
			PaidDelegators:     1,
			OwnStakedBalance:   mavryk.NewZ(1000000000),
			EarnedFees:         mavryk.NewZ(1000000000),
			EarnedRewards:      mavryk.NewZ(1000000000),
			DistributedRewards: mavryk.NewZ(1000000000),
			BondIncome:         mavryk.NewZ(1000000000),
			FeeIncome:          mavryk.NewZ(1000000000),
			IncomeTotal:        mavryk.NewZ(1000000000),
			DonatedBonds:       mavryk.NewZ(1000000000),
			DonatedFees:        mavryk.NewZ(1000000000),
			DonatedTotal:       mavryk.NewZ(1000000000),
			Timestamp:          t,
		},
	}

	app := prepare.AfterPayoutsPreapered{
		ValidPayouts: []common.PayoutRecipe{
			simulatedCandidate.ToPayoutRecipe(mavryk.ZeroAddress, 1, enums.PAYOUT_KIND_DELEGATOR_REWARD),
		},
		ReportsOfPastSuccesfulPayouts: common.NewSuccessBatchResult([]common.PayoutRecipe{simulatedCandidate.ToPayoutRecipe(mavryk.ZeroAddress, 1, enums.PAYOUT_KIND_DELEGATOR_REWARD)}, mavryk.ZeroOpHash).ToReports(),
	}

	result := "\n"
	result += "NOTE: *all bellow examples are just sample data to showcase fields used in data passed to hooks.*\n\n"

	result += fmt.Sprintf("## %s\n\n", enums.EXTENSION_HOOK_AFTER_CANDIDATES_GENERATED)
	result += "This hook is capable of mutating data.\n"
	result += "```json\n"
	acgSerialized, _ := json.MarshalIndent(acg, "", "  ")
	result += string(acgSerialized)
	result += "\n```\n\n"

	result += fmt.Sprintf("## %s\n\n", enums.EXTENSION_HOOK_AFTER_BONDS_DISTRIBUTED)
	result += "This hook is capable of mutating data.\n"
	result += "```json\n"
	abdSerialized, _ := json.MarshalIndent(abd, "", "  ")
	result += string(abdSerialized)
	result += "\n```\n\n"

	result += fmt.Sprintf("## %s\n\n", enums.EXTENSION_HOOK_CHECK_BALANCE)
	result += "This hook is NOT capable of mutating data.\n"
	result += "```json\n"
	acbSerialized, _ := json.MarshalIndent(acb, "", "  ")
	result += string(acbSerialized)
	result += "\n```\n\n"

	result += fmt.Sprintf("## %s\n\n", enums.EXTENSION_HOOK_ON_FEES_COLLECTION)
	result += "This hook is capable of mutating data.\n"
	result += "```json\n"
	ofcSerialized, _ := json.MarshalIndent(ofc, "", "  ")
	result += string(ofcSerialized)
	result += "\n```\n\n"

	result += fmt.Sprintf("## %s\n\n", enums.EXTENSION_HOOK_AFTER_PAYOUTS_BLUEPRINT_GENERATED)
	result += "This hook is NOT capable of mutating data *currently*.\n"
	result += "```json\n"
	apgSerialized, _ := json.MarshalIndent(apg, "", "  ")
	result += string(apgSerialized)
	result += "\n```\n\n"

	result += fmt.Sprintf("## %s\n\n", enums.EXTENSION_HOOK_AFTER_PAYOUTS_PREPARED)
	result += "This hook is capable of mutating data *currently*.\n"
	result += "```json\n"
	appSerialized, _ := json.MarshalIndent(app, "", "  ")
	result += string(appSerialized)
	result += "\n```\n\n"

	// write to docs/extensions/Hooks.md
	os.WriteFile("docs/extensions/Hooks.md", []byte(result), 0644)
}

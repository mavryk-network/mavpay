package main

import (
	"encoding/json"
	"fmt"
	"os"

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
				Source:           mavryk.ZeroAddress,
				Recipient:        mavryk.ZeroAddress,
				FeeRate:          5.0,
				DelegatedBalance: mavryk.NewZ(1000000000),
				StakedBalance:    mavryk.NewZ(1000000000),
				IsInvalid:        true,
				IsEmptied:        true,
				InvalidBecause:   "reason",
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
	ofc := generate.OnFeesCollectionHookData{
		580,
		[]generate.PayoutCandidateWithBondAmountAndFee{payoutCandidate},
	}

	apg := generate.AfterPayoutsBlueprintGeneratedHookData{
		Cycle: 1,
		Payouts: []common.PayoutRecipe{
			payoutCandidate.ToPayoutRecipe(mavryk.ZeroAddress, 1, enums.PAYOUT_KIND_DELEGATOR_REWARD),
		},
		OwnStakedBalance: mavryk.NewZ(1000000000),
		EarnedBlockFees:  mavryk.NewZ(1000000000),
		EarnedRewards:    mavryk.NewZ(1000000000),
		EarnedTotal:      mavryk.NewZ(2000000000),
		BondIncome:       mavryk.NewZ(1000000000),
		DonatedBonds:     mavryk.NewZ(1000000000),
		// DonatedFees:      mavryk.NewZ(1000000000),
		// DonatedTotal:     mavryk.NewZ(1000000000),
	}

	recipe := payoutCandidate.ToPayoutRecipe(mavryk.ZeroAddress, 1, enums.PAYOUT_KIND_DELEGATOR_REWARD)

	acb := prepare.CheckBalanceHookData{
		SkipMavCheck: true,
		Message:      "This message is used to carry errors from hook to the caller.",
		IsSufficient: true,
		Payouts:      []*common.AccumulatedPayoutRecipe{recipe.AsAccumulated()},
	}
	app := prepare.AfterPayoutsPreapered{
		Payouts: []common.PayoutRecipe{
			recipe,
		},
		ReportsOfPastSuccesfulPayouts: common.NewSuccessBatchResult([]*common.AccumulatedPayoutRecipe{recipe.AsAccumulated()}, mavryk.ZeroOpHash).ToIndividualReports(),
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

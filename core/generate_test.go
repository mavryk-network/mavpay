package core

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mavpay/state"
	"github.com/mavryk-network/mavpay/test/mock"
	"github.com/mavryk-network/mavpay/utils"
	"github.com/mavryk-network/gomavryk/mavryk"
)

type mockGenerateCollector struct {
	mock.EmptyCollector
}

func (engine *mockGenerateCollector) GetId() string {
	return "mockGenerateCollector"
}

func (engine *mockGenerateCollector) IsRevealed(address mavryk.Address) (bool, error) {
	return true, nil
}

func (engine *mockGenerateCollector) GetCycleStakingData(baker mavryk.Address, cycle int64) (*common.BakersCycleData, error) {
	rawCycleData := `{"OwnDelegatedBalance":"275708698","ExternalDelegatedBalance":"49100747788","BlockDelegatedRewards":"1197688","IdealBlockDelegatedRewards":"1197688","AttestationsDelegatedRewards":"1920302","IdealAttestationsDelegatedRewards":"1920302","DalDelegatedRewards":"427318","IdealDalDelegatedRewards":"427318","BlockDelegatedFees":"17100","DelegatorsCount":57,"OwnStakedBalance":"16421212933","ExternalStakedBalance":"29383795329","BlockStakingRewardsEdge":"192434","AttestationStakingRewardsEdge":"308534","BlockStakingFees":"0","StakersCount":8,"FrozenDepositLimit":"7300000000","Delegators":[{"Address":"mv1XHWTsUduZo49c36XvUeBggTVG4FCGz12p","DelegatedBalance":"18996459719","StakedBalance":"0","Emptied":false},{"Address":"mv1X9TXhN3ywRxytgkDvGq6Bcaf7smz4chYU","DelegatedBalance":"16995702660","StakedBalance":"0","Emptied":false},{"Address":"mv1X2SWH5ewCXYzTE8G37tw9pHAz8LUuceJj","DelegatedBalance":"6704943203","StakedBalance":"0","Emptied":false},{"Address":"mv1WfXzCwhHES8YcDaJDArzxhCLQLQFxYQGU","DelegatedBalance":"2067404437","StakedBalance":"0","Emptied":false},{"Address":"mv1NRwexo4LRtPixUXhUQLbQyDnWnQ682GTd","DelegatedBalance":"767347858","StakedBalance":"0","Emptied":false},{"Address":"mv1Wc4TxW2pNrZe4fvUh6Ux9Z6bdSdCJgtte","DelegatedBalance":"664759189","StakedBalance":"0","Emptied":false},{"Address":"mv1WV8nm6rnkpk6inA8Eqe74Gr36xkmhy9aC","DelegatedBalance":"598755901","StakedBalance":"0","Emptied":false},{"Address":"mv1WLJr1tSbZ1AcxNLN9BaocXw654EtaXUGX","DelegatedBalance":"507845869","StakedBalance":"0","Emptied":false},{"Address":"mv1WG5a7s8o8KU5x1vVGgWHWEBQvLveB766Q","DelegatedBalance":"405103709","StakedBalance":"0","Emptied":false},{"Address":"mv1WCTuyzCrBZMtHEJKXcxR9RWNCEBfoyM3Z","DelegatedBalance":"313906924","StakedBalance":"0","Emptied":false},{"Address":"mv1Vv3mBK3Q166DndgzfaMux48Ye5Hmq9oGK","DelegatedBalance":"170719784","StakedBalance":"0","Emptied":false},{"Address":"mv1VnvaH7sWiY6jRX57VCFnpPcax3FnQ2zjb","DelegatedBalance":"165652388","StakedBalance":"0","Emptied":false},{"Address":"mv1VWJUsAUkWhXN4EqdeRQEHF1jobTN4KjMc","DelegatedBalance":"161298995","StakedBalance":"0","Emptied":false},{"Address":"mv1VJqqe9A4otkv32xZZuj663PszaBxXAvV3","DelegatedBalance":"94164895","StakedBalance":"0","Emptied":false},{"Address":"mv1V7DmkL7ApfzW9x4Z888K9SBwMvhY12S3G","DelegatedBalance":"75318303","StakedBalance":"0","Emptied":false},{"Address":"mv1V4ntsKBfPWueYCSVjBLA8WsRwr89otXqC","DelegatedBalance":"74317010","StakedBalance":"0","Emptied":false},{"Address":"mv1V1jSM6s66pPjoSFh2ysafeJubmQuH6EPy","DelegatedBalance":"70644509","StakedBalance":"0","Emptied":false},{"Address":"mv1UwtfjoqsW2VSnh6ua9bXWqQXC75Hc1sRT","DelegatedBalance":"50000001","StakedBalance":"0","Emptied":false},{"Address":"mv1UqEZgmL7M5X836tX5eU3ToTrrBeSwpsQ5","DelegatedBalance":"34441415","StakedBalance":"0","Emptied":false},{"Address":"mv1UaZuPQcBSZ6vBk8dqL8uLpMJvc4rFEwJt","DelegatedBalance":"32415471","StakedBalance":"0","Emptied":false},{"Address":"mv1UZF5SdgTcosigsnovTEbgUaFEiNCsG9NJ","DelegatedBalance":"20256873","StakedBalance":"0","Emptied":false},{"Address":"mv1UYABMCi3JHWKhsjA5uqjQkX6VpbbMRg6V","DelegatedBalance":"18888009","StakedBalance":"0","Emptied":false},{"Address":"mv1UVKnfHCrzbUcfHS63xFeYbD4PfBCNAfYw","DelegatedBalance":"17427981","StakedBalance":"0","Emptied":false},{"Address":"mv1UPcgbYJCHWj6xg15VD2z1mAPMHH1VY3Sn","DelegatedBalance":"14207578","StakedBalance":"0","Emptied":false},{"Address":"mv1U7MzxbVWSVnjYWMwvrkWxReQwSzPkXcvB","DelegatedBalance":"12750656","StakedBalance":"0","Emptied":false},{"Address":"mv1TnGFAE2ehvQe5F2U5HbWZWb2FX423Vd8i","DelegatedBalance":"11966885","StakedBalance":"0","Emptied":false},{"Address":"mv1Tf75a2fUGmrThNuyse4vKFbPWM697SrNY","DelegatedBalance":"11307149","StakedBalance":"0","Emptied":false},{"Address":"mv1Te9TpceACAx9ovhntt2VdsvkpLEDqvBvn","DelegatedBalance":"10889299","StakedBalance":"0","Emptied":false},{"Address":"mv1TWUZkUjutJHgo9t7FqxTNiqoFPuGEC4Bq","DelegatedBalance":"7355283","StakedBalance":"0","Emptied":false},{"Address":"mv1TUaXztR5AK1BiRBxkFVLL49rkYLY3UqvZ","DelegatedBalance":"5221495","StakedBalance":"0","Emptied":false},{"Address":"mv1TSNZFAJaf8EQw4wreZjXtrcqZ8h4DFguU","DelegatedBalance":"2939210","StakedBalance":"0","Emptied":false},{"Address":"mv1TNgdUnE5udE6azQKy4HmDQJUWnKq2XQHu","DelegatedBalance":"1984760","StakedBalance":"0","Emptied":false},{"Address":"mv1TGgFyRj2X6HovN3bCVSzMNZpJAuZGfmah","DelegatedBalance":"1866216","StakedBalance":"0","Emptied":false},{"Address":"mv1SqGXUmNrBnnxE8zGueLteyK8nGf8MEX3A","DelegatedBalance":"1641573","StakedBalance":"0","Emptied":false},{"Address":"mv1SpdbvjSeBH2sVGbwWdS6PjT3fTjcJXrND","DelegatedBalance":"1502458","StakedBalance":"0","Emptied":false},{"Address":"mv1SoxkkpjGAxigEX9jYEZjnhytAwYiJQXSM","DelegatedBalance":"1459627","StakedBalance":"0","Emptied":false},{"Address":"mv1SXRSyhLnM9gSCc4YK3Ve6usdUN3pxNUtP","DelegatedBalance":"1425001","StakedBalance":"0","Emptied":false},{"Address":"mv1SEq4K9xCso46WdqhDTPLHU6intuatA5Rz","DelegatedBalance":"1406177","StakedBalance":"0","Emptied":false},{"Address":"mv1S8DRaNqHFRMNkpPgz3m5VnDGNSp4g4E71","DelegatedBalance":"1322354","StakedBalance":"0","Emptied":false},{"Address":"mv1RJjoUbdiZhG5ifQteRXVMPbooozCXwSbQ","DelegatedBalance":"990125","StakedBalance":"0","Emptied":false},{"Address":"mv1REFFS2cy5hwqiDtxgeQmMpWTgoJdTwnBQ","DelegatedBalance":"960009","StakedBalance":"0","Emptied":false},{"Address":"mv1QuSnBhYgaYwSq9Xy2985c5paAsFp9wBC5","DelegatedBalance":"370571","StakedBalance":"0","Emptied":false},{"Address":"KT1AmQTRDjTwfJDASJRJZdd7uJwDqs5W2mjA","DelegatedBalance":"289285","StakedBalance":"0","Emptied":false},{"Address":"mv1QtzoiuEztATbwANYP4gtEAdczn5ddV4Nj","DelegatedBalance":"275000","StakedBalance":"0","Emptied":false},{"Address":"mv1QmgAgBAMDHrx5th182efP3oygKWr1C2uU","DelegatedBalance":"200001","StakedBalance":"0","Emptied":false},{"Address":"mv1Qa4KUBhQKbwypURu3DCZVd1eyfAGwiVWp","DelegatedBalance":"172996","StakedBalance":"0","Emptied":false},{"Address":"KT1Kmai449TQT76GZXbihwNFHTUy432y1Z6Y","DelegatedBalance":"164629","StakedBalance":"0","Emptied":false},{"Address":"mv1QVVWyMc5yUJSja4ekcSceiw2kr8QnMZbF","DelegatedBalance":"162848","StakedBalance":"0","Emptied":false},{"Address":"mv1QFwxFkC9c7fezjNE4VPATXfefbam1GFAt","DelegatedBalance":"63473","StakedBalance":"0","Emptied":false},{"Address":"mv1Pf1ZpN3v4fBW3oSYhTj4wnf9XTEQATfDK","DelegatedBalance":"40787","StakedBalance":"0","Emptied":false},{"Address":"mv1PXgWu6bukpMfSCMgmLHqPvwCxUFkVuT3A","DelegatedBalance":"27475","StakedBalance":"0","Emptied":false},{"Address":"mv1PM6eS9LcrNZGQGiAuistbK64odUvw5CqG","DelegatedBalance":"8937","StakedBalance":"0","Emptied":false},{"Address":"KT19XE62UbrJ2gWW4ZWq2UxTQLhrnjBHLvBm","DelegatedBalance":"826","StakedBalance":"0","Emptied":false},{"Address":"mv1PJLDRgEZGwxwCJsjM4hv3Rhx5AMbYuaJa","DelegatedBalance":"1","StakedBalance":"0","Emptied":false},{"Address":"mv1PHN4Ck48rVjhj35FGWq23vYVSDKRD5sny","DelegatedBalance":"1","StakedBalance":"0","Emptied":false},{"Address":"KT1FtGbyLR1KV9oQGEYgBUpKPkEC8BcQn4cD","DelegatedBalance":"0","StakedBalance":"0","Emptied":false},{"Address":"KT1B5KPckWy2Mw99ii3wKuE4TQWKKQtSNXFE","DelegatedBalance":"0","StakedBalance":"0","Emptied":false}]}`
	var cycleData common.BakersCycleData
	err := json.Unmarshal([]byte(rawCycleData), &cycleData)
	if err != nil {
		panic(err)
	}
	return &cycleData, err
}

func assertBlueprintsEqual(assert *assert.Assertions, expected, actual *common.CyclePayoutBlueprint) {
	assert.Equal(len(expected.Payouts), len(actual.Payouts))
	assert.Equal(expected.Cycle, actual.Cycle)
	assert.Equal(expected.OwnStakedBalance, actual.OwnStakedBalance)
	assert.Equal(expected.OwnDelegatedBalance, actual.OwnDelegatedBalance)
	assert.Equal(expected.ExternalStakedBalance, actual.ExternalStakedBalance)
	assert.Equal(expected.ExternalDelegatedBalance, actual.ExternalDelegatedBalance)
	assert.Equal(expected.EarnedBlockFees, actual.EarnedBlockFees)
	assert.Equal(expected.EarnedRewards, actual.EarnedRewards)
	assert.Equal(expected.EarnedTotal, actual.EarnedTotal)
	assert.Equal(expected.BondIncome, actual.BondIncome)
	assert.Equal(expected.FeeIncome, actual.FeeIncome)
	assert.Equal(expected.IncomeTotal, actual.IncomeTotal)
	assert.Equal(expected.DonatedBonds, actual.DonatedBonds)
	assert.Equal(expected.DonatedFees, actual.DonatedFees)
	assert.Equal(expected.DonatedTotal, actual.DonatedTotal)

	utils.SortPayouts(expected.Payouts)
	utils.SortPayouts(actual.Payouts)
	for i, payout := range expected.Payouts {
		if payout.Recipient == mavryk.MustParseAddress("mv1VnvaH7sWiY6jRX57VCFnpPcax3FnQ2zjb") {
			// skip address which was used to generate test data
			// this address was used to generate test data, we do not use it in tests
			// so if it is delegated to baker it would appear in result as RECIPIENT_TARGETS_PAYOUT
			// which can not be replicated in test data as we use a random payout address for tests
			continue
		}

		assert.Equal(expected.Payouts[i].Baker, actual.Payouts[i].Baker)
		assert.Equal(expected.Payouts[i].Delegator, actual.Payouts[i].Delegator)
		assert.Equal(expected.Payouts[i].Cycle, actual.Payouts[i].Cycle)
		assert.Equal(expected.Payouts[i].Recipient, actual.Payouts[i].Recipient)
		assert.Equal(expected.Payouts[i].Kind, actual.Payouts[i].Kind)
		assert.Equal(expected.Payouts[i].TxKind, actual.Payouts[i].TxKind)
		assert.Equal(expected.Payouts[i].FATokenId, actual.Payouts[i].FATokenId)
		assert.Equal(expected.Payouts[i].FAContract, actual.Payouts[i].FAContract)
		assert.Equal(expected.Payouts[i].FAAlias, actual.Payouts[i].FAAlias)
		assert.Equal(expected.Payouts[i].FADecimals, actual.Payouts[i].FADecimals)
		assert.Equal(expected.Payouts[i].DelegatedBalance, actual.Payouts[i].DelegatedBalance)
		assert.Equal(expected.Payouts[i].StakedBalance, actual.Payouts[i].StakedBalance)
		assert.Equal(expected.Payouts[i].Amount, actual.Payouts[i].Amount)
		assert.Equal(expected.Payouts[i].FeeRate, actual.Payouts[i].FeeRate)
		assert.Equal(expected.Payouts[i].Fee, actual.Payouts[i].Fee)
		assert.Equal(expected.Payouts[i].TxFee, actual.Payouts[i].TxFee)
		assert.Equal(expected.Payouts[i].Note, actual.Payouts[i].Note)
		assert.Equal(expected.Payouts[i].IsValid, actual.Payouts[i].IsValid)
	}
}

func Test_Generate(t *testing.T) {
	var err error
	assert := assert.New(t)

	state.Init(".", state.StateInitOptions{})
	rawConfig := `{
        mavpay_config_version: 0
        baker: mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs
        payouts: {
                fee: 0.11
                transaction_fee_buffer: 0
                kt_transaction_fee_buffer: 0
        }
        delegators: {
                requirements: {
                        minimum_balance: 10
                }
        }
        network: {
                rpc_pool: [
					https://rpc.mavryk.network/
                ]
                mvkt_url: https://api.mavryk.network/
        }
        overdelegation: {
                protect: true
        }
        income_recipients: {
                donate: 0.01
        }
        notifications: [
        ]
	}`

	config, _ := configuration.LoadFromString([]byte(rawConfig))
	assert.NotNil(config)

	collector := mockGenerateCollector{}
	signer := mock.InitSimpleSigner()

	engineContext := common.NewGeneratePayoutsEngines(&collector, signer, func(msg string) {})

	result, err := GeneratePayouts(config, engineContext, &common.GeneratePayoutsOptions{
		Cycle: 1016,
	})

	assert.Nil(err)
	assert.NotNil(result)
	fmt.Println(len(result.Payouts))

	var expectedResult common.CyclePayoutBlueprint = common.CyclePayoutBlueprint{
		Cycle:                    1016,
		OwnStakedBalance:         mavryk.NewZ(16421212933),
		OwnDelegatedBalance:      mavryk.NewZ(275708698),
		ExternalStakedBalance:    mavryk.NewZ(29383795329),
		ExternalDelegatedBalance: mavryk.NewZ(49100747788),
		EarnedBlockFees:          mavryk.NewZ(17100),
		EarnedRewards:            mavryk.NewZ(3545308),
		EarnedTotal:              mavryk.NewZ(3562408),
		BondIncome:               mavryk.NewZ(19693),
		FeeIncome:                mavryk.NewZ(385753),
		IncomeTotal:              mavryk.NewZ(405446),
		DonatedBonds:             mavryk.NewZ(198),
		DonatedFees:              mavryk.NewZ(3896),
		DonatedTotal:             mavryk.NewZ(4094),
		Payouts: []common.PayoutRecipe{
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1XHWTsUduZo49c36XvUeBggTVG4FCGz12p"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1XHWTsUduZo49c36XvUeBggTVG4FCGz12p"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(18996459719),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(1219794),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(150761),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1X9TXhN3ywRxytgkDvGq6Bcaf7smz4chYU"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1X9TXhN3ywRxytgkDvGq6Bcaf7smz4chYU"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(16995702660),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(1091322),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(134882),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1X2SWH5ewCXYzTE8G37tw9pHAz8LUuceJj"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1X2SWH5ewCXYzTE8G37tw9pHAz8LUuceJj"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(6704943203),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(430535),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(53212),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1WfXzCwhHES8YcDaJDArzxhCLQLQFxYQGU"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1WfXzCwhHES8YcDaJDArzxhCLQLQFxYQGU"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(2067404437),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(132751),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(16407),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1NRwexo4LRtPixUXhUQLbQyDnWnQ682GTd"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1NRwexo4LRtPixUXhUQLbQyDnWnQ682GTd"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(767347858),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(49273),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(6089),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1Wc4TxW2pNrZe4fvUh6Ux9Z6bdSdCJgtte"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1Wc4TxW2pNrZe4fvUh6Ux9Z6bdSdCJgtte"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(664759189),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(42685),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(5275),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1WV8nm6rnkpk6inA8Eqe74Gr36xkmhy9aC"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1WV8nm6rnkpk6inA8Eqe74Gr36xkmhy9aC"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(598755901),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(38447),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(4751),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1WLJr1tSbZ1AcxNLN9BaocXw654EtaXUGX"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1WLJr1tSbZ1AcxNLN9BaocXw654EtaXUGX"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(507845869),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(32610),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(4030),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1WG5a7s8o8KU5x1vVGgWHWEBQvLveB766Q"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1WG5a7s8o8KU5x1vVGgWHWEBQvLveB766Q"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(405103709),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(26013),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(3214),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1WCTuyzCrBZMtHEJKXcxR9RWNCEBfoyM3Z"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1WCTuyzCrBZMtHEJKXcxR9RWNCEBfoyM3Z"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(313906924),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(20156),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(2491),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1Vv3mBK3Q166DndgzfaMux48Ye5Hmq9oGK"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1Vv3mBK3Q166DndgzfaMux48Ye5Hmq9oGK"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(170719784),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(10963),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(1354),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1VnvaH7sWiY6jRX57VCFnpPcax3FnQ2zjb"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1VnvaH7sWiY6jRX57VCFnpPcax3FnQ2zjb"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(165652388),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(10637),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(1314),
				Note:             "RECIPIENT_TARGETS_PAYOUT",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1VWJUsAUkWhXN4EqdeRQEHF1jobTN4KjMc"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1VWJUsAUkWhXN4EqdeRQEHF1jobTN4KjMc"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(161298995),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(10357),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(1280),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1VJqqe9A4otkv32xZZuj663PszaBxXAvV3"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1VJqqe9A4otkv32xZZuj663PszaBxXAvV3"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(94164895),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(6046),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(747),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1V7DmkL7ApfzW9x4Z888K9SBwMvhY12S3G"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1V7DmkL7ApfzW9x4Z888K9SBwMvhY12S3G"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(75318303),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(4837),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(597),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1V4ntsKBfPWueYCSVjBLA8WsRwr89otXqC"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1V4ntsKBfPWueYCSVjBLA8WsRwr89otXqC"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(74317010),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(4772),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(589),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1V1jSM6s66pPjoSFh2ysafeJubmQuH6EPy"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1V1jSM6s66pPjoSFh2ysafeJubmQuH6EPy"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(70644509),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(4536),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(560),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1UwtfjoqsW2VSnh6ua9bXWqQXC75Hc1sRT"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1UwtfjoqsW2VSnh6ua9bXWqQXC75Hc1sRT"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(50000001),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(3211),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(396),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1UqEZgmL7M5X836tX5eU3ToTrrBeSwpsQ5"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1UqEZgmL7M5X836tX5eU3ToTrrBeSwpsQ5"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(34441415),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(2211),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(273),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1UaZuPQcBSZ6vBk8dqL8uLpMJvc4rFEwJt"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1UaZuPQcBSZ6vBk8dqL8uLpMJvc4rFEwJt"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(32415471),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(2081),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(257),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1UZF5SdgTcosigsnovTEbgUaFEiNCsG9NJ"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1UZF5SdgTcosigsnovTEbgUaFEiNCsG9NJ"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(20256873),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(1301),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(160),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1UYABMCi3JHWKhsjA5uqjQkX6VpbbMRg6V"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1UYABMCi3JHWKhsjA5uqjQkX6VpbbMRg6V"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(18888009),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(1213),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(149),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1UVKnfHCrzbUcfHS63xFeYbD4PfBCNAfYw"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1UVKnfHCrzbUcfHS63xFeYbD4PfBCNAfYw"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(17427981),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(1119),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(138),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1UPcgbYJCHWj6xg15VD2z1mAPMHH1VY3Sn"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1UPcgbYJCHWj6xg15VD2z1mAPMHH1VY3Sn"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(14207578),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(913),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(112),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1U7MzxbVWSVnjYWMwvrkWxReQwSzPkXcvB"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1U7MzxbVWSVnjYWMwvrkWxReQwSzPkXcvB"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(12750656),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(818),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(101),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1TnGFAE2ehvQe5F2U5HbWZWb2FX423Vd8i"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1TnGFAE2ehvQe5F2U5HbWZWb2FX423Vd8i"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(11966885),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(769),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(94),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1Tf75a2fUGmrThNuyse4vKFbPWM697SrNY"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1Tf75a2fUGmrThNuyse4vKFbPWM697SrNY"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(11307149),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(726),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(89),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1Te9TpceACAx9ovhntt2VdsvkpLEDqvBvn"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1Te9TpceACAx9ovhntt2VdsvkpLEDqvBvn"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(10889299),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(699),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(86),
				IsValid:          true,
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1TWUZkUjutJHgo9t7FqxTNiqoFPuGEC4Bq"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1TWUZkUjutJHgo9t7FqxTNiqoFPuGEC4Bq"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(7355283),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(472),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(58),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1TUaXztR5AK1BiRBxkFVLL49rkYLY3UqvZ"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1TUaXztR5AK1BiRBxkFVLL49rkYLY3UqvZ"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(5221495),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(335),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(41),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1TSNZFAJaf8EQw4wreZjXtrcqZ8h4DFguU"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1TSNZFAJaf8EQw4wreZjXtrcqZ8h4DFguU"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(2939210),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(189),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(23),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1TNgdUnE5udE6azQKy4HmDQJUWnKq2XQHu"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1TNgdUnE5udE6azQKy4HmDQJUWnKq2XQHu"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(1984760),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(128),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(15),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1TGgFyRj2X6HovN3bCVSzMNZpJAuZGfmah"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1TGgFyRj2X6HovN3bCVSzMNZpJAuZGfmah"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(1866216),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(120),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(14),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1SqGXUmNrBnnxE8zGueLteyK8nGf8MEX3A"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1SqGXUmNrBnnxE8zGueLteyK8nGf8MEX3A"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(1641573),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(106),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(12),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1SpdbvjSeBH2sVGbwWdS6PjT3fTjcJXrND"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1SpdbvjSeBH2sVGbwWdS6PjT3fTjcJXrND"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(1502458),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(97),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(11),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1SoxkkpjGAxigEX9jYEZjnhytAwYiJQXSM"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1SoxkkpjGAxigEX9jYEZjnhytAwYiJQXSM"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(1459627),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(94),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(11),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1SXRSyhLnM9gSCc4YK3Ve6usdUN3pxNUtP"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1SXRSyhLnM9gSCc4YK3Ve6usdUN3pxNUtP"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(1425001),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(91),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(11),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1SEq4K9xCso46WdqhDTPLHU6intuatA5Rz"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1SEq4K9xCso46WdqhDTPLHU6intuatA5Rz"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(1406177),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(90),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(11),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1S8DRaNqHFRMNkpPgz3m5VnDGNSp4g4E71"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1S8DRaNqHFRMNkpPgz3m5VnDGNSp4g4E71"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(1322354),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(85),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(10),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1RJjoUbdiZhG5ifQteRXVMPbooozCXwSbQ"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1RJjoUbdiZhG5ifQteRXVMPbooozCXwSbQ"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(990125),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(64),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(7),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1REFFS2cy5hwqiDtxgeQmMpWTgoJdTwnBQ"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1REFFS2cy5hwqiDtxgeQmMpWTgoJdTwnBQ"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(960009),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(62),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(7),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1QuSnBhYgaYwSq9Xy2985c5paAsFp9wBC5"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1QuSnBhYgaYwSq9Xy2985c5paAsFp9wBC5"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(370571),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(24),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(2),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("KT1AmQTRDjTwfJDASJRJZdd7uJwDqs5W2mjA"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("KT1AmQTRDjTwfJDASJRJZdd7uJwDqs5W2mjA"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(289285),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(18),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(2),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1QtzoiuEztATbwANYP4gtEAdczn5ddV4Nj"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1QtzoiuEztATbwANYP4gtEAdczn5ddV4Nj"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(275000),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(17),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(2),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1QmgAgBAMDHrx5th182efP3oygKWr1C2uU"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1QmgAgBAMDHrx5th182efP3oygKWr1C2uU"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(200001),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(13),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(1),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1Qa4KUBhQKbwypURu3DCZVd1eyfAGwiVWp"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1Qa4KUBhQKbwypURu3DCZVd1eyfAGwiVWp"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(172996),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(11),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(1),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("KT1Kmai449TQT76GZXbihwNFHTUy432y1Z6Y"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("KT1Kmai449TQT76GZXbihwNFHTUy432y1Z6Y"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(164629),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(10),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(1),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1QVVWyMc5yUJSja4ekcSceiw2kr8QnMZbF"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1QVVWyMc5yUJSja4ekcSceiw2kr8QnMZbF"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(162848),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(10),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(1),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1QFwxFkC9c7fezjNE4VPATXfefbam1GFAt"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1QFwxFkC9c7fezjNE4VPATXfefbam1GFAt"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(63473),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(4),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(0),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1Pf1ZpN3v4fBW3oSYhTj4wnf9XTEQATfDK"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1Pf1ZpN3v4fBW3oSYhTj4wnf9XTEQATfDK"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(40787),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(2),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(0),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1PXgWu6bukpMfSCMgmLHqPvwCxUFkVuT3A"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1PXgWu6bukpMfSCMgmLHqPvwCxUFkVuT3A"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(27475),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(1),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(0),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1PM6eS9LcrNZGQGiAuistbK64odUvw5CqG"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1PM6eS9LcrNZGQGiAuistbK64odUvw5CqG"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(8937),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(0),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(0),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("KT19XE62UbrJ2gWW4ZWq2UxTQLhrnjBHLvBm"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("KT19XE62UbrJ2gWW4ZWq2UxTQLhrnjBHLvBm"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(826),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(0),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(0),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1PJLDRgEZGwxwCJsjM4hv3Rhx5AMbYuaJa"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1PJLDRgEZGwxwCJsjM4hv3Rhx5AMbYuaJa"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(1),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(0),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(0),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("mv1PHN4Ck48rVjhj35FGWq23vYVSDKRD5sny"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1PHN4Ck48rVjhj35FGWq23vYVSDKRD5sny"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(1),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(0),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(0),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("KT1FtGbyLR1KV9oQGEYgBUpKPkEC8BcQn4cD"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("KT1FtGbyLR1KV9oQGEYgBUpKPkEC8BcQn4cD"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(0),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(0),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(0),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.MustParseAddress("KT1B5KPckWy2Mw99ii3wKuE4TQWKKQtSNXFE"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("KT1B5KPckWy2Mw99ii3wKuE4TQWKKQtSNXFE"),
				Kind:             "delegator reward",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(0),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(0),
				FeeRate:          0.11,
				Fee:              mavryk.NewZ(0),
				Note:             "DELEGATOR_LOW_BALANCE",
			},
			{
				Baker:            mavryk.MustParseAddress("mv1T9xoFWkkNgy6wH5xeDg9XgdwnqznpuDXs"),
				Delegator:        mavryk.Address{}, // Represents an empty delegator address
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("mv1WLJr1tSbZ1AcxNLN9BaocXw654EtaXUGX"),
				Kind:             "donation",
				TxKind:           "mav",
				FATokenId:        mavryk.NewZ(0),
				DelegatedBalance: mavryk.NewZ(0),
				StakedBalance:    mavryk.NewZ(0),
				Amount:           mavryk.NewZ(4094),
				Fee:              mavryk.NewZ(0),
				IsValid:          true,
			},
		},
	}

	assert.Equal(len(expectedResult.Payouts), len(result.Payouts))
	assert.Equal(expectedResult.Cycle, result.Cycle)
	assert.Equal(expectedResult.OwnStakedBalance, result.OwnStakedBalance)
	assert.Equal(expectedResult.OwnDelegatedBalance, result.OwnDelegatedBalance)
	assert.Equal(expectedResult.ExternalStakedBalance, result.ExternalStakedBalance)
	assert.Equal(expectedResult.ExternalDelegatedBalance, result.ExternalDelegatedBalance)
	assert.Equal(expectedResult.EarnedBlockFees, result.EarnedBlockFees)
	assert.Equal(expectedResult.EarnedRewards, result.EarnedRewards)
	assert.Equal(expectedResult.EarnedTotal, result.EarnedTotal)
	assert.Equal(expectedResult.BondIncome, result.BondIncome)
	assert.Equal(expectedResult.FeeIncome, result.FeeIncome)
	assert.Equal(expectedResult.IncomeTotal, result.IncomeTotal)
	assert.Equal(expectedResult.DonatedBonds, result.DonatedBonds)
	assert.Equal(expectedResult.DonatedFees, result.DonatedFees)
	assert.Equal(expectedResult.DonatedTotal, result.DonatedTotal)

	utils.SortPayouts(expectedResult.Payouts)
	utils.SortPayouts(result.Payouts)
	for i, payout := range expectedResult.Payouts {
		if payout.Recipient == mavryk.MustParseAddress("mv1VnvaH7sWiY6jRX57VCFnpPcax3FnQ2zjb") {
			// skip address which was used to generate test data
			// this address was used to generate test data, we do not use it in tests
			// so if it is delegated to baker it would appear in result as RECIPIENT_TARGETS_PAYOUT
			// which can not be replicated in test data as we use a random payout address for tests
			continue
		}

		assert.Equal(expectedResult.Payouts[i].Baker, result.Payouts[i].Baker)
		assert.Equal(expectedResult.Payouts[i].Delegator, result.Payouts[i].Delegator)
		assert.Equal(expectedResult.Payouts[i].Cycle, result.Payouts[i].Cycle)
		assert.Equal(expectedResult.Payouts[i].Recipient, result.Payouts[i].Recipient)
		assert.Equal(expectedResult.Payouts[i].Kind, result.Payouts[i].Kind)
		assert.Equal(expectedResult.Payouts[i].TxKind, result.Payouts[i].TxKind)
		assert.Equal(expectedResult.Payouts[i].FATokenId, result.Payouts[i].FATokenId)
		assert.Equal(expectedResult.Payouts[i].FAContract, result.Payouts[i].FAContract)
		assert.Equal(expectedResult.Payouts[i].FAAlias, result.Payouts[i].FAAlias)
		assert.Equal(expectedResult.Payouts[i].FADecimals, result.Payouts[i].FADecimals)
		assert.Equal(expectedResult.Payouts[i].DelegatedBalance, result.Payouts[i].DelegatedBalance)
		assert.Equal(expectedResult.Payouts[i].StakedBalance, result.Payouts[i].StakedBalance)
		assert.Equal(expectedResult.Payouts[i].Amount, result.Payouts[i].Amount)
		assert.Equal(expectedResult.Payouts[i].FeeRate, result.Payouts[i].FeeRate)
		assert.Equal(expectedResult.Payouts[i].Fee, result.Payouts[i].Fee)
		assert.Equal(expectedResult.Payouts[i].TxFee, result.Payouts[i].TxFee)
		assert.Equal(expectedResult.Payouts[i].Note, result.Payouts[i].Note)
		assert.Equal(expectedResult.Payouts[i].IsValid, result.Payouts[i].IsValid)
	}
}

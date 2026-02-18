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
	"github.com/mavryk-network/mvgo/mavryk"
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
	rawCycleData := `{"OwnDelegatedBalance":"275708698","ExternalDelegatedBalance":"49100747788","BlockDelegatedRewards":"1197688","IdealBlockDelegatedRewards":"1197688","AttestationsDelegatedRewards":"1920302","IdealAttestationsDelegatedRewards":"1920302","DalDelegatedRewards":"427318","IdealDalDelegatedRewards":"427318","BlockDelegatedFees":"17100","DelegatorsCount":57,"OwnStakedBalance":"16421212933","ExternalStakedBalance":"29383795329","BlockStakingRewardsEdge":"192434","AttestationStakingRewardsEdge":"308534","BlockStakingFees":"0","StakersCount":8,"FrozenDepositLimit":"7300000000","Delegators":[{"Address":"tz1dKeXpumUQD5aCgk3jb7i7psWiRkzqJQdE","DelegatedBalance":"18996459719","StakedBalance":"0","Emptied":false},{"Address":"tz1WZRZ6ciwRvkx5YVe5uVXKxrnzDQQevFCL","DelegatedBalance":"16995702660","StakedBalance":"0","Emptied":false},{"Address":"tz1WpPaE47vtUT56MZGgKVi4VQCgEF95mvng","DelegatedBalance":"6704943203","StakedBalance":"0","Emptied":false},{"Address":"tz2X42QzdRqoqeKCqZaGpoXD4VmzmwbuMJMb","DelegatedBalance":"2067404437","StakedBalance":"0","Emptied":false},{"Address":"tz1PABDPXbg9rqz62umCaQ8YhCUajmoYK2Cb","DelegatedBalance":"767347858","StakedBalance":"0","Emptied":false},{"Address":"tz1dhJPohjESC3zzRnaJH7krf5wZPwLm4oH7","DelegatedBalance":"664759189","StakedBalance":"0","Emptied":false},{"Address":"tz1gnuBF9TbBcgHPV2mUE96tBrW7PxqRmx1h","DelegatedBalance":"598755901","StakedBalance":"0","Emptied":false},{"Address":"tz1UGkfyrT9yBt6U5PV7Qeui3pt3a8jffoWv","DelegatedBalance":"507845869","StakedBalance":"0","Emptied":false},{"Address":"tz1XS2hdV3ygveRn1gxRJZgJ4yxFK1YnRrLn","DelegatedBalance":"405103709","StakedBalance":"0","Emptied":false},{"Address":"tz1P9A1noGEeoAQ6WNYM87BbJ3iaioorrws2","DelegatedBalance":"313906924","StakedBalance":"0","Emptied":false},{"Address":"tz1eYiA2SFVBKPK2UW24puEzbTiLHxMDjVtj","DelegatedBalance":"170719784","StakedBalance":"0","Emptied":false},{"Address":"tz1X7U9XxVz6NDxL4DSZhijME61PW45bYUJE","DelegatedBalance":"165652388","StakedBalance":"0","Emptied":false},{"Address":"tz1WWvSczbkko8p14qknj6d1KiK4T3ckQ33X","DelegatedBalance":"161298995","StakedBalance":"0","Emptied":false},{"Address":"tz1gKDahpLhd85yuJ6TEFAwrcnmcEG9gTkSi","DelegatedBalance":"94164895","StakedBalance":"0","Emptied":false},{"Address":"tz2BZibvVaRzZq5kUWz1rcrJychdbHmoNZJE","DelegatedBalance":"75318303","StakedBalance":"0","Emptied":false},{"Address":"tz1cXDqHit4wsNv6dFHggKxDUBdpTkAaizyx","DelegatedBalance":"74317010","StakedBalance":"0","Emptied":false},{"Address":"tz2NvuB26dtjUg11Gb49uocbYrJf26zN2QqW","DelegatedBalance":"70644509","StakedBalance":"0","Emptied":false},{"Address":"tz1XyRgFTiNNZU7WzG3LDBi1bXWYPupBNBL2","DelegatedBalance":"50000001","StakedBalance":"0","Emptied":false},{"Address":"tz1dLZQyaopB2MDMdQC5nRt3Yx6RArVRubRh","DelegatedBalance":"34441415","StakedBalance":"0","Emptied":false},{"Address":"tz1gMRKQWdhY5ZPUehpHofLswyL8Vexvfpbm","DelegatedBalance":"32415471","StakedBalance":"0","Emptied":false},{"Address":"tz2Lhty58UJvywseQH9vpGasHjuHtiitKZ56","DelegatedBalance":"20256873","StakedBalance":"0","Emptied":false},{"Address":"tz1hVZBTeTs9GZtpShYYuF92f7F2g4gWSQo9","DelegatedBalance":"18888009","StakedBalance":"0","Emptied":false},{"Address":"tz1Nry5YPddf4VFDyunXLR8b8AtpvGJLMLZA","DelegatedBalance":"17427981","StakedBalance":"0","Emptied":false},{"Address":"tz1boUGNrp8Z1WEAkqTHAPr5DennP7tQjF71","DelegatedBalance":"14207578","StakedBalance":"0","Emptied":false},{"Address":"tz1Sf9KWCfjN3nsPHB4cek2MXBuXq5qjBZrs","DelegatedBalance":"12750656","StakedBalance":"0","Emptied":false},{"Address":"tz1frE7ArsC1spvQGizY5gBbJBtjnzGCgQ3y","DelegatedBalance":"11966885","StakedBalance":"0","Emptied":false},{"Address":"tz1hbseHXGm2XB7ivo1r3mbpH29c4GZE1aV2","DelegatedBalance":"11307149","StakedBalance":"0","Emptied":false},{"Address":"tz1f8LD3RZ9SMZUgsHQsQAX85EApD2cTXKES","DelegatedBalance":"10889299","StakedBalance":"0","Emptied":false},{"Address":"tz1dTmMLY7tan1HKfW5e3KxJcfZput6UuTH6","DelegatedBalance":"7355283","StakedBalance":"0","Emptied":false},{"Address":"tz1SRqikWtR2KNpugCApreA9QeZuqh8swQMa","DelegatedBalance":"5221495","StakedBalance":"0","Emptied":false},{"Address":"tz1iKqZc4CfrPwjcD7kx8MZtFtfL6YysKnPv","DelegatedBalance":"2939210","StakedBalance":"0","Emptied":false},{"Address":"tz2DWwmauKy4Dkak6X8ePY4Z91EsM36miYxB","DelegatedBalance":"1984760","StakedBalance":"0","Emptied":false},{"Address":"tz1aCb6QKL7K4JxRTUTYjAER8qZH4zKAbigq","DelegatedBalance":"1866216","StakedBalance":"0","Emptied":false},{"Address":"tz1f4zxfTtjTKkrzzXNtQU73zGBcasW3xEBS","DelegatedBalance":"1641573","StakedBalance":"0","Emptied":false},{"Address":"tz1TG2bxjvmRMQakt8KVQ5S6CyLzMuSHx2qF","DelegatedBalance":"1502458","StakedBalance":"0","Emptied":false},{"Address":"tz1YibTr9kUchgG6wJjMwRar2NUfawxH48FT","DelegatedBalance":"1459627","StakedBalance":"0","Emptied":false},{"Address":"tz1YtK9qUJffq2SqjhaJEGbufVHEBWDLPrSV","DelegatedBalance":"1425001","StakedBalance":"0","Emptied":false},{"Address":"tz2DTjmLqMXSwhkpCFnV977vyhtYWbsP6nwo","DelegatedBalance":"1406177","StakedBalance":"0","Emptied":false},{"Address":"tz1eaCp9cc4FtTs3Wm6UgZZnEtZ9b1Peqewi","DelegatedBalance":"1322354","StakedBalance":"0","Emptied":false},{"Address":"tz1MZ7ypArEoYocyyfscGaneNKM7bADhnQC4","DelegatedBalance":"990125","StakedBalance":"0","Emptied":false},{"Address":"tz1Qzov1LwEvSfHonRHr2EwQGto7ExrH43or","DelegatedBalance":"960009","StakedBalance":"0","Emptied":false},{"Address":"tz28e4LURT9Kv6TXysL6szWYkKKqRmA1wDGq","DelegatedBalance":"370571","StakedBalance":"0","Emptied":false},{"Address":"KT1AmQTRDjTwfJDASJRJZdd7uJwDqs5W2mjA","DelegatedBalance":"289285","StakedBalance":"0","Emptied":false},{"Address":"tz1UC8X57hKRmTbgJWsoos9njMy2SW3bKAAh","DelegatedBalance":"275000","StakedBalance":"0","Emptied":false},{"Address":"tz1XX6ykvpfPGAkzeDaFPvvPqmRQoTwCvP4S","DelegatedBalance":"200001","StakedBalance":"0","Emptied":false},{"Address":"tz1P91PPAfukVdpgj153WqoLymrWDYCoyLRA","DelegatedBalance":"172996","StakedBalance":"0","Emptied":false},{"Address":"KT1Kmai449TQT76GZXbihwNFHTUy432y1Z6Y","DelegatedBalance":"164629","StakedBalance":"0","Emptied":false},{"Address":"tz2JvpjaCGsynphe78Xbyu2kHUU2Ym8rxi8v","DelegatedBalance":"162848","StakedBalance":"0","Emptied":false},{"Address":"tz1YDHyB7aBCvj5TCYiRQkjeVHfgtEYUDAd2","DelegatedBalance":"63473","StakedBalance":"0","Emptied":false},{"Address":"tz1VWKeeSqxdLsTfAyCSrwFVZnfgw4ZwaHWV","DelegatedBalance":"40787","StakedBalance":"0","Emptied":false},{"Address":"tz1Xfkbwwa9ewNvYnesSXEYWGgS5gXVSQQVj","DelegatedBalance":"27475","StakedBalance":"0","Emptied":false},{"Address":"tz1Mv79qqYx2QX1NXYawrSrya13T7QoxCba9","DelegatedBalance":"8937","StakedBalance":"0","Emptied":false},{"Address":"KT19XE62UbrJ2gWW4ZWq2UxTQLhrnjBHLvBm","DelegatedBalance":"826","StakedBalance":"0","Emptied":false},{"Address":"tz2H4arakxXuzfo8NBVqTNgL5NBZ9b497Pfv","DelegatedBalance":"1","StakedBalance":"0","Emptied":false},{"Address":"tz1WxZvegEFTFSA8mqHWsBq97S15kS34JXqu","DelegatedBalance":"1","StakedBalance":"0","Emptied":false},{"Address":"KT1FtGbyLR1KV9oQGEYgBUpKPkEC8BcQn4cD","DelegatedBalance":"0","StakedBalance":"0","Emptied":false},{"Address":"KT1B5KPckWy2Mw99ii3wKuE4TQWKKQtSNXFE","DelegatedBalance":"0","StakedBalance":"0","Emptied":false}]}`
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
		if payout.Recipient == mavryk.MustParseAddress("tz1X7U9XxVz6NDxL4DSZhijME61PW45bYUJE") {
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
        baker: tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1dKeXpumUQD5aCgk3jb7i7psWiRkzqJQdE"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1dKeXpumUQD5aCgk3jb7i7psWiRkzqJQdE"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1WZRZ6ciwRvkx5YVe5uVXKxrnzDQQevFCL"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1WZRZ6ciwRvkx5YVe5uVXKxrnzDQQevFCL"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1WpPaE47vtUT56MZGgKVi4VQCgEF95mvng"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1WpPaE47vtUT56MZGgKVi4VQCgEF95mvng"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz2X42QzdRqoqeKCqZaGpoXD4VmzmwbuMJMb"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz2X42QzdRqoqeKCqZaGpoXD4VmzmwbuMJMb"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1PABDPXbg9rqz62umCaQ8YhCUajmoYK2Cb"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1PABDPXbg9rqz62umCaQ8YhCUajmoYK2Cb"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1dhJPohjESC3zzRnaJH7krf5wZPwLm4oH7"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1dhJPohjESC3zzRnaJH7krf5wZPwLm4oH7"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1gnuBF9TbBcgHPV2mUE96tBrW7PxqRmx1h"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1gnuBF9TbBcgHPV2mUE96tBrW7PxqRmx1h"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1UGkfyrT9yBt6U5PV7Qeui3pt3a8jffoWv"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1UGkfyrT9yBt6U5PV7Qeui3pt3a8jffoWv"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1XS2hdV3ygveRn1gxRJZgJ4yxFK1YnRrLn"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1XS2hdV3ygveRn1gxRJZgJ4yxFK1YnRrLn"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1P9A1noGEeoAQ6WNYM87BbJ3iaioorrws2"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1P9A1noGEeoAQ6WNYM87BbJ3iaioorrws2"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1eYiA2SFVBKPK2UW24puEzbTiLHxMDjVtj"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1eYiA2SFVBKPK2UW24puEzbTiLHxMDjVtj"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1X7U9XxVz6NDxL4DSZhijME61PW45bYUJE"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1X7U9XxVz6NDxL4DSZhijME61PW45bYUJE"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1WWvSczbkko8p14qknj6d1KiK4T3ckQ33X"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1WWvSczbkko8p14qknj6d1KiK4T3ckQ33X"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1gKDahpLhd85yuJ6TEFAwrcnmcEG9gTkSi"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1gKDahpLhd85yuJ6TEFAwrcnmcEG9gTkSi"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz2BZibvVaRzZq5kUWz1rcrJychdbHmoNZJE"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz2BZibvVaRzZq5kUWz1rcrJychdbHmoNZJE"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1cXDqHit4wsNv6dFHggKxDUBdpTkAaizyx"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1cXDqHit4wsNv6dFHggKxDUBdpTkAaizyx"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz2NvuB26dtjUg11Gb49uocbYrJf26zN2QqW"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz2NvuB26dtjUg11Gb49uocbYrJf26zN2QqW"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1XyRgFTiNNZU7WzG3LDBi1bXWYPupBNBL2"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1XyRgFTiNNZU7WzG3LDBi1bXWYPupBNBL2"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1dLZQyaopB2MDMdQC5nRt3Yx6RArVRubRh"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1dLZQyaopB2MDMdQC5nRt3Yx6RArVRubRh"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1gMRKQWdhY5ZPUehpHofLswyL8Vexvfpbm"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1gMRKQWdhY5ZPUehpHofLswyL8Vexvfpbm"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz2Lhty58UJvywseQH9vpGasHjuHtiitKZ56"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz2Lhty58UJvywseQH9vpGasHjuHtiitKZ56"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1hVZBTeTs9GZtpShYYuF92f7F2g4gWSQo9"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1hVZBTeTs9GZtpShYYuF92f7F2g4gWSQo9"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1Nry5YPddf4VFDyunXLR8b8AtpvGJLMLZA"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1Nry5YPddf4VFDyunXLR8b8AtpvGJLMLZA"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1boUGNrp8Z1WEAkqTHAPr5DennP7tQjF71"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1boUGNrp8Z1WEAkqTHAPr5DennP7tQjF71"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1Sf9KWCfjN3nsPHB4cek2MXBuXq5qjBZrs"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1Sf9KWCfjN3nsPHB4cek2MXBuXq5qjBZrs"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1frE7ArsC1spvQGizY5gBbJBtjnzGCgQ3y"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1frE7ArsC1spvQGizY5gBbJBtjnzGCgQ3y"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1hbseHXGm2XB7ivo1r3mbpH29c4GZE1aV2"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1hbseHXGm2XB7ivo1r3mbpH29c4GZE1aV2"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1f8LD3RZ9SMZUgsHQsQAX85EApD2cTXKES"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1f8LD3RZ9SMZUgsHQsQAX85EApD2cTXKES"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1dTmMLY7tan1HKfW5e3KxJcfZput6UuTH6"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1dTmMLY7tan1HKfW5e3KxJcfZput6UuTH6"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1SRqikWtR2KNpugCApreA9QeZuqh8swQMa"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1SRqikWtR2KNpugCApreA9QeZuqh8swQMa"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1iKqZc4CfrPwjcD7kx8MZtFtfL6YysKnPv"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1iKqZc4CfrPwjcD7kx8MZtFtfL6YysKnPv"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz2DWwmauKy4Dkak6X8ePY4Z91EsM36miYxB"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz2DWwmauKy4Dkak6X8ePY4Z91EsM36miYxB"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1aCb6QKL7K4JxRTUTYjAER8qZH4zKAbigq"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1aCb6QKL7K4JxRTUTYjAER8qZH4zKAbigq"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1f4zxfTtjTKkrzzXNtQU73zGBcasW3xEBS"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1f4zxfTtjTKkrzzXNtQU73zGBcasW3xEBS"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1TG2bxjvmRMQakt8KVQ5S6CyLzMuSHx2qF"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1TG2bxjvmRMQakt8KVQ5S6CyLzMuSHx2qF"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1YibTr9kUchgG6wJjMwRar2NUfawxH48FT"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1YibTr9kUchgG6wJjMwRar2NUfawxH48FT"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1YtK9qUJffq2SqjhaJEGbufVHEBWDLPrSV"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1YtK9qUJffq2SqjhaJEGbufVHEBWDLPrSV"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz2DTjmLqMXSwhkpCFnV977vyhtYWbsP6nwo"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz2DTjmLqMXSwhkpCFnV977vyhtYWbsP6nwo"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1eaCp9cc4FtTs3Wm6UgZZnEtZ9b1Peqewi"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1eaCp9cc4FtTs3Wm6UgZZnEtZ9b1Peqewi"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1MZ7ypArEoYocyyfscGaneNKM7bADhnQC4"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1MZ7ypArEoYocyyfscGaneNKM7bADhnQC4"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1Qzov1LwEvSfHonRHr2EwQGto7ExrH43or"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1Qzov1LwEvSfHonRHr2EwQGto7ExrH43or"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz28e4LURT9Kv6TXysL6szWYkKKqRmA1wDGq"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz28e4LURT9Kv6TXysL6szWYkKKqRmA1wDGq"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1UC8X57hKRmTbgJWsoos9njMy2SW3bKAAh"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1UC8X57hKRmTbgJWsoos9njMy2SW3bKAAh"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1XX6ykvpfPGAkzeDaFPvvPqmRQoTwCvP4S"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1XX6ykvpfPGAkzeDaFPvvPqmRQoTwCvP4S"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1P91PPAfukVdpgj153WqoLymrWDYCoyLRA"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1P91PPAfukVdpgj153WqoLymrWDYCoyLRA"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz2JvpjaCGsynphe78Xbyu2kHUU2Ym8rxi8v"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz2JvpjaCGsynphe78Xbyu2kHUU2Ym8rxi8v"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1YDHyB7aBCvj5TCYiRQkjeVHfgtEYUDAd2"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1YDHyB7aBCvj5TCYiRQkjeVHfgtEYUDAd2"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1VWKeeSqxdLsTfAyCSrwFVZnfgw4ZwaHWV"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1VWKeeSqxdLsTfAyCSrwFVZnfgw4ZwaHWV"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1Xfkbwwa9ewNvYnesSXEYWGgS5gXVSQQVj"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1Xfkbwwa9ewNvYnesSXEYWGgS5gXVSQQVj"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1Mv79qqYx2QX1NXYawrSrya13T7QoxCba9"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1Mv79qqYx2QX1NXYawrSrya13T7QoxCba9"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz2H4arakxXuzfo8NBVqTNgL5NBZ9b497Pfv"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz2H4arakxXuzfo8NBVqTNgL5NBZ9b497Pfv"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.MustParseAddress("tz1WxZvegEFTFSA8mqHWsBq97S15kS34JXqu"),
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1WxZvegEFTFSA8mqHWsBq97S15kS34JXqu"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
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
				Baker:            mavryk.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
				Delegator:        mavryk.Address{}, // Represents an empty delegator address
				Cycle:            1016,
				Recipient:        mavryk.MustParseAddress("tz1UGkfyrT9yBt6U5PV7Qeui3pt3a8jffoWv"),
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
		if payout.Recipient == mavryk.MustParseAddress("tz1X7U9XxVz6NDxL4DSZhijME61PW45bYUJE") {
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

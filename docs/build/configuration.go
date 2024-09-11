package main

import (
	"encoding/json"
	"os"

	"github.com/hjson/hjson-go/v4"
	"github.com/mavryk-network/mavpay/common"
	mavpay_configuration "github.com/mavryk-network/mavpay/configuration/v"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mvgo/mavryk"
)

func GenerateDefaultHJson() {
	config := mavpay_configuration.GetDefaultV0()

	sample, _ := hjson.MarshalWithOptions(config, hjson.EncoderOptions{
		Eol:                   "\n",
		BracesSameLine:        true,
		EmitRootBraces:        true,
		QuoteAlways:           false,
		QuoteAmbiguousStrings: true,
		IndentBy:              "  ",
		BaseIndentation:       "",
		Comments:              false,
	})
	_ = os.WriteFile("docs/configuration/config.default.hjson", sample, 0644)
}

func genrateSample() *mavpay_configuration.ConfigurationV0 {
	logExtensionConfiguration := json.RawMessage(`{"LOG_FILE": "path/to/my/extension.log"}`)
	feeExtensionConfiguration := json.RawMessage(`{"FEE": 0, "TOKEN": "1", "CONTRACT": "KT1Hkg6qgV3VykjgUXKbWcU3h6oJ1qVxUxZV"}`)

	fee := 0.0
	donate := 0.025
	donateFees := 0.05
	donateBonds := 0.03
	gasLimitBuffer := int64(200)
	deserializationGasBuffer := int64(5)
	feeBuffer := int64(10)
	ktFeeBuffer := int64(50)
	bellowMinimumBalanceRewardDestination := enums.REWARD_DESTINATION_EVERYONE
	maximumBalance := float64(1000.0)
	minimumDelayBlocks := int64(10)
	maximumDelayBlocks := int64(250)

	return &mavpay_configuration.ConfigurationV0{
		Version:  0,
		BakerPKH: mavryk.InvalidAddress,
		Delegators: mavpay_configuration.DelegatorsConfigurationV0{
			Requirements: mavpay_configuration.DelegatorRequirementsV0{
				MinimumBalance:                        float64(0.5),
				BellowMinimumBalanceRewardDestination: &bellowMinimumBalanceRewardDestination,
			},
			Overrides: map[string]mavpay_configuration.DelegatorOverrideV0{
				"mv1HCXRedE7zVSwmSqxDe3XZcMPLeF7xYqP3": {
					Recipient:      mavryk.InvalidAddress,
					Fee:            &fee,
					MinimumBalance: 2.5,
				},
				"mv1Qe2hoRHRHYxYCHzD8vUX2We8uEJrEdWAb": {
					MaximumBalance: &maximumBalance,
				},
			},
			FeeOverrides: map[string][]mavryk.Address{
				"1":  {mavryk.ZeroAddress, mavryk.BurnAddress},
				".5": {mavryk.InvalidAddress},
			},
			Ignore:    []mavryk.Address{mavryk.ZeroAddress, mavryk.BurnAddress},
			Prefilter: []mavryk.Address{mavryk.MustParseAddress("mv1HCXRedE7zVSwmSqxDe3XZcMPLeF7xYqP3"), mavryk.MustParseAddress("mv1Qe2hoRHRHYxYCHzD8vUX2We8uEJrEdWAb")},
		},
		Network: mavpay_configuration.MavrykNetworkConfigurationV0{
			RpcUrl:                 constants.DEFAULT_RPC_URL,
			MvktUrl:                constants.DEFAULT_MVKT_URL,
			ProtocolRewardsUrl:     constants.DEFAULT_PROTOCOL_REWARDS_URL,
			Explorer:               "https://tzstats.com/",
			DoNotPaySmartContracts: true,
		},
		Overdelegation: mavpay_configuration.OverdelegationConfigurationV0{
			IsProtectionEnabled: true,
		},
		PayoutConfiguration: mavpay_configuration.PayoutConfigurationV0{
			WalletMode:                 enums.WALLET_MODE_LOCAL_PRIVATE_KEY,
			PayoutMode:                 enums.PAYOUT_MODE_IDEAL,
			BalanceCheckMode:           enums.PROTOCOL_BALANCE_CHECK_MODE,
			Fee:                        .075,
			IsPayingTxFee:              true,
			IsPayingAllocationTxFee:    true,
			MinimumAmount:              10.5,
			TxGasLimitBuffer:           &gasLimitBuffer,
			TxDeserializationGasBuffer: &deserializationGasBuffer,
			TxFeeBuffer:                &feeBuffer,
			KtTxFeeBuffer:              &ktFeeBuffer,
			MinimumDelayBlocks:         &minimumDelayBlocks,
			MaximumDelayBlocks:         &maximumDelayBlocks,
		},
		NotificationConfigurations: []json.RawMessage{
			json.RawMessage(`{
				"type":             "discord",
				"webhook_url":      "https://my-discord-webhook.com/",
				"message_template": "my awesome message",
			}`),
			json.RawMessage(`{
				"type":             "discord",
				"webhook_url":      "https://my-admin-discord-webhook.com/",
				"message_template": "my awesome message",
				"admin":            true,
			}`),
			json.RawMessage(`{
				"type":             "discord",
				"webhook_id":       "webhook id",
				"webhook_token":    "webhook token",
				"message_template": "my awesome message",
			}`),
			json.RawMessage(`{
				"type":                "twitter",
				"access_token":        "your access token",
				"access_token_secret": "your access token secret",
				"consumer_key":        "your consumer key",
				"consumer_secret":     "your consumer secret",
				"message_template":    "my awesome message",
			}`),
			json.RawMessage(`{
				"type":             "telegram",
				"api_token":        "your api token",
				"receivers":        ["list of chat numbers without quotes", -1234567890],
				"message_template": "my awesome message",
			}`),
			json.RawMessage(`{
				"type":             "email",
				"sender":           "my@email.is",
				"smtp_server":      "smtp.gmail.com:443",
				"smtp_identity":    "",
				"smtp_username":    "my@email.is",
				"smtp_password":    "password123",
				"recipients":       ["my-follower1@email.is", "my-follower2@email.is"],
				"message_template": "my awesome message",
			}`),
			json.RawMessage(`{
				"type": "external",
				"path": "path to external notificator binary",
				"args": ["--kind", "<kind>", "<data>"],
			}`),
		},
		IncomeRecipients: mavpay_configuration.IncomeRecipientsV0{
			Bonds: map[string]float64{
				"mv1HCXRedE7zVSwmSqxDe3XZcMPLeF7xYqP3": 0.455,
				"tz1X7U9XxVz6NDxL4DSZhijME61PW45bYUJE": 0.545,
			},
			Fees: map[string]float64{
				"mv1HCXRedE7zVSwmSqxDe3XZcMPLeF7xYqP3": 0.455,
				"tz1X7U9XxVz6NDxL4DSZhijME61PW45bYUJE": 0.545,
			},
			Donate:      &donate,
			DonateFees:  &donateFees,
			DonateBonds: &donateBonds,
			Donations: map[string]float64{
				"mv1HCXRedE7zVSwmSqxDe3XZcMPLeF7xYqP3": 0.10,
				"mv1V4h45W3p4e1sjSBvRkK2uYbvkTnSuHg8g": 0.90,
			},
		},
		Extensions: []mavpay_configuration.ExtensionConfigurationV0{
			common.ExtensionDefinition{
				Name:    "log-extension",
				Command: "python3",
				Args:    []string{"/path/to/my/extension.py"},
				Kind:    enums.EXTENSION_STDIO_RPC,
				Hooks: []common.ExtensionHook{{
					Id:   enums.EXTENSION_HOOK_ALL,
					Mode: enums.EXTENSION_HOOK_MODE_READ_ONLY,
				}},
				Configuration: &logExtensionConfiguration,
			},
			common.ExtensionDefinition{
				Name:    "fee-extension",
				Command: `/path/to/my/extension.bin`,
				Args:    []string{"--config", "/path/to/my/extension.config"},
				Kind:    enums.EXTENSION_STDIO_RPC,
				Hooks: []common.ExtensionHook{{
					Id:   enums.EXTENSION_HOOK_AFTER_CANDIDATES_GENERATED,
					Mode: enums.EXTENSION_HOOK_MODE_READ_WRITE,
				}},
				Configuration: &feeExtensionConfiguration,
			},
		},
		DisableAnalytics: true,
	}
}

func GenerateSampleHJson() {
	config := genrateSample()

	sample, _ := hjson.Marshal(config)
	_ = os.WriteFile("docs/configuration/config.sample.hjson", sample, 0644)
}

func genrateStarter() *mavpay_configuration.ConfigurationV0 {
	return &mavpay_configuration.ConfigurationV0{
		Version:  0,
		BakerPKH: mavryk.InvalidAddress,
		Delegators: mavpay_configuration.DelegatorsConfigurationV0{
			Requirements: mavpay_configuration.DelegatorRequirementsV0{
				MinimumBalance: float64(10),
			},
		},
		Overdelegation: mavpay_configuration.OverdelegationConfigurationV0{
			IsProtectionEnabled: true,
		},
		PayoutConfiguration: mavpay_configuration.PayoutConfigurationV0{
			WalletMode:       enums.WALLET_MODE_LOCAL_PRIVATE_KEY,
			BalanceCheckMode: enums.PROTOCOL_BALANCE_CHECK_MODE,
			PayoutMode:       enums.PAYOUT_MODE_ACTUAL,
			Fee:              .10,
			MinimumAmount:    0.01,
		},
	}
}

func GenerateStarterHJson() {
	config := genrateStarter()

	defaultMarshaled, _ := hjson.Marshal(config)
	var node hjson.Node
	_ = hjson.Unmarshal(defaultMarshaled, &node)

	node.Cm.InsideFirst = "\n#=====================================================================================================\n" +
		"# This is mavpay starter configuration template. Please refer to https://mavpay.mavryk.org/mavpay/\n" +
		"# - for default configuration (list of default values) see https://mavpay.mavryk.org/mavpay/configuration/examples/default/.\n" +
		"# - for sample of all available fields see https://mavpay.mavryk.org/mavpay/configuration/examples/sample/.\n" +
		"#=====================================================================================================\n"

	node.DeleteKey("network")
	node.DeleteKey("income_recipients")
	node.NKC("baker").Value = "your-baker-address"
	node.NKC("payouts").DeleteKey("wallet_mode")
	node.NKC("payouts").DeleteKey("payout_mode")

	defaultMarshaled, _ = hjson.Marshal(node)
	_ = os.WriteFile("docs/configuration/config.starter.hjson", defaultMarshaled, 0644)
}

package seed

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/hjson/hjson-go/v4"
	trd_seed "github.com/mavryk-network/mavpay/configuration/seed/trd"
	mavpay_configuration "github.com/mavryk-network/mavpay/configuration/v"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mavpay/notifications"
	"github.com/mavryk-network/mvgo/mavryk"
	"gopkg.in/yaml.v3"
)

// %CYCLE%, %NDELEGATORS%, %TREWARDS%
func trdAliasing(configuration []byte) []byte {
	config := string(configuration)
	config = strings.ReplaceAll(config, "%TREWARDS%", "<DistributedRewards>")
	config = strings.ReplaceAll(config, "%CYCLE%", "<Cycle>")
	config = strings.ReplaceAll(config, "%NDELEGATORS%", "<Delegators>")
	return []byte(config)
}

func MigrateTrdv1ToTPv0(sourceBytes []byte) ([]byte, error) {
	slog.Debug("migrating trd configuration to mavpay")
	configuration := trd_seed.GetDefault()
	err := yaml.Unmarshal(trdAliasing(sourceBytes), &configuration)
	if err != nil {
		return []byte{}, err
	}

	address, err := mavryk.ParseAddress(configuration.BakingAddress)
	if err != nil {
		return []byte{}, err
	}

	feeRecipients := make(map[string]float64, len(configuration.FoundersMap))
	if len(configuration.FoundersMap) > 0 {
		for recipient, share := range configuration.FoundersMap {
			feeRecipients[recipient] = share
		}
	}

	bondRecipients := make(map[string]float64, len(configuration.OwnersMap))
	if len(configuration.OwnersMap) > 0 {
		for recipient, share := range configuration.OwnersMap {
			bondRecipients[recipient] = share
		}
	}

	delegatorOverrides := make(map[string]mavpay_configuration.DelegatorOverrideV0, len(configuration.SpecialsMap)+len(configuration.SupportersSet))
	if len(configuration.SpecialsMap) > 0 {
		for recipient, share := range configuration.SpecialsMap {
			if addr, err := mavryk.ParseAddress(recipient); err == nil {
				delegatorOverrides[recipient] = mavpay_configuration.DelegatorOverrideV0{
					Recipient:      addr,
					Fee:            &share,
					MinimumBalance: 0,
				}
			}
		}
	}

	if len(configuration.SupportersSet) > 0 {
		fee := 0.0
		for recipient := range configuration.SupportersSet {
			if _, err := mavryk.ParseAddress(recipient); err == nil {
				if v, ok := delegatorOverrides[recipient]; ok {
					if v.Fee == nil {
						v.Fee = &fee
					}
					continue
				}
				delegatorOverrides[recipient] = mavpay_configuration.DelegatorOverrideV0{
					Fee: &fee,
				}
			}
		}
	}

	delegatorBellowMinimumBalanceRewardDestination := enums.REWARD_DESTINATION_NONE
	feeOverrides := make(map[string][]mavryk.Address, 0)
	ignores := make([]mavryk.Address, 0)
	if len(configuration.RulesMap) > 0 {
		for k, v := range configuration.RulesMap {
			if k == "mindelegation" {
				if v == "TOE" {
					delegatorBellowMinimumBalanceRewardDestination = enums.REWARD_DESTINATION_EVERYONE
				}
				continue
			}

			switch v {
			// if TOE -> ignore
			case "TOE":
				if addr, err := mavryk.ParseAddress(k); err == nil {
					ignores = append(ignores, addr)
				} else {
					slog.Warn("failed to parse address, please adjust configuration manually", "address", k)
				}
				// if TOB -> fee 1
			case "TOB":
				fallthrough
				// if TOF -> fee 1
			case "TOF":
				if addr, err := mavryk.ParseAddress(k); err == nil {
					if _, ok := feeOverrides["1"]; !ok {
						feeOverrides["1"] = make([]mavryk.Address, 0)
					}
					feeOverrides["1"] = append(feeOverrides["1"], addr)
				} else {
					slog.Warn("failed to parse address, please adjust configuration manually", "address", k)
				}
			case "Dexter":
				slog.Warn("we do not support dexter right now, please check your configuration file and migrate it manually")
			default:
				targetAddr, err := mavryk.ParseAddress(v)
				if err == nil { // if address redirect
					if sourceAddr, err := mavryk.ParseAddress(k); err == nil {
						if v, ok := delegatorOverrides[sourceAddr.String()]; ok {
							if v.Recipient != mavryk.ZeroAddress {
								slog.Warn("address already has a recipient, please adjust configuration manually", "address", k, "recipient", v.Recipient)
							} else {
								v.Recipient = targetAddr
							}
						} else {
							delegatorOverrides[sourceAddr.String()] = mavpay_configuration.DelegatorOverrideV0{
								Recipient: targetAddr,
							}
						}
					} else {
						slog.Warn("failed to parse address, please adjust configuration manually", "address", k)
					}
				} else {
					slog.Warn("failed to parse address - unknown rules map value, please adjust configuration manually", "address", v)
				}
			}
		}
	}

	notificationConfigurations := make([]json.RawMessage, 0)
	if configuration.Plugins != nil {
		for t, plugin := range configuration.Plugins {
			switch t {
			case "email":
				oldConfig := trd_seed.EmailPluginConfigurationV1{}
				err := plugin.Decode(&oldConfig)
				if err != nil {
					slog.Warn("we are not able to migrate email plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					continue
				}
				configuration := notifications.EmailNotificatorConfiguration{
					Type:       "email",
					Sender:     oldConfig.SMTPSender,
					SmtpUser:   oldConfig.SMTPUser,
					SmtpPass:   oldConfig.SMTPPass,
					Recipients: oldConfig.SMTPRecipients,
				}
				if oldConfig.SMTPPort == 0 {
					configuration.SmtpServer = oldConfig.SMTPHost
				} else {
					configuration.SmtpServer = fmt.Sprintf("%s:%d", oldConfig.SMTPHost, oldConfig.SMTPPort)
				}
				result, err := json.Marshal(configuration)
				if err != nil {
					slog.Warn("we are not able to migrate twitter plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					continue
				}
				notificationConfigurations = append(notificationConfigurations, result)
			case "webhook":
				slog.Warn("POST request by trd webhook plugin differs from mavpay, you may have to adjust your webhook logic")
				var configuration trd_seed.WebhookPluginConfigurationV1
				err := plugin.Decode(&configuration)
				if err != nil {
					// log and skip
					slog.Warn("we are not able to migrate webhook plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					continue
				}
				result, err := json.Marshal(map[string]any{
					"type":  "webhook",
					"url":   configuration.Endpoint,
					"token": configuration.Token,
				})
				if err != nil {
					slog.Warn("we are not able to migrate webhook plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					continue
				}
				notificationConfigurations = append(notificationConfigurations, result)
			case "telegram":
				var configuration trd_seed.TelegramPluginConfigurationV1
				err := plugin.Decode(&configuration)
				if err != nil {
					// log and skip
					slog.Warn("we are not able to migrate telegram plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					continue
				}
				if len(configuration.AdminChatsIds) > 0 {
					config, err := json.Marshal(map[string]any{
						"type":             "telegram",
						"admin":            true,
						"recipients":       configuration.AdminChatsIds,
						"api_token":        configuration.BotApiKey,
						"message_template": configuration.TelegramText,
					})
					if err == nil {
						notificationConfigurations = append(notificationConfigurations, config)
					} else {
						slog.Warn("we are not able to migrate telegram plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					}
				}
				if len(configuration.ChatIds) > 0 {
					configuration.PayoutChatsIds = append(configuration.PayoutChatsIds, configuration.ChatIds...)
				}
				if len(configuration.PayoutChatsIds) > 0 {
					config, err := json.Marshal(map[string]any{
						"type":             "telegram",
						"admin":            false,
						"recipients":       configuration.PayoutChatsIds,
						"api_token":        configuration.BotApiKey,
						"message_template": configuration.TelegramText,
					})
					if err == nil {
						notificationConfigurations = append(notificationConfigurations, config)
					} else {
						slog.Warn("we are not able to migrate telegram plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					}
				}
			case "twitter":
				var configuration trd_seed.TwitterPluginConfigurationV1
				err := plugin.Decode(&configuration)
				if err != nil {
					// log and skip
					slog.Warn("we are not able to migrate twitter plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					continue
				}
				configuration.Type = "twitter"
				result, err := json.Marshal(configuration)
				if err != nil {
					slog.Warn("we are not able to migrate twitter plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					continue
				}
				notificationConfigurations = append(notificationConfigurations, result)
			case "discord":
				var configuration trd_seed.DiscordPluginConfigurationV1
				err := plugin.Decode(&configuration)
				if err != nil {
					// log and skip
					slog.Warn("we are not able to migrate discord plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					continue
				}
				configuration.Type = "discord"
				result, err := json.Marshal(configuration)
				if err != nil {
					slog.Warn("we are not able to migrate discord plugin configuration right now, please check your configuration file and migrate it manually", "error", err.Error())
					continue
				}
				notificationConfigurations = append(notificationConfigurations, result)
			}
		}
	}

	donate := 0.05

	migrated := mavpay_configuration.ConfigurationV0{
		Version:  0,
		BakerPKH: address,
		IncomeRecipients: mavpay_configuration.IncomeRecipientsV0{
			Bonds:  bondRecipients,
			Fees:   feeRecipients,
			Donate: &donate,
		},
		Delegators: mavpay_configuration.DelegatorsConfigurationV0{
			Requirements: mavpay_configuration.DelegatorRequirementsV0{
				MinimumBalance:                        configuration.MinDelegation,
				BellowMinimumBalanceRewardDestination: &delegatorBellowMinimumBalanceRewardDestination,
			},
			Overrides:    delegatorOverrides,
			FeeOverrides: feeOverrides,
			Ignore:       ignores,
		},
		Network: mavpay_configuration.MavrykNetworkConfigurationV0{
			RpcUrl:                 constants.DEFAULT_RPC_URL,
			MvktUrl:                constants.DEFAULT_MVKT_URL,
			ProtocolRewardsUrl:     constants.DEFAULT_PROTOCOL_REWARDS_URL,
			DoNotPaySmartContracts: false,
		},
		Overdelegation: mavpay_configuration.OverdelegationConfigurationV0{
			IsProtectionEnabled: true,
		},
		PayoutConfiguration: mavpay_configuration.PayoutConfigurationV0{
			Fee:                     configuration.ServiceFee / 100,
			IsPayingTxFee:           !configuration.DelPaysXferFee,
			IsPayingAllocationTxFee: !configuration.DelPaysRaFee,
			IgnoreEmptyAccounts:     !configuration.ReactivateZero,
			WalletMode:              enums.WALLET_MODE_LOCAL_PRIVATE_KEY,
			PayoutMode:              enums.EPayoutMode(configuration.RewardsType),
			BalanceCheckMode:        enums.PROTOCOL_BALANCE_CHECK_MODE,
			MinimumAmount:           configuration.MinPayment,
		},
		NotificationConfigurations: notificationConfigurations,
	}

	migratedBytes, err := hjson.MarshalWithOptions(migrated, getSerializeHjsonOptions())
	if err != nil {
		return []byte{}, err
	}
	slog.Debug("migrated bc configuration successfully")
	return migratedBytes, nil
}

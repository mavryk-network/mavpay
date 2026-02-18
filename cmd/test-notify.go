package cmd

import (
	"log/slog"
	"time"

	"github.com/mavryk-network/mavpay/notifications"
	"github.com/spf13/cobra"
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/notifications"
	"github.com/mavryk-network/mvgo/mavryk"
)

var notificationTestCmd = &cobra.Command{
	Use:   "test-notify",
	Short: "notification test",
	Long:  "sends test notification",
	Run: func(cmd *cobra.Command, args []string) {
		config, _, _, _ := assertRunWithResult(loadConfigurationEnginesExtensions, EXIT_CONFIGURATION_LOAD_FAILURE).Unwrap()
		notificator, _ := cmd.Flags().GetString(NOTIFICATOR_FLAG)
		wantsAdmin, _ := cmd.Flags().GetBool("admin")
		wantsPayoutSummary, _ := cmd.Flags().GetBool("payout-summary")
		for _, notificatorConfiguration := range config.NotificationConfigurations {
			if notificator != "" && string(notificatorConfiguration.Type) != notificator {
				continue
			}

			slog.Info("sending notification", "notificator", notificatorConfiguration.Type)
			notificator, err := notifications.LoadNotificatior(notificatorConfiguration.Type, notificatorConfiguration.Configuration)
			if err != nil {
				slog.Warn("failed to send notification", "error", err.Error())
				continue
			}

			switch {
			case wantsPayoutSummary:
				err = notificator.PayoutSummaryNotify(&common.PayoutSummary{
					Cycles: []int64{123, 124, 125},
					CyclePayoutSummary: common.CyclePayoutSummary{
						Delegators:               60,
						PaidDelegators:           27,
						OwnStakedBalance:         mavryk.NewZ(11041500351),
						OwnDelegatedBalance:      mavryk.NewZ(3659574532),
						ExternalStakedBalance:    mavryk.NewZ(23868693294),
						ExternalDelegatedBalance: mavryk.NewZ(78248642581),
						EarnedBlockFees:          mavryk.NewZ(49075),
						EarnedRewards:            mavryk.NewZ(30287160),
						DistributedRewards:       mavryk.NewZ(25538039),
						BondIncome:               mavryk.NewZ(1339668),
						FeeIncome:                mavryk.NewZ(3356503),
						IncomeTotal:              mavryk.NewZ(4696171),
						TxFeesPaid:               mavryk.NewZ(68564),
						DonatedBonds:             mavryk.NewZ(13531),
						DonatedFees:              mavryk.NewZ(33904),
						DonatedTotal:             mavryk.NewZ(47435),
						Timestamp:                time.Now(),
					},
				}, map[string]string{})
				if err != nil {
					slog.Warn("failed to send notification", "error", err.Error())
					continue
				}
			case wantsAdmin:
				if !notificatorConfiguration.IsAdmin {
					continue
				}
				err = notificator.AdminNotify("test admin notification")
				if err != nil {
					slog.Warn("failed to send notification", "error", err.Error())
					continue
				}
			default:
				err = notificator.TestNotify()
				if err != nil {
					slog.Warn("failed to send notification", "error", err.Error())
					continue
				}
			}

		}
	},
}

func init() {
	notificationTestCmd.Flags().String(NOTIFICATOR_FLAG, "", "Notify through specific notificator")
	notificationTestCmd.Flags().Bool("admin", false, "Notify through admin notificators")
	notificationTestCmd.Flags().Bool("payout-summary", false, "Send payout-summary notification with dummy data")

	RootCmd.AddCommand(notificationTestCmd)
}

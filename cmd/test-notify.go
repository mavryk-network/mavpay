package cmd

import (
	"log/slog"

	"github.com/mavryk-network/mavpay/notifications"
	"github.com/spf13/cobra"
)

var notificationTestCmd = &cobra.Command{
	Use:   "test-notify",
	Short: "notification test",
	Long:  "sends test notification",
	Run: func(cmd *cobra.Command, args []string) {
		config, _, _, _ := assertRunWithResult(loadConfigurationEnginesExtensions, EXIT_CONFIGURATION_LOAD_FAILURE).Unwrap()
		notificator, _ := cmd.Flags().GetString(NOTIFICATOR_FLAG)
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

			err = notificator.TestNotify()
			if err != nil {
				slog.Warn("failed to send notification", "error", err.Error())
				continue
			}
		}
	},
}

func init() {
	notificationTestCmd.Flags().String(NOTIFICATOR_FLAG, "", "Notify through specific notificator")

	RootCmd.AddCommand(notificationTestCmd)
}

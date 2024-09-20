package cmd

import (
	"encoding/json"
	"log/slog"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mavpay/extension"
	"github.com/mavryk-network/mavpay/notifications"
)

func collectAdditionalData(_ *common.CyclePayoutSummary) map[string]string {
	data := make(map[string]json.RawMessage)

	err := extension.ExecuteHook(enums.EXTENSION_HOOK_COLLECT_ADDITIONAL_NOTIFICATION_DATA, "0.1", &data)
	if err != nil {
		slog.Warn("failed to execute hook", "error", err.Error())
	}
	result := make(map[string]string)
	for key, value := range data {
		result[key] = string(value)
	}

	return result
}

func notifyPayoutsProcessed(configuration *configuration.RuntimeConfiguration, summary *common.CyclePayoutSummary, filter string) {
	for _, notificatorConfiguration := range configuration.NotificationConfigurations {
		if filter != "" && string(notificatorConfiguration.Type) != filter {
			continue
		}

		if notificatorConfiguration.IsAdmin {
			continue
		}

		slog.Info("sending notification", "notificator", notificatorConfiguration.Type)
		notificator, err := notifications.LoadNotificatior(notificatorConfiguration.Type, notificatorConfiguration.Configuration)
		if err != nil {
			slog.Warn("failed to send notification", "error", err.Error())
			continue
		}

		additionalData := collectAdditionalData(summary)
		err = notificator.PayoutSummaryNotify(summary, additionalData)
		if err != nil {
			slog.Warn("failed to send notification", "error", err.Error())
			continue
		}
	}
	slog.Info("notifications sent")
}
func notifyPayoutsProcessedThroughAllNotificators(configuration *configuration.RuntimeConfiguration, summary *common.CyclePayoutSummary) {
	notifyPayoutsProcessed(configuration, summary, "")
}

func notifyAdmin(configuration *configuration.RuntimeConfiguration, msg string) {
	for _, notificatorConfiguration := range configuration.NotificationConfigurations {
		if !notificatorConfiguration.IsAdmin {
			continue
		}

		slog.Debug("sending admin notification", "notificator", notificatorConfiguration.Type)
		notificator, err := notifications.LoadNotificatior(notificatorConfiguration.Type, notificatorConfiguration.Configuration)
		if err != nil {
			slog.Warn("failed to send notification", "error", err.Error())
			continue
		}

		err = notificator.AdminNotify(msg)
		if err != nil {
			slog.Warn("failed to send notification", "error", err.Error())
			continue
		}
	}
	slog.Debug("admin notifications sent")
}

func notifyAdminFactory(configuration *configuration.RuntimeConfiguration) func(string) {
	return func(msg string) {
		notifyAdmin(configuration, msg)
	}
}

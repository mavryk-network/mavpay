package notifications

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mvgo/mavryk"
)

type NotificationKind string

const (
	PAYOUT_SUMMARY_NOTIFICATION NotificationKind = "payout_summary"
	ADMIN_NOTIFICATION          NotificationKind = "admin"
	TEST_NOTIFICATION           NotificationKind = "test"
	TEXT_NOTIFICATION           NotificationKind = "text"
)

type NotificatorKind string

const (
	TELEGRAM_NOTIFICATOR NotificatorKind = "telegram"
	TWITTER_NOTIFICATOR  NotificatorKind = "twitter"
	DISCORD_NOTIFICATOR  NotificatorKind = "discord"
	EMAIL_NOTIFICATOR    NotificatorKind = "email"
	EXTERNAL_NOTIFICATOR NotificatorKind = "external"
	WEBHOOK_NOTIFICATOR  NotificatorKind = "webhook"
)

func PopulateMessageTemplate(messageTempalte string, summary *common.CyclePayoutSummary, additionalData map[string]string) string {
	v := reflect.ValueOf(*summary)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		val := fmt.Sprintf("%v", v.Field(i).Interface())
		if typeOfS.Field(i).Type.Name() == "Z" && strings.Contains(typeOfS.Field(i).Type.PkgPath(), "mvgo/mavryk") {
			val = fmt.Sprintf("%v", common.MumavToMavS(v.Field(i).Interface().(mavryk.Z).Int64()))
		}
		messageTempalte = strings.ReplaceAll(messageTempalte, fmt.Sprintf("<%s>", typeOfS.Field(i).Name), val)
	}

	for k, v := range additionalData {
		messageTempalte = strings.ReplaceAll(messageTempalte, fmt.Sprintf("<%s>", k), v)
	}

	return messageTempalte
}

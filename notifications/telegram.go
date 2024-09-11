package notifications

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/nikoksr/notify/service/telegram"
	"golang.org/x/exp/slog"
)

type telegramNotificatorConfiguration struct {
	Type            string  `json:"type"`
	Token           string  `json:"api_token"`
	Receivers       []int64 `json:"receivers"`
	MessageTemplate string  `json:"message_template"`
}

type TelegramNotificator struct {
	session         *telegram.Telegram
	messageTemplate string
}

const (
	DEFAULT_TELEGRAM_MESSAGE_TEMPLATE = "A total of <DistributedRewards> was distributed for cycle <Cycle> to <Delegators> delegators and donated <DonatedTotal> using #tezpay on the #tezos blockchain."
)

func InitTelegramNotificator(configurationBytes []byte) (*TelegramNotificator, error) {
	configuration := telegramNotificatorConfiguration{}
	err := json.Unmarshal(configurationBytes, &configuration)
	if err != nil {
		return nil, err
	}
	msgTemplate := configuration.MessageTemplate
	if msgTemplate == "" {
		msgTemplate = DEFAULT_TELEGRAM_MESSAGE_TEMPLATE
	}

	session, err := telegram.New(configuration.Token)
	if err != nil {
		return nil, err
	}
	for _, v := range configuration.Receivers {
		session.AddReceivers(v)
	}

	slog.Debug("telegram notificator initialized")

	return &TelegramNotificator{
		session:         session,
		messageTemplate: msgTemplate,
	}, nil
}

func ValidateTelegramConfiguration(configurationBytes []byte) error {
	configuration := telegramNotificatorConfiguration{}
	err := json.Unmarshal(configurationBytes, &configuration)
	if err != nil {
		return err
	}
	if configuration.Token == "" {
		return errors.Join(constants.ErrInvalidNotificatorConfiguration, errors.New("invalid telegram api token"))
	}
	if len(configuration.Receivers) == 0 {
		return errors.Join(constants.ErrInvalidNotificatorConfiguration, errors.New("no telegram receivers specified"))
	}
	return nil
}

func (tn *TelegramNotificator) PayoutSummaryNotify(summary *common.CyclePayoutSummary, additionalData map[string]string) error {
	return tn.session.Send(context.Background(), fmt.Sprintf("Report of cycle #%d", summary.Cycle), PopulateMessageTemplate(tn.messageTemplate, summary, additionalData))
}

func (tn *TelegramNotificator) AdminNotify(msg string) error {
	return tn.session.Send(context.Background(), string(ADMIN_NOTIFICATION), msg)
}

func (tn *TelegramNotificator) TestNotify() error {
	return tn.session.Send(context.Background(), "test notification", "telegram test")
}

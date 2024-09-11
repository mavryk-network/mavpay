package notifications

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/nikoksr/notify/service/mail"
)

type EmailNotificatorConfiguration struct {
	Type            string   `json:"type"`
	Sender          string   `json:"sender"`
	SmtpServer      string   `json:"smtp_server"`
	SmtpIdentity    string   `json:"smtp_identity"`
	SmtpUser        string   `json:"smtp_username"`
	SmtpPass        string   `json:"smtp_password"`
	Recipients      []string `json:"recipients"`
	MessageTemplate string   `json:"message_template"`
}

type EmailNotificator struct {
	session         *mail.Mail
	messageTemplate string
}

const (
	DEFAULT_EMAIL_MESSAGE_TEMPLATE = "A total of <DistributedRewards> was distributed for cycle <Cycle> to <Delegators> delegators and donated <DonatedTotal> using #tezpay on the #tezos blockchain."
)

func InitEmailNotificator(configurationBytes []byte) (*EmailNotificator, error) {
	configuration := EmailNotificatorConfiguration{}
	err := json.Unmarshal(configurationBytes, &configuration)
	if err != nil {
		return nil, err
	}
	msgTemplate := configuration.MessageTemplate
	if msgTemplate == "" {
		msgTemplate = DEFAULT_EMAIL_MESSAGE_TEMPLATE
	}

	session := mail.New(configuration.Sender, configuration.SmtpServer)
	session.AddReceivers(configuration.Recipients...)

	smtpHost, _, err := net.SplitHostPort(configuration.SmtpServer)
	if err != nil {
		return nil, err
	}

	if configuration.SmtpUser != "" && configuration.SmtpPass != "" {
		session.AuthenticateSMTP(configuration.SmtpIdentity, configuration.SmtpUser, configuration.SmtpPass, smtpHost)
	}

	slog.Debug("email notificator initialized")

	return &EmailNotificator{
		session:         session,
		messageTemplate: msgTemplate,
	}, nil
}

func ValidateEmailConfiguration(configurationBytes []byte) error {
	configuration := EmailNotificatorConfiguration{}
	err := json.Unmarshal(configurationBytes, &configuration)
	if err != nil {
		return err
	}
	if configuration.Sender == "" {
		return errors.Join(constants.ErrInvalidNotificatorConfiguration, errors.New("invalid email sender"))
	}
	if len(configuration.Recipients) == 0 {
		return errors.Join(constants.ErrInvalidNotificatorConfiguration, errors.New("no email recipients specified"))
	}
	return nil
}

func (en *EmailNotificator) PayoutSummaryNotify(summary *common.CyclePayoutSummary, additionalData map[string]string) error {
	return en.session.Send(context.Background(), fmt.Sprintf("Report of cycle #%d", summary.Cycle), PopulateMessageTemplate(en.messageTemplate, summary, additionalData))
}

func (en *EmailNotificator) AdminNotify(msg string) error {
	return en.session.Send(context.Background(), string(ADMIN_NOTIFICATION), msg)
}

func (en *EmailNotificator) TestNotify() error {
	return en.session.Send(context.Background(), "test notification", "email test")
}

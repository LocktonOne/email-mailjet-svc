package mailjet

import (
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/pkg/errors"
)

type Message struct {
	Destination string
	Subject     string
	Text        *string
	Html        *string
	From        *string
}

type Connector interface {
	Send(message Message) (string, error)
}

func NewConnector(cfg Mailjet) Connector {
	return &connector{
		cfg:    cfg,
		client: mailjet.NewMailjetClient(cfg.PublicAPIKey, cfg.PrivateAPIKey),
	}
}

type connector struct {
	cfg    Mailjet
	client *mailjet.Client
}

func (c *connector) Send(message Message) (string, error) {
	mjMes := c.buildMailjetMessage(message)
	res, err := c.client.SendMailV31(mjMes)
	if err != nil {
		return "", errors.Wrap(err, "failed to send message")
	}

	id := res.ResultsV31[0].CustomID

	return id, nil
}

func (c *connector) buildMailjetMessage(message Message) *mailjet.MessagesV31 {
	text := ""
	if message.Text != nil {
		text = *message.Text
	}

	from := c.cfg.FromEmail
	if message.From != nil {
		from = *message.From
	}

	messageInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: from,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31 {
					Email: message.Destination,
				},
			},
			Subject:  message.Subject,
			TextPart: text,
			HTMLPart: "",
		},
	}

	if message.Html != nil {
		messageInfo[0].HTMLPart = *message.Html
	}

	messages := mailjet.MessagesV31{Info: messageInfo }

	return &messages
}


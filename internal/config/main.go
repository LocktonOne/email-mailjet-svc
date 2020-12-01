package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/tokend/notifications/email-mailjet-svc/internal/mailjet"
)

type Config interface {
	comfig.Logger
	comfig.Listenerer
	Notificator() *Notificator
	NotificationsRouterRegistry() error
	mailjet.Mailjeter
}

type config struct {
	comfig.Logger
	comfig.Listenerer
	*notificatorer
	getter kv.Getter
	mailjet.Mailjeter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:        getter,
		Listenerer:    comfig.NewListenerer(getter),
		Logger:        comfig.NewLogger(getter, comfig.LoggerOpts{}),
		notificatorer: newNotificatorer(getter),
		Mailjeter:     mailjet.NewMailjeter(getter),
	}
}

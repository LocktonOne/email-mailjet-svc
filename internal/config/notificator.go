package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type notificatorer struct {
	getter kv.Getter
	once   comfig.Once
}

// Notificator is a config struct that defines common settings for the service
type Notificator struct {
	Channel          string `fig:"channel,required"`
	RegisterEndpoint string `fig:"register_endpoint,required"`
	Url              string `fig:"url,required"`
	Upstream         string `fig:"upstream,required"`
}

func newNotificatorer(getter kv.Getter) *notificatorer {
	return &notificatorer{
		getter: getter,
	}
}

// Notificator returns new instance of Notificator
func (m *notificatorer) Notificator() *Notificator {
	return m.once.Do(func() interface{} {
		var config Notificator
		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(m.getter, "notificator")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out notificator config"))
		}
		return &config
	}).(*Notificator)
}

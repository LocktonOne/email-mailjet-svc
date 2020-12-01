package mailjet

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Mailjeter interface {
	Mailjet() Connector
}

type mailjeter struct {
	getter kv.Getter
	once   comfig.Once
}

// Mailjet is a config structure for the Mailjet service
type Mailjet struct {
	PublicAPIKey  string `fig:"public_api_key,required"`
	PrivateAPIKey string `fig:"private_api_key,required"`
	FromEmail     string `fig:"from_email,required"`
}

func NewMailjeter(getter kv.Getter) Mailjeter {
	return &mailjeter{
		getter: getter,
	}
}

// Mailjet returns new instance of Mailjet
func (m *mailjeter) Mailjet() Connector {
	return m.once.Do(func() interface{} {
		var config Mailjet
		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(m.getter, "mailjet")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out mailjet config"))
		}
		return NewConnector(config)
	}).(Connector)
}

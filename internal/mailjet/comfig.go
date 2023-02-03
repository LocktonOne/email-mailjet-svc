package mailjet

import (
	"log"

	"github.com/ivanlele/envconfig"
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
	PublicAPIKey  string `fig:"public_api_key" envconfig:"MAILJET_PUBLIC_API_KEY" required:"true"`
	PrivateAPIKey string `fig:"private_api_key" envconfig:"MAILJET_PRIVATE_API_KEY" required:"true"`
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

		if config.PrivateAPIKey == "" || config.PublicAPIKey == "" {
			if err := envconfig.Process("MAILJET", &config); err != nil {
				return err
			}
		} else {
			log.Println("The use of public_api_key and private_api_key has been deprecated, please switch to using an environment variable instead")
		}

		return NewConnector(config)
	}).(Connector)
}

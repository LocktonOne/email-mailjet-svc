package config

import (
	"context"
	"net/http"
	"net/url"

	"gitlab.com/tokend/connectors/signed"

	jsonapi "gitlab.com/distributed_lab/json-api-connector"

	"github.com/pkg/errors"
)

func (c *config) NotificationsRouterRegistry() error {
	notificator := c.Notificator()
	url, err := url.Parse(notificator.Url)
	if err != nil {
		return errors.Wrap(err, "failed to parse url")
	}
	horizon := jsonapi.NewConnector(signed.NewClient(http.DefaultClient, url))
	endpoint, err := url.Parse(notificator.RegisterEndpoint)
	if err != nil {
		return errors.Wrap(err, "failed to parse notificator's endpoint")
	}

	request := registerServiceRequest{
		Data: registerServiceData{
			Type: "notificator-service",
			Attributes: registerServiceAttributes{
				Endpoint: notificator.Upstream,
				Channels: []string{notificator.Channel},
			},
		},
	}

	err = horizon.PostJSON(endpoint, request, context.Background(), nil)
	if err != nil {
		return errors.Wrap(err, "failed to send post request")
	}
	return nil
}

type registerServiceRequest struct {
	Data registerServiceData
}

type registerServiceData struct {
	Type       string
	Attributes registerServiceAttributes
}

type registerServiceAttributes struct {
	Endpoint string   `json:"endpoint"`
	Channels []string `json:"channels"`
}

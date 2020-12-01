package requests

import (
	"encoding/json"
	"net/http"

	"gitlab.com/tokend/notifications/email-mailjet-svc/internal/resources"

	"github.com/pkg/errors"
)

type CreateNotification struct {
	Notification resources.CreateNotificationResponse
}

func NewCreateNotification(r *http.Request) (*CreateNotification, error) {
	request := CreateNotification{}

	var body resources.CreateNotificationResponse
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	request.Notification = body

	return &request, request.validate()
}

func (c *CreateNotification) validate() error {
	return nil
}

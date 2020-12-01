package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/notifications/email-mailjet-svc/internal/mailjet"
	"gitlab.com/tokend/notifications/email-mailjet-svc/internal/service/requests"
	"net/http"
)

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateNotification(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	_, err = MailjetClient(r).Send(mailjet.Message{
		Destination: request.Notification.Data.Relationships.Destination.Data.ID,
		Subject:     request.Notification.Data.Attributes.Message.Attributes.Subject,
		Text:        request.Notification.Data.Attributes.Message.Attributes.Text,
		Html:        request.Notification.Data.Attributes.Message.Attributes.Html,
		From:        request.Notification.Data.Attributes.Message.Attributes.From,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to send notification")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

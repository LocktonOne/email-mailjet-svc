package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/notifications/email-mailjet-svc/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log), // this line may cause compilation error but in general case `dep ensure -v` will fix it
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxMailjetClient(s.cfg.Mailjet()),
		),
	)

	r.Post("/notifications", handlers.CreateNotification)

	return r
}

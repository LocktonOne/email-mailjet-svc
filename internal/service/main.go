package service

import (
	"net"
	"net/http"

	"github.com/pkg/errors"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/notifications/email-mailjet-svc/internal/config"
)

type service struct {
	log      *logan.Entry
	listener net.Listener
	cfg      config.Config
}

func (s *service) run() error {
	r := s.router()

	if err := s.cfg.NotificationsRouterRegistry(); err != nil {
		return errors.Wrap(err, "failed to register in notifications router")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		listener: cfg.Listener(),
		cfg:      cfg,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}

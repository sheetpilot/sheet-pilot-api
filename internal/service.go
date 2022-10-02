package internal

import (
	"context"
	"net/http"

	"sheetpilot/sheet-pilot-api/api/router"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ListenAddress string

	SheetPilotServiceAddress string
}

type Service struct {
	Router *router.Router
	Addr   string
}

func New(ctx context.Context, log *logrus.Entry, configs Config) (*Service, error) {
	router := router.New(router.Config{
		CorsHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		CorsMethods: []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})

	router.SetupRouter()
	log.Info("starting the router")

	return &Service{
		Router: router,
		Addr:   configs.ListenAddress,
	}, nil
}

func (s *Service) Start() {

	http.ListenAndServe(s.Addr, s.Router)
}

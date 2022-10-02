package internal

import (
	"context"
	"log"
	"net/http"

	"github.com/sheetpilot/sheet-pilot-api/api/router"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ListenAddress string

	SheetPilotServiceAddress string
}

type Service struct {
	Handler http.Handler
	Addr    string
}

func New(ctx context.Context, log *logrus.Entry, configs Config) (*Service, error) {
	router := router.New(router.Config{
		CorsHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		CorsMethods: []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})

	handler := router.SetupRouter()
	log.Info("starting the router")

	return &Service{
		Handler: handler,
		Addr:    configs.ListenAddress,
	}, nil
}

func (s *Service) Start() {
	log.Printf("public server Listening at %s", s.Addr)

	http.ListenAndServe(s.Addr, s.Handler)
}

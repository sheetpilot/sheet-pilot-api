package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/sheetpilot/sheet-pilot-api/internal"
)

var log *logrus.Entry

func init() {
	l := logrus.New()

	log = l.WithFields(logrus.Fields{
		"app": map[string]string{
			"host": os.Getenv("HOST"),
		},
	})
}

func main() {
	options := internal.Config{
		ListenAddress:            os.Getenv("LISTEN_ADDRESS"),
		SheetPilotServiceAddress: os.Getenv("SHEETPILOT_SERVICE_ADDRESS"),
	}

	svc, err := internal.New(context.Background(), log, options)
	if err != nil {
		log.WithError(err).Fatal("internal.New()")
	}

	svc.Start()
}

package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/sheetpilot/sheet-pilot-api/internal"
	"github.com/sheetpilot/sheet-pilot-api/internal/util"
)

var log *logrus.Entry

func init() {
	l := logrus.New()

	log = l.WithFields(logrus.Fields{
		"app": map[string]string{
			"host": util.GetEnv("HOST", "sheet-pilot-api"),
		},
	})
}

func main() {
	options := internal.Config{
		ListenAddress:            util.GetEnv("LISTEN_ADDRESS", ":4001"),
		SheetPilotServiceAddress: util.GetEnv("SHEETPILOT_SERVICE_ADDRESS", ":10001"),
	}

	svc, err := internal.New(context.Background(), log, options)
	if err != nil {
		log.WithError(err).Fatal("internal.New()")
	}

	svc.Start()
}

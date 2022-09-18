package main

import (
	"fmt"

	"github.com/b3lb/monitoring/pkg/app"
	"github.com/b3lb/monitoring/pkg/config"
	b3lbConfig "github.com/SLedunois/b3lb/v2/pkg/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	initLog()
	if err := run(); err != nil {
		panic(fmt.Errorf("failed to launch server: %s", err))
	}
}

func initLog() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
}

func run() error {
	conf, err := config.Load(b3lbConfig.Path())
	if err != nil {
		log.Error("Unable to load configuration")
		return err
	}

	err = app.NewServer(conf).Run()
	if err != nil {
		return err
	}

	return nil
}

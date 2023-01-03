package main

import (
	"flag"
	"fmt"

	bbsConfig "github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
	"github.com/bigblueswarm/monitoring/pkg/app"
	"github.com/bigblueswarm/monitoring/pkg/config"

	log "github.com/sirupsen/logrus"
)

var (
	configPath = ""
)

func main() {
	parseFlags()
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

func parseFlags() {
	flag.StringVar(&configPath, "config", config.DefaultConfigPath(), "Config file path")
	flag.Parse()
}

func run() error {
	configPath, err := bbsConfig.FormalizeConfigPath(configPath)
	if err != nil {
		panic(fmt.Errorf("unable to parse configuration: %s", err.Error()))
	}

	conf, err := config.Load(configPath)
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

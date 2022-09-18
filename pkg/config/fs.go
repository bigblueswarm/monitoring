// Package config provides the configuration loading
package config

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func loadConfigFromFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Error(fmt.Sprintf("unable to close config file: %s", err))
		}
	}()

	conf := &Config{}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	yaml.Unmarshal(b, &conf)

	return conf, nil
}

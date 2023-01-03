// Package config provides the configuration loading
package config

import (
	"fmt"

	"github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
)

const defaultConfigFileName = "monitoring.bbs.yaml"

// DefaultConfigPath return the default config path file
func DefaultConfigPath() string {
	return fmt.Sprintf("%s/%s", config.DefaultConfigFolder, defaultConfigFileName)
}

// Load configuration for the given path
func Load(path string) (*Config, error) {
	var conf *Config
	var err error
	if config.IsConsulEnabled(path) {
		conf, err = loadConfigFromConsul(path)
	} else {
		conf, err = loadConfigFromFile(path)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %s", err)
	}

	return conf, nil
}

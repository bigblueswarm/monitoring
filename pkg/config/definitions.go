// Package config provides the configuration loading
package config

import "github.com/bigblueswarm/bigblueswarm/v2/pkg/config"

// Config is the global application config
type Config struct {
	IDB        config.IDB            `yaml:"influxdb"`
	RDB        config.RDB            `yaml:"redis"`
	Balancer   config.BalancerConfig `yaml:"balancer"`
	Monitoring *MonitoringConfig     `yaml:"monitoring"`
}

// Port is the monitoring application port
type Port int

// MonitoringConfig is the global monitoring application config
type MonitoringConfig struct {
	Auth *string `yaml:"auth"`
	Port Port    `yaml:"port"`
}

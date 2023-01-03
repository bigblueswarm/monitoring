// Package config provides the configuration loading
package config

import (
	"fmt"
	"strings"

	"github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func getConsulKV(path string) (*api.KV, error) {
	conf := api.DefaultConfig()
	conf.Address = strings.ReplaceAll(path, config.ConsulPrefix, "")
	client, err := api.NewClient(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve KV api: %s", err)
	}

	config.SetConsulConfig(conf)

	return client.KV(), nil

}

func loadConfigFromConsul(path string) (*Config, error) {
	kv, err := getConsulKV(path)
	if err != nil {
		return nil, err
	}

	bbsConf := &config.Config{}
	if err := bbsConf.LoadInfluxDBConf(kv); err != nil {
		return nil, fmt.Errorf("failed to load influxdb configuration from consul: %s", err)
	}

	if err := bbsConf.LoadRedisConf(kv); err != nil {
		return nil, fmt.Errorf("failed to load redis configuration from consul: %s", err)
	}

	mc, err := loadMonitoringConfig(kv)
	if err != nil {
		return nil, fmt.Errorf("failed to load monitoring configuration from consul: %s", err)
	}

	return &Config{
		IDB:        bbsConf.IDB,
		RDB:        bbsConf.RDB,
		Monitoring: mc,
	}, nil
}

func loadMonitoringConfig(kv *api.KV) (*MonitoringConfig, error) {
	key := "monitoring"
	pair, _, err := kv.Get(config.ConsulKey(key), nil)

	if err != nil {
		return nil, err
	}

	var conf *MonitoringConfig
	if err := yaml.Unmarshal(pair.Value, &conf); err != nil {
		return nil, err
	}

	err = config.WatchChanges(log.WithField("key", key), key, func(value []byte) {
		if err := yaml.Unmarshal(value, &conf); err != nil {
			log.Error(fmt.Errorf("unable to parse new config value: %s", err))
			return
		}
	})

	if err != nil {
		return nil, fmt.Errorf("failed to watch changes: %s", err.Error())
	}

	return conf, nil
}

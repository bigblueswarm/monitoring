package config

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
	"github.com/bigblueswarm/test_utils/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConfigPath(t *testing.T) {
	t.Run("default config path should return $HOME/.bigblueswarm/monitoring.bbs.yaml", func(t *testing.T) {
		assert.Equal(t, "$HOME/.bigblueswarm/monitoring.bbs.yaml", DefaultConfigPath())
	})
}

func TestFSConfigLoad(t *testing.T) {

	type test struct {
		name  string
		path  string
		check func(t *testing.T, config *Config, err error)
	}

	tests := []test{
		{
			name: "Configuration loading does not returns any error with a valid path",
			path: "../../scripts/config.yml",
			check: func(t *testing.T, config *Config, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, config)
			},
		},
		{
			name: "Configuration loading returns an error with an invalid path",
			path: "config.yml",
			check: func(t *testing.T, config *Config, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, config)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config, err := Load(test.path)
			test.check(t, config, err)
		})
	}
}

func TestConsulConfigLoad(t *testing.T) {
	var url string
	var rdbConf string
	var idbConf string
	var monitoringConf string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := strings.ReplaceAll(r.RequestURI, "/v1/kv/configuration/", "")

		switch key {
		case "redis":
			w.Write([]byte(rdbConf))
		case "influxdb":
			w.Write([]byte(idbConf))
		case "monitoring":
			w.Write([]byte(monitoringConf))
		}
	}))

	defer server.Close()

	tests := []test.Test{
		{
			Name: "an invalid url should return an error",
			Mock: func() {
				url = "invalid_url:333333"
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "an error while loading influx configuration should return an error",
			Mock: func() {
				url = server.URL
				rdbConf = `[{"LockIndex":0,"Key":"configuration/redis","Flags":0,"Value":"YWRkcmVzczogbG9jYWxob3N0OjYzNzkKcGFzc3dvcmQ6CmRhdGFiYXNlOiAw","CreateIndex":46,"ModifyIndex":46}]`
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "an error while loading monitoring configuration should return an error",
			Mock: func() {
				idbConf = `[{"LockIndex":0,"Key":"configuration/influxdb","Flags":0,"Value":"YWRkcmVzczogaHR0cDovL2xvY2FsaG9zdDo4MDg2CnRva2VuOiBacTl3THNtaG5XNVV0T2lQSkFwVXYxY1RWSmZ3WHNUZ2xfcENraVRpa1EzZzJZR1B0UzVIcXNYZWYtV2Y1cFVVM3dqWTNuVldUWVJJLVdjOExqYkRmZz09Cm9yZ2FuaXphdGlvbjogYjNsYgpidWNrZXQ6IGJ1Y2tldA==","CreateIndex":50,"ModifyIndex":50}]`
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "no error should return a valid configuration",
			Mock: func() {
				monitoringConf = `[{"LockIndex":0,"Key":"configuration/monitoring","Flags":0,"Value":"YXV0aDogcGFzc3dvcmQKcG9ydDogODA5MQ==","CreateIndex":25,"ModifyIndex":406}]`
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				conf := value.(*Config)
				assert.Nil(t, err)
				auth := "password"
				expected := &Config{
					IDB: config.IDB{
						Address:      "http://localhost:8086",
						Token:        "Zq9wLsmhnW5UtOiPJApUv1cTVJfwXsTgl_pCkiTikQ3g2YGPtS5HqsXef-Wf5pUU3wjY3nVWTYRI-Wc8LjbDfg==",
						Organization: "b3lb",
						Bucket:       "bucket",
					},
					RDB: config.RDB{
						Address:  "localhost:6379",
						Password: "",
						DB:       0,
					},
					Monitoring: &MonitoringConfig{
						Auth: &auth,
						Port: 8091,
					},
				}

				assert.Equal(t, expected, conf)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Mock()
			config, err := Load(fmt.Sprintf("%s%s", config.ConsulPrefix, url))
			test.Validator(t, config, err)
		})
	}
}

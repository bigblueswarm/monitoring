package service

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var idbResponse string
var idbHttpStatus int
var csMock *ClusterService

func TestMain(m *testing.M) {
	influxDBServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(idbHttpStatus)
		rw.Header().Set("Content-Type", "application/csv")
		rw.Write([]byte(idbResponse))
	}))

	defer influxDBServer.Close()

	influxClient := influxdb2.NewClient(influxDBServer.URL, "token")

	csMock = &ClusterService{
		Service: Service{
			client: influxClient.QueryAPI("org"),
		},
	}
	status := m.Run()
	os.Exit(status)
}

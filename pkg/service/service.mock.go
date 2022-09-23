package service

import (
	"context"
	"errors"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
)

type InfluxDBQueryMock struct{}

var (
	// InfluxDBQueryMockFunction is the function that will be called when mock is called
	InfluxDBQueryMockFunction func(ctx context.Context, query string) (*influxdb2.QueryTableResult, error)
)

// QueryRaw executes flux query on the InfluxDB server and returns complete query result as a string with table annotations according to dialect
func (i *InfluxDBQueryMock) QueryRaw(ctx context.Context, query string, dialect *domain.Dialect) (string, error) {
	panic(errors.New("method has no implementation"))
}

// QueryRawWithParams executes flux parametrized query on the InfluxDB server and returns complete query result as a string with table annotations according to dialect

func (i *InfluxDBQueryMock) QueryRawWithParams(ctx context.Context, query string, dialect *domain.Dialect, params interface{}) (string, error) {
	panic(errors.New("method has no implementation"))
}

// Query executes flux query on the InfluxDB server and returns QueryTableResult which parses streamed response into structures representing flux table parts
func (i *InfluxDBQueryMock) Query(ctx context.Context, query string) (*influxdb2.QueryTableResult, error) {
	return InfluxDBQueryMockFunction(ctx, query)
}

// QueryWithParams executes flux parametrized query  on the InfluxDB server and returns QueryTableResult which parses streamed response into structures representing flux table parts
func (i *InfluxDBQueryMock) QueryWithParams(ctx context.Context, query string, params interface{}) (*influxdb2.QueryTableResult, error) {
	panic(errors.New("method has no implementation"))
}

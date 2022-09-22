// Package service provide every services used by application
package service

import (
	"context"

	"github.com/b3lb/monitoring/pkg/model"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// NewClusterService instantiate a new ClusterService object
func NewClusterService(idb influxdb2.Client, org string, b string) IClusterService {
	return &ClusterService{
		client: idb.QueryAPI(org),
		bucket: b,
	}
}

// GetUserCount retrieve active users count
func (c *ClusterService) GetUserCount() (*model.ActiveUserGauge, error) {
	q := `
	from(bucket: "bucket")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "bigbluebutton_meetings")
		|> filter(fn: (r) => r["_field"] == "participant_count")
		|> group(columns: ["_time", "_measurement"], mode:"by")
		|> last(column: "_time")
		|> yield(name: "sum")
	`

	c.client.Query(context.Background(), q)
	val := 123
	return &model.ActiveUserGauge{
		Value: &val,
	}, nil
}

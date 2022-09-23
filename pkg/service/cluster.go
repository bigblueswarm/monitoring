// Package service provide every services used by application
package service

import (
	"github.com/b3lb/monitoring/pkg/model"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// NewClusterService instantiate a new ClusterService object
func NewClusterService(idb influxdb2.Client, org string, b string, a string) IClusterService {
	return &ClusterService{
		Service: Service{
			client: idb.QueryAPI(org),
			bucket: b,
			aggregationInterval: a,
		},
	}
}

// GetUserCount retrieve active users count
func (c *ClusterService) GetUserCount() (*model.Gauge, error) {
	return c.getGaugeValue(bigbluebuttonMeetings, participantCount)
}

// GetUserTimeserie retrieve active users in the cluster as a timeserie
func (c *ClusterService) GetUserTimeserie(start string, stop string, every string) ([]*model.Point, error) {
	return c.getTimeserie(bigbluebuttonMeetings, participantCount, start, stop, every)
}

// GetUserTrend retrive active users trend in the cluster
func (c *ClusterService) GetUserTrend(start string, stop string, every string) (*model.Trend, error) {
	return c.getTrend(bigbluebuttonMeetings, participantCount, start, stop, every)
}

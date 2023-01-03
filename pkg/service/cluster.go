// Package service provide every services used by application
package service

import (
	"github.com/bigblueswarm/monitoring/pkg/model"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// NewClusterService instantiate a new ClusterService object
func NewClusterService(idb influxdb2.Client, org string, b string, a string) IClusterService {
	return &ClusterService{
		Service: Service{
			client:              idb.QueryAPI(org),
			bucket:              b,
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

// GetMeetingCount retrieve actives meetings in the cluster
func (c *ClusterService) GetMeetingCount() (*model.Gauge, error) {
	return c.getGaugeValue(bigbluebuttonMeetings, activeMeetings)
}

// GetMeetingTimeserie retrieve active meetings in the cluster as a timeserie
func (c *ClusterService) GetMeetingTimeserie(start string, stop string, every string) ([]*model.Point, error) {
	return c.getTimeserie(bigbluebuttonMeetings, activeMeetings, start, stop, every)
}

// GetMeetingTrend retrieve active meetings trend in the cluster
func (c *ClusterService) GetMeetingTrend(start string, stop string, every string) (*model.Trend, error) {
	return c.getTrend(bigbluebuttonMeetings, activeMeetings, start, stop, every)
}

// GetRecordingsCount retrieve actives meetings in the cluster
func (c *ClusterService) GetRecordingsCount() (*model.Gauge, error) {
	return c.getGaugeValue(bigbluebuttonMeetings, activeRecordings)
}

// GetRecoringTimeserie retrieve active meetings in the cluster as a timeserie
func (c *ClusterService) GetRecoringTimeserie(start string, stop string, every string) ([]*model.Point, error) {
	return c.getTimeserie(bigbluebuttonMeetings, activeRecordings, start, stop, every)
}

// GetRecordingTrend retrieve active meetings trend in the cluster
func (c *ClusterService) GetRecordingTrend(start string, stop string, every string) (*model.Trend, error) {
	return c.getTrend(bigbluebuttonMeetings, activeRecordings, start, stop, every)
}

// GetAggregationInterval returns the aggregation interval configured in balancer
func (c *ClusterService) GetAggregationInterval() string {
	return c.aggregationInterval
}

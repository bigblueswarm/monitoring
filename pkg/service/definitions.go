// Package service provide every services used by application
package service

import (
	"github.com/bigblueswarm/monitoring/pkg/model"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2/api"
)

const bigbluebuttonMeetings = "bigbluebutton_meetings"
const participantCount = "participant_count"

// IClusterService is the interface that manage cluster data
type IClusterService interface {
	// GetUserCount retrieve actives users in the cluster
	GetUserCount() (*model.Gauge, error)
	// GetUserTimeserie retrieve active users in the cluster as a timeserie
	GetUserTimeserie(start string, stop string, every string) ([]*model.Point, error)
	// GetUserTrend retrive active users trend in the cluster
	GetUserTrend(start string, stop string, every string) (*model.Trend, error)
}

// Service is a common struct that represents a service
type Service struct {
	client              influxdb2.QueryAPI
	bucket              string
	aggregationInterval string
}

// ClusterService manage cluster data
type ClusterService struct {
	Service
}

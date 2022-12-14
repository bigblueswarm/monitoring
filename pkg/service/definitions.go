// Package service provide every services used by application
package service

import (
	"github.com/bigblueswarm/monitoring/pkg/model"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2/api"
)

const (
	bigbluebuttonMeetings = "bigbluebutton_meetings"
	participantCount      = "participant_count"
	activeMeetings        = "active_meetings"
	activeRecordings      = "active_recordings"
)

// IClusterService is the interface that manage cluster data
type IClusterService interface {
	// GetUserCount retrieve actives users in the cluster
	GetUserCount() (*model.Gauge, error)
	// GetUserTimeserie retrieve active users in the cluster as a timeserie
	GetUserTimeserie(start string, stop string, every string) ([]*model.Point, error)
	// GetUserTrend retrieve active users trend in the cluster
	GetUserTrend(start string, stop string, every string) (*model.Trend, error)
	// GetMeetingsCount retrieve actives meetings in the cluster
	GetMeetingCount() (*model.Gauge, error)
	// GetMeetingTimeserie retrieve active meetings in the cluster as a timeserie
	GetMeetingTimeserie(start string, stop string, every string) ([]*model.Point, error)
	// GetMeetingTrend retrieve active meetings trend in the cluster
	GetMeetingTrend(start string, stop string, every string) (*model.Trend, error)
	// GetRecordingsCount retrieve actives meetings in the cluster
	GetRecordingsCount() (*model.Gauge, error)
	// GetRecoringTimeserie retrieve active meetings in the cluster as a timeserie
	GetRecoringTimeserie(start string, stop string, every string) ([]*model.Point, error)
	// GetRecordingTrend retrieve active meetings trend in the cluster
	GetRecordingTrend(start string, stop string, every string) (*model.Trend, error)
	// GetAggregationInterval returns the aggregation interval configured in balancer
	GetAggregationInterval() string
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

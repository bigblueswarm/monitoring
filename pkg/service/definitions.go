// Package service provide every services used by application
package service

import (
	"github.com/b3lb/monitoring/pkg/model"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2/api"
)

// IClusterService is the interface that manage cluster data
type IClusterService interface {
	// GetUserCount retrieve actives users in the cluster
	GetUserCount() (*model.ActiveUserGauge, error)
}

// ClusterService manage cluster data
type ClusterService struct {
	client influxdb2.QueryAPI
	bucket string
}

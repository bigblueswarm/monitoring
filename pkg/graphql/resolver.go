package graphql

import "github.com/b3lb/monitoring/pkg/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ClusterService service.IClusterService
}

package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/b3lb/monitoring/pkg/graphql/generated"
	"github.com/b3lb/monitoring/pkg/model"
)

// ActiveUsers is the resolver for the activeUsers field.
func (r *queryResolver) ActiveUsers(ctx context.Context) (*model.ActiveUserGauge, error) {
	return r.ClusterService.GetUserCount()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

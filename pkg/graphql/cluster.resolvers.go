package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/b3lb/monitoring/pkg/graphql/generated"
	"github.com/b3lb/monitoring/pkg/model"
	"github.com/b3lb/monitoring/pkg/pointer"
)

// ActiveUsers is the resolver for the activeUsers field.
func (r *queryResolver) ActiveUsers(ctx context.Context, start *string, stop *string, every *string) (*model.ActiveUsersStat, error) {
	if every == (*string)(nil) {
		every = pointer.SPtr("")
	}

	gauge, err := r.ClusterService.GetUserCount()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user count gauge: %s", err)
	}

	sparkline, err := r.ClusterService.GetUserTimeserie(*start, *stop, "45s")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user count sparkline: %s", err)
	}

	trend, err := r.ClusterService.GetUserTrend(*start, *stop, *every)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user trend: %s", err)
	}

	return &model.ActiveUsersStat{
		Gauge:     gauge,
		Trend:     trend,
		Sparkline: sparkline,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bigblueswarm/monitoring/pkg/graphql/generated"
	"github.com/bigblueswarm/monitoring/pkg/model"
)

// ActiveUsers is the resolver for the activeUsers field.
func (r *queryResolver) ActiveUsers(ctx context.Context, start *string, stop *string) (*model.ActiveMetricStat, error) {
	gauge, err := r.ClusterService.GetUserCount()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user count gauge: %s", err)
	}

	sparkline, err := r.ClusterService.GetUserTimeserie(*start, *stop, "45s")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user count sparkline: %s", err)
	}

	trend, err := r.ClusterService.GetUserTrend(*start, *stop, r.ClusterService.GetAggregationInterval())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user trend: %s", err)
	}

	return &model.ActiveMetricStat{
		Gauge:     gauge,
		Trend:     trend,
		Sparkline: sparkline,
	}, nil
}

// ActiveMeetings is the resolver for the activeMeetings field.
func (r *queryResolver) ActiveMeetings(ctx context.Context, start *string, stop *string) (*model.ActiveMetricStat, error) {
	gauge, err := r.ClusterService.GetMeetingCount()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve meeting count gauge: %s", err)
	}

	sparkline, err := r.ClusterService.GetMeetingTimeserie(*start, *stop, "45s")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve meetings count sparkline: %s", err)
	}

	trend, err := r.ClusterService.GetMeetingTrend(*start, *stop, r.ClusterService.GetAggregationInterval())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve meetings trend: %s", err)
	}

	return &model.ActiveMetricStat{
		Gauge:     gauge,
		Trend:     trend,
		Sparkline: sparkline,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

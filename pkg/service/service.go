// Package service provide every services used by application
package service

import (
	"context"
	"fmt"
	"math"

	"github.com/bigblueswarm/monitoring/pkg/model"
)

func (s *Service) getGaugeValue(measurement string, field string) (*model.Gauge, error) {
	q := fmt.Sprintf(`
	from(bucket: "%s")
		|> range(start: -60s)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["_field"] == "%s")
		|> aggregateWindow(every: %s, fn: sum, createEmpty: false)
		|> last()
		|> map(fn: (r) => ({r with _value: float(v: r._value)}))
	`, s.bucket, measurement, field, s.aggregationInterval)

	result, err := s.client.Query(context.Background(), q)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve gauge value: %s", err)
	}

	val := float64(0)
	if result.Next() {
		val = result.Record().Value().(float64)
	}

	if result.Err() != nil {
		return nil, fmt.Errorf("gauge request return an error: %s", result.Err())
	}

	return &model.Gauge{
		Value: &val,
	}, nil
}

func (s *Service) getTimeserie(measurement string, field string, start string, stop string, every string) ([]*model.Point, error) {
	if every == "" {
		every = s.aggregationInterval
	}

	q := fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: %s, stop: %s)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["_field"] == "%s")
		|> group(columns: ["_field"])
		|> aggregateWindow(every: %s, fn: sum, createEmpty: false)
		|> map(fn: (r) => ({r with _value: float(v: r._value)}))
		`, s.bucket, start, stop, measurement, field, every)

	result, err := s.client.Query(context.Background(), q)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve timeserie: %s", err)
	}

	points := []*model.Point{}

	for result.Next() {
		record := result.Record()
		t := record.Time().String()
		v := record.Value().(float64)
		point := &model.Point{
			Time:  &t,
			Value: &v,
		}

		points = append(points, point)
	}

	if result.Err() != nil {
		return nil, fmt.Errorf("timeserie request return an error: %s", result.Err())
	}

	return points, nil
}

func (s *Service) getTrend(measurement string, field string, start string, stop string, every string) (*model.Trend, error) {
	if every == "" {
		every = s.aggregationInterval
	}

	q := fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: %s, stop: %s)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["_field"] == "%s")
		|> group(columns: ["_field"])
		|> aggregateWindow(every: %s, fn: sum, createEmpty: false)
		|> sort(columns: ["_time"], desc: true)
		|> limit(n: 2)
		|> map(fn: (r) => ({r with _value: float(v: r._value)}))
		`, s.bucket, start, stop, measurement, field, every)

	result, err := s.client.Query(context.Background(), q)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve trend: %s", err)
	}

	value := float64(0)
	if result.Next() {
		val := result.Record().Value().(float64)
		previousVal := float64(0)
		if result.Next() {
			previousVal = result.Record().Value().(float64)
		}

		value = ((val - previousVal) / previousVal) * 100
		value = math.Round(value*100) / 100
	}

	if result.Err() != nil {
		return nil, fmt.Errorf("trend request return an error: %s", result.Err())
	}

	return &model.Trend{
		Value: &value,
	}, nil
}

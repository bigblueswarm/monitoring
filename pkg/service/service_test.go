package service

import (
	"net/http"
	"testing"

	"github.com/b3lb/monitoring/pkg/model"
	"github.com/b3lb/monitoring/pkg/pointer"
	"github.com/b3lb/test_utils/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestGetGaugeValue(t *testing.T) {
	tests := []test.Test{
		{
			Name: "an error returned by influxdb should return an error",
			Mock: func() {
				idbResponse = ``
				idbHttpStatus = http.StatusInternalServerError
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			Name: "a empty response should return a gauge value equals to 0",
			Mock: func() {
				idbResponse = ``
				idbHttpStatus = http.StatusOK
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				gauge := value.(*model.Gauge)
				assert.Equal(t, float64(0), *gauge.Value)
			},
		},
		{
			Name: "a valid response should return a valid gauge equals to 210",
			Mock: func() {
				idbResponse = `#group,false,false,true,true,true,true,false,false
#datatype,string,long,string,string,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double
#default,_result,,,,,,,
,result,table,_field,_measurement,_start,_stop,_time,_value
,,0,participant_count,bigbluebutton_meetings,2022-10-11T15:00:33.120844652Z,2022-10-11T15:05:33.120844652Z,2022-10-11T15:05:30Z,210`
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				gauge := value.(*model.Gauge)
				assert.Equal(t, float64(210), *gauge.Value)
			},
		},
	}

	for _, test := range tests {
		test.Mock()
		gauge, err := csMock.getGaugeValue(bigbluebuttonMeetings, participantCount)
		test.Validator(t, gauge, err)
	}
}

func TestGetTimeserie(t *testing.T) {
	tests := []test.Test{
		{
			Name: "an error returned by influxdb should return an error",
			Mock: func() {
				idbResponse = ``
				idbHttpStatus = http.StatusInternalServerError
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			Name: "an empty response should return an empty points array",
			Mock: func() {
				idbHttpStatus = http.StatusOK
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
				res := value.([]*model.Point)
				assert.Equal(t, 0, len(res))
			},
		},
		{
			Name: "A valid request should parse points and return a valid point array",
			Mock: func() {
				idbResponse = `#group,false,false,true,true,true,false,false
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,string,double,dateTime:RFC3339
#default,_result,,,,,,
,result,table,_start,_stop,_field,_value,_time
,,0,2022-10-09T01:23:00.011171912Z,2022-10-09T01:23:30.011171912Z,participant_count,6795.75,2022-10-09T01:23:10Z
,,0,2022-10-09T01:23:00.011171912Z,2022-10-09T01:23:30.011171912Z,participant_count,7001,2022-10-09T01:23:20Z`
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				expected := []*model.Point{
					{
						Time:  pointer.SPtr("2022-10-09 01:23:10 +0000 UTC"),
						Value: pointer.F64Ptr(6795.75),
					},
					{
						Time:  pointer.SPtr("2022-10-09 01:23:20 +0000 UTC"),
						Value: pointer.F64Ptr(7001),
					},
				}
				res := value.([]*model.Point)
				assert.Equal(t, expected, res)
			},
		},
	}
	for _, test := range tests {
		test.Mock()
		timeserie, err := csMock.getTimeserie(bigbluebuttonMeetings, participantCount, "-30s", "now()", "10s")
		test.Validator(t, timeserie, err)
	}
}

func TestGetTrend(t *testing.T) {
	tests := []test.Test{
		{
			Name: "an error returned by influxdb should return an error",
			Mock: func() {
				idbResponse = ``
				idbHttpStatus = http.StatusInternalServerError
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			Name: "a empty response should return a trend value equals to 0",
			Mock: func() {
				idbResponse = ``
				idbHttpStatus = http.StatusOK
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				gauge := value.(*model.Trend)
				assert.Equal(t, float64(0), *gauge.Value)
			},
		},
		{
			Name: "a valid request should return a trend equals to ",
			Mock: func() {
				idbResponse = `#group,false,false,true,true,true,false,false
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,string,double,dateTime:RFC3339
#default,_result,,,,,,
,result,table,_start,_stop,_field,_value,_time
,,0,2022-10-09T01:39:04.099691144Z,2022-10-09T01:39:34.099691144Z,participant_count,6898.25,2022-10-09T01:39:30Z
,,0,2022-10-09T01:39:04.099691144Z,2022-10-09T01:39:34.099691144Z,participant_count,6859,2022-10-09T01:39:20Z`
			},
			Validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
				res := value.(*model.Trend)
				expected := &model.Trend{
					Value: pointer.F64Ptr(0.57),
				}
				assert.Equal(t, expected, res)
			},
		},
	}

	for _, test := range tests {
		test.Mock()
		trend, err := csMock.getTrend(bigbluebuttonMeetings, participantCount, "-30s", "now()", "10s")
		test.Validator(t, trend, err)
	}
}

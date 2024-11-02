package model

import (
	"nebeng-jek/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToTrackUserLocationRequest(t *testing.T) {
	data := map[string]interface{}{
		"ride_id":   111,
		"timestamp": 12345678,
		"location": map[string]float64{
			"longitude": 1.111,
			"latitude":  2.0001,
		},
	}
	expected := TrackUserLocationRequest{
		RideID:    111,
		Timestamp: 12345678,
		Location: Coordinate{
			Longitude: 1.111, Latitude: 2.0001,
		},
	}

	t.Run("return track user location req", func(t *testing.T) {
		res, err := ToTrackUserLocationRequest(data)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})
	t.Run("return error", func(t *testing.T) {
		_, err := ToTrackUserLocationRequest(utils.FailingType{})
		assert.Error(t, err)
	})
}

package haversine

import (
	"math"
)

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// Function to calculate the distance between two coordinates
func Calculate(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Radius of the Earth in km

	// Convert coordinates from degrees to radians
	lat1Rad := toRadians(lat1)
	lon1Rad := toRadians(lon1)
	lat2Rad := toRadians(lat2)
	lon2Rad := toRadians(lon2)

	// Differences between coordinates
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Apply the Haversine formula
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return distance
}

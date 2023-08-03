package utils

import (
	"math"
)

const (
	earthRadius = 6371 // Earth's radius in kilometers
)

// ToRadians converts degrees to radians.
func ToRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// ToDegrees converts radians to degrees.
func ToDegrees(rad float64) float64 {
	return rad * (180 / math.Pi)
}

// CalculateNewCoordinates calculates the new latitude and longitude based on initial coordinates, bearing, and distance.
func CalculateNewCoordinates(initialLat, initialLon, bearing, distance float64) (newLat, newLon float64) {
	// Convert to radians
	initialLatRad := ToRadians(initialLat)
	initialLonRad := ToRadians(initialLon)
	bearingRad := ToRadians(bearing)

	// Calculate new latitude
	newLatRad := math.Asin(math.Sin(initialLatRad)*math.Cos(distance/earthRadius) +
		math.Cos(initialLatRad)*math.Sin(distance/earthRadius)*math.Cos(bearingRad))

	// Calculate new longitude
	newLonRad := initialLonRad + math.Atan2(math.Sin(bearingRad)*math.Sin(distance/earthRadius)*math.Cos(initialLatRad),
		math.Cos(distance/earthRadius)-math.Sin(initialLatRad)*math.Sin(newLatRad))

	// Convert back to degrees
	newLat = ToDegrees(newLatRad)
	newLon = ToDegrees(newLonRad)

	return newLat, newLon
}

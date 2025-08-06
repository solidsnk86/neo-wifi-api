package utils

import (
	"math"
	"neo-wifi-api/types"
)

const EARTH_RADIUS = 6378137.0

// square returns the square of a number
func square(num float64) float64 {
	return num * num
}

// degreesToRadians converts degrees to radians
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

// Haversine calculates the distance between two coordinates using the Haversine formula
func Haversine(locationA, locationB interface{}) float64 {
	var latA, lonA, latB, lonB float64

	// Handle different data types
	switch a := locationA.(type) {
	case types.Coords:
		latA = a.Lat
		lonA = a.Lon
		if latA == 0 && a.Latitude != 0 {
			latA = a.Latitude
		}
		if lonA == 0 && a.Longitude != 0 {
			lonA = a.Longitude
		}
	case types.Antenna:
		latA = a.Lat
		lonA = a.Lon
	case types.City:
		latA = a.Lat
		lonA = a.Lon
	case types.Airport:
		latA = a.Lat
		lonA = a.Lon
	}

	switch b := locationB.(type) {
	case types.Coords:
		latB = b.Lat
		lonB = b.Lon
		if latB == 0 && b.Latitude != 0 {
			latB = b.Latitude
		}
		if lonB == 0 && b.Longitude != 0 {
			lonB = b.Longitude
		}
	case types.Antenna:
		latB = b.Lat
		lonB = b.Lon
	case types.City:
		latB = b.Lat
		lonB = b.Lon
	case types.Airport:
		latB = b.Lat
		lonB = b.Lon
	}

	latitudeA := degreesToRadians(latA)
	latitudeB := degreesToRadians(latB)
	longitudeA := degreesToRadians(lonA)
	longitudeB := degreesToRadians(lonB)

	formula := square(math.Sin((latitudeB-latitudeA)/2)) +
		math.Cos(latitudeA)*
			math.Cos(latitudeB)*
			square(math.Sin((longitudeB-longitudeA)/2))

	distance := 2 * EARTH_RADIUS * math.Asin(math.Sqrt(formula))
	return distance / 1000
}

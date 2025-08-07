package utils

import (
	"math"
	"neo-wifi-api/types"
)

// GetClosest finds the three closest targets to given coordinates
func GetClosest(coordinates types.Coords, allData interface{}) types.ClosestResult {
	var closestTarget interface{}
	var secondClosestTarget interface{}
	var thirdClosestTarget interface{}

	minDistance := math.Inf(1)
	secondMinDistance := math.Inf(1)
	thirdMinDistance := math.Inf(1)

	coords := types.Coords{Lat: 0, Lon: 0}
	secondCoords := types.Coords{Lat: 0, Lon: 0}
	thirdCoords := types.Coords{Lat: 0, Lon: 0}

	// Handle different data types
	switch data := allData.(type) {
	case []types.Antenna:
		for _, item := range data {
			distance := Haversine(coordinates, item)

			if distance < minDistance {
				thirdMinDistance = secondMinDistance
				thirdClosestTarget = secondClosestTarget
				thirdCoords = secondCoords

				secondMinDistance = minDistance
				secondClosestTarget = closestTarget
				secondCoords = coords

				minDistance = distance
				closestTarget = item
				coords = types.Coords{Lat: item.Lat, Lon: item.Lon}
			} else if distance < secondMinDistance {
				thirdMinDistance = secondMinDistance
				thirdClosestTarget = secondClosestTarget
				thirdCoords = secondCoords

				secondMinDistance = distance
				secondClosestTarget = item
				secondCoords = types.Coords{Lat: item.Lat, Lon: item.Lon}
			} else if distance < thirdMinDistance {
				thirdMinDistance = distance
				thirdClosestTarget = item
				thirdCoords = types.Coords{Lat: item.Lat, Lon: item.Lon}
			}
		}
	case []types.City:
		for _, item := range data {
			distance := Haversine(coordinates, item)

			if distance < minDistance {
				minDistance = distance
				closestTarget = item
				coords = types.Coords{Lat: item.Lat, Lon: item.Lon}
			}
		}
	case []types.Airport:
		for _, item := range data {
			distance := Haversine(coordinates, item)

			if distance < minDistance {
				minDistance = distance
				closestTarget = item
				coords = types.Coords{Lat: item.Lat, Lon: item.Lon}
			}
		}
	}

	return types.ClosestResult{
		ClosestTarget:       closestTarget,
		SecondClosestTarget: secondClosestTarget,
		ThirdClosestTarget:  thirdClosestTarget,
		MinDistance:         minDistance,
		SecondMinDistance:   secondMinDistance,
		ThirdMinDistance:    thirdMinDistance,
		Coords:              coords,
		SecondCoords:        secondCoords,
		ThirdCoords:         thirdCoords,
	}
}

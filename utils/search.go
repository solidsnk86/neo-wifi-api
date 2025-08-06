package utils

import (
	"math"
	"neo-wifi-api/types"
	"strings"
)

// SearchAntenna searches for a specific antenna by name
func SearchAntenna(coordinates types.Coords, allData []types.Antenna, query string) types.SearchAntennaResult {
	targetDistance := math.Inf(1)
	var searchedTarget, searchedTarget5g, mac, mac5, antennaType string
	coords := types.Coords{Latitude: 0, Longitude: 0}

	for _, data := range allData {
		distance := Haversine(coordinates, data)
		nameA := strings.ToLower(data.Name)
		nameB := strings.ToLower(data.Name5g)

		if query == nameA || query == nameB {
			searchedTarget = data.Name
			searchedTarget5g = data.Name5g
			targetDistance = distance
			coords.Latitude = data.Lat
			coords.Longitude = data.Lon
			mac = data.MAC
			mac5 = data.MAC5g
			antennaType = data.Type
			break
		}
	}

	return types.SearchAntennaResult{
		TargetDistance:   targetDistance,
		SearchedTarget:   searchedTarget,
		SearchedTarget5g: searchedTarget5g,
		Coordinates:      coords,
		MAC:              mac,
		MAC5:             mac5,
		Type:             antennaType,
	}
}

// WriteMAC formats MAC address
func WriteMAC(mac string) string {
	if mac == "" {
		return "No disponible"
	}
	return strings.ReplaceAll(mac, " ", "-")
}

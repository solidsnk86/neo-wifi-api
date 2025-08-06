package utils

import (
	"math"
	"sort"
)

type Coord struct {
	Lat  float64
	Lon  float64
	Name string
}

type Result struct {
	Point    Coord
	Distance float64
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371.0 // Radio de la Tierra en km
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180
	lat1 = lat1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Asin(math.Sqrt(a))
	return R * c
}

// FindClosestN devuelve los N puntos más cercanos ordenados por distancia
func FindClosestN(origin Coord, list []Coord, n int) []Result {
	var results []Result

	for _, point := range list {
		dist := haversine(origin.Lat, origin.Lon, point.Lat, point.Lon)
		results = append(results, Result{Point: point, Distance: dist})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})

	if len(results) < n {
		return results
	}
	return results[:n]
}

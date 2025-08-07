package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"neo-wifi-api/types"
	"neo-wifi-api/utils"
	"net/http"
	"strconv"
	"strings"
)

// Global data
var antennas []types.Antenna
var cities []types.City
var airports []types.Airport

// LoadData loads JSON data from files
func LoadData() error {
	// Load antennas
	antennaData, err := ioutil.ReadFile("data/wifi-v15.json")
	if err != nil {
		return fmt.Errorf("error loading antennas: %v", err)
	}
	if err := json.Unmarshal(antennaData, &antennas); err != nil {
		return fmt.Errorf("error parsing antennas: %v", err)
	}

	// Load cities
	cityData, err := ioutil.ReadFile("data/geodata-v4-mgc.json")
	if err != nil {
		return fmt.Errorf("error loading cities: %v", err)
	}
	if err := json.Unmarshal(cityData, &cities); err != nil {
		return fmt.Errorf("error parsing cities: %v", err)
	}

	// Load airports
	airportData, err := ioutil.ReadFile("data/airports.json")
	if err != nil {
		return fmt.Errorf("error loading airports: %v", err)
	}
	if err := json.Unmarshal(airportData, &airports); err != nil {
		return fmt.Errorf("error parsing airports: %v", err)
	}

	return nil
}

// WifiHandler handles the main API endpoint
func WifiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse query parameters
	latStr := r.URL.Query().Get("lat")
	lonStr := r.URL.Query().Get("lon")
	query := strings.ToLower(r.URL.Query().Get("query"))

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil || lat == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Debes proporcionar los parámetros de latitud y longitud",
		})
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil || lon == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Debes proporcionar los parámetros de latitud y longitud",
		})
		return
	}

	clientIP := r.Header.Get("X-Real-IP")
	if clientIP == "" {
		clientIP = "No disponible"
	}

	coords := types.Coords{Lat: lat, Lon: lon}

	// Get closest city
	cityResult := utils.GetClosest(coords, cities)
	closestCity, _ := cityResult.ClosestTarget.(types.City)

	// Get closest antennas
	antennaResult := utils.GetClosest(coords, antennas)
	closestAntenna, _ := antennaResult.ClosestTarget.(types.Antenna)
	secondAntenna, _ := antennaResult.SecondClosestTarget.(types.Antenna)
	thirdAntenna, _ := antennaResult.ThirdClosestTarget.(types.Antenna)

	// Get closest airport
	airportResult := utils.GetClosest(coords, airports)
	closestAirport, _ := airportResult.ClosestTarget.(types.Airport)

	w.Header().Set("Content-Type", "application/json")

	// Handle query search
	if query != "" {
		searchResult := utils.SearchAntenna(coords, antennas, query)

		coordsStr := "No disponible"
		if searchResult.Coordinates.Latitude != 0 && searchResult.Coordinates.Longitude != 0 {
			coordsStr = fmt.Sprintf("%.6f,%.6f", searchResult.Coordinates.Latitude, searchResult.Coordinates.Longitude)
		}

		response := types.QueryResponse{
			Antenna: types.AntennaInfo{
				Name:   searchResult.SearchedTarget,
				Name5g: searchResult.SearchedTarget5g,
			},
			Distance: fmt.Sprintf("%.3fmts", searchResult.TargetDistance),
			Coords:   coordsStr,
			MAC:      utils.WriteMAC(searchResult.MAC),
			MAC5G:    utils.WriteMAC(searchResult.MAC5),
			Type:     searchResult.Type,
		}

		if searchResult.SearchedTarget == "" {
			response.Antenna.Name = "Antena inexistente"
		}
		if searchResult.SearchedTarget5g == "" {
			response.Antenna.Name5g = "No disponible"
		}
		if searchResult.Type == "" {
			response.Type = "No disponible"
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// Full response
	response := types.FullResponse{
		IP:          clientIP,
		City:        closestCity.Nombre,
		State:       closestCity.Provincia,
		Country:     closestCity.Pais,
		Departament: closestCity.Departamento,
		CityCoords: types.CityCoords{
			Latitude:  closestCity.Lat,
			Longitude: closestCity.Lon,
		},
		CenterDistance: fmt.Sprintf("%.3fmts", cityResult.MinDistance),
		CurrentPosition: types.CurrentPosition{
			Latitude:  lat,
			Longitude: lon,
		},
		ClosestWifi: types.WifiInfo{
			Antenna:  closestAntenna.Name,
			Name:     closestAntenna.Name5g,
			Distance: fmt.Sprintf("%.3fmts", antennaResult.MinDistance),
			Type:     closestAntenna.Type,
			MAC:      utils.WriteMAC(closestAntenna.MAC),
			MAC5G:    utils.WriteMAC(closestAntenna.MAC5g),
			Coords:   fmt.Sprintf("%.6f,%.6f", closestAntenna.Lat, closestAntenna.Lon),
			Users:    closestAntenna.Users,
		},
		SecondWifi: types.WifiInfo{
			Antenna:  secondAntenna.Name,
			Name:     secondAntenna.Name5g,
			Distance: fmt.Sprintf("%.3fmts", antennaResult.SecondMinDistance),
			Type:     secondAntenna.Type,
			MAC:      utils.WriteMAC(secondAntenna.MAC),
			MAC5G:    utils.WriteMAC(secondAntenna.MAC5g),
			Coords:   fmt.Sprintf("%.6f,%.6f", secondAntenna.Lat, secondAntenna.Lon),
			Users:    secondAntenna.Users,
		},
		ThirdWifi: types.WifiInfo{
			Antenna:  thirdAntenna.Name,
			Name:     thirdAntenna.Name5g,
			Distance: fmt.Sprintf("%.3fmts", antennaResult.ThirdMinDistance),
			Type:     thirdAntenna.Type,
			MAC:      utils.WriteMAC(thirdAntenna.MAC),
			MAC5G:    utils.WriteMAC(thirdAntenna.MAC5g),
			Coords:   fmt.Sprintf("%.6f,%.6f", thirdAntenna.Lat, thirdAntenna.Lon),
			Users:    thirdAntenna.Users,
		},
		AirportLocation: types.AirportLocation{
			City:    closestAirport.State,
			Country: closestAirport.Country,
			ClosestAirport: types.ClosestAirport{
				Airport:  closestAirport.Name,
				Distance: fmt.Sprintf("%.3fmts", airportResult.MinDistance),
			},
		},
	}

	json.NewEncoder(w).Encode(response)
}

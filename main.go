package main

import (
	"encoding/json"
	"fmt"
	"log"
	"neo-wifi-api/utils"
	"net/http"
	"os"
	"strconv"
)

type Coord = utils.Coord

var antennas []Coord
var cities []Coord
var airports []Coord

func loadData() {
	load := func(filePath string, target *[]Coord) {
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error leyendo %s: %v", filePath, err)
		}
		if err := json.Unmarshal(data, target); err != nil {
			log.Fatalf("Error parseando %s: %v", filePath, err)
		}
	}
	load("data/wifi-v15.json", &antennas)
	load("data/geodata-v4-mgc.json", &cities)
	load("data/airports.json", &airports)
}

func getClosestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	lat, _ := strconv.ParseFloat(query.Get("lat"), 64)
	lon, _ := strconv.ParseFloat(query.Get("lon"), 64)

	if lat == 0 || lon == 0 {
		http.Error(w, "Parámetros lat y lon requeridos", http.StatusBadRequest)
		return
	}

	current := Coord{Lat: lat, Lon: lon}

	closest, dist := utils.FindClosestN(current, antennas)

	resp := map[string]interface{}{
		"closest_wifi": map[string]interface{}{
			"name":     closest.Name,
			"distance": fmt.Sprintf("%.3fkm", dist),
			"coords": map[string]float64{
				"lat": closest.Lat,
				"lon": closest.Lon,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	loadData()

	http.HandleFunc("/api/closest", getClosestHandler)

	fmt.Println("Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

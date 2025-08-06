package types

// ClosestResult represents the result of finding closest targets
type ClosestResult struct {
	ClosestTarget       interface{} `json:"closestTarget"`
	SecondClosestTarget interface{} `json:"secondClosestTarget"`
	ThirdClosestTarget  interface{} `json:"thirdClosestTarget"`
	MinDistance         float64     `json:"minDistance"`
	SecondMinDistance   float64     `json:"secondMinDistance"`
	ThirdMinDistance    float64     `json:"thirdMinDistance"`
	Coords              Coords      `json:"coords"`
	SecondCoords        Coords      `json:"secondCoords"`
	ThirdCoords         Coords      `json:"thirdCoords"`
}

// SearchAntennaResult represents the result of searching for a specific antenna
type SearchAntennaResult struct {
	TargetDistance   float64 `json:"targetDistance"`
	SearchedTarget   string  `json:"searchedTarget"`
	SearchedTarget5g string  `json:"searchedTarget5g"`
	Coordinates      Coords  `json:"coordinates"`
	MAC              string  `json:"mac"`
	MAC5             string  `json:"mac5"`
	Type             string  `json:"type"`
}

// QueryResponse represents response when searching for specific antenna
type QueryResponse struct {
	Antenna  AntennaInfo `json:"antenna"`
	Distance string      `json:"distance"`
	Coords   string      `json:"coords"`
	MAC      string      `json:"MAC"`
	MAC5G    string      `json:"MAC5G"`
	Type     string      `json:"type"`
}

type AntennaInfo struct {
	Name   string `json:"name"`
	Name5g string `json:"name5g"`
}

// FullResponse represents the complete location response
type FullResponse struct {
	IP              string          `json:"ip"`
	City            string          `json:"city"`
	State           string          `json:"state"`
	Country         string          `json:"country"`
	Departament     string          `json:"departament"`
	CityCoords      CityCoords      `json:"city_coords"`
	CenterDistance  string          `json:"center_distance"`
	CurrentPosition CurrentPosition `json:"current_position"`
	ClosestWifi     WifiInfo        `json:"closest_wifi"`
	SecondWifi      WifiInfo        `json:"second_closest_wifi"`
	ThirdWifi       WifiInfo        `json:"third_closest_wifi"`
	AirportLocation AirportLocation `json:"airport_location"`
}

type CityCoords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CurrentPosition struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type WifiInfo struct {
	Antenna  string `json:"antenna"`
	Name     string `json:"name"`
	Distance string `json:"distance"`
	Type     string `json:"type"`
	MAC      string `json:"MAC"`
	MAC5G    string `json:"MAC5G"`
	Coords   string `json:"coords"`
	Users    int    `json:"users"`
}

type AirportLocation struct {
	City           string         `json:"city"`
	Country        string         `json:"country"`
	ClosestAirport ClosestAirport `json:"closest_airport"`
}

type ClosestAirport struct {
	Airport  string `json:"airport"`
	Distance string `json:"distance"`
}

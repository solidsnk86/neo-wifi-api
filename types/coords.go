package types

type Coords struct {
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

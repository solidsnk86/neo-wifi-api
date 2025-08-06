package types

type Antenna struct {
	Name   string  `json:"name"`
	Name5g string  `json:"name5g"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	MAC    string  `json:"MAC"`
	MAC5g  string  `json:"MAC5g"`
	Type   string  `json:"type"`
	Users  int     `json:"users"`
}

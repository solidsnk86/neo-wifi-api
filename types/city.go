package types

type City struct {
	Nombre       string  `json:"nombre"`
	Provincia    string  `json:"provincia"`
	Pais         string  `json:"pais"`
	Departamento string  `json:"departamento"`
	Lat          float64 `json:"lat"`
	Lon          float64 `json:"lon"`
}

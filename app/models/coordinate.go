package models

type OwnUnit struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type GenGeoJson struct {
	C240       Cat240s
	LatAwalAz  float64
	LonAwalAz  float64
	LatAkhirAz float64
	LonAkhirAz float64
	Ranges     float64
}

package models

type FeatureCollection struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
}

// Feature represents a feature in the FeatureCollection
type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

// Geometry represents the geometry of a feature
type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type Properties struct {
	Opacity float64 `json:"opacity"`
	Color   string  `json:"color"`
}

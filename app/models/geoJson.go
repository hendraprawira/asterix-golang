package models

type FeatureCollection struct {
	Type     string     `json:"type"`
	MsgIndex uint32			`json:"msgIndex"`
	EndAz    float64    `json:"endAzi"`
	StartAz  float64    `json:"startAzi"`
	Features []*Feature `json:"features"`
	Radius   float64    `json:"radius"`
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
	EndAz   float64 `json:"endAz"`
	// Radius  float64 `json:"radius"`
}

package model

type OSRMResponse struct {
	Code      string     `json:"code"`
	Routes    []Route    `json:"routes"`
	Waypoints []Waypoint `json:"waypoints"`
}

type Route struct {
	Distance float64  `json:"distance"`
	Duration float64  `json:"duration"`
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][]float64   `json:"coordinates"`
}

type Waypoint struct {
	Name     string    `json:"name"`
	Location []float64 `json:"location"`
	Distance float64   `json:"distance"`
}
package models

type Forecast struct {
	Max    float64 `json:"max"`
	Min    float64 `json:"min"`
	Sample int64   `json:"sample"`
}
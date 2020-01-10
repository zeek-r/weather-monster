package models

import (
	"github.com/jinzhu/gorm"
)

type Temperature struct {
	gorm.Model
	City      City `gorm:"foreignkey:CityID"`
	CityID    uint
	Max       int64
	Min       int64
	Timestamp int64
}

type Forecast struct {
	Max    float64 `json:"max"`
	Min    float64 `json:"min"`
	Sample int64   `json:"sample"`
}

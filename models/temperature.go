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

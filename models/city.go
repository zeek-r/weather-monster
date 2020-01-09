package models

import (
	"github.com/jinzhu/gorm"
)

type City struct {
	gorm.Model
	Name      string `gorm:"unique_index"`
	Latitude  string `gorm:"not null"`
	Longitude string `gorm:"not null"`
	Deleted   bool
}

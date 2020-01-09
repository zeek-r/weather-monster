package models

import (
	"github.com/jinzhu/gorm"
)

type Webhook struct {
	gorm.Model
	City        City `gorm:"foreignkey:CityID"`
	CityID      uint
	CallbackUrl string
}

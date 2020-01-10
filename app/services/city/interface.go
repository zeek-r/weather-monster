package cityservice

import (
	"github.com/zeek-r/weather-monster/models"
)

type Service interface {
	Create(toInsert *models.City) error
	Update(toUpdate *models.City) error
	DeleteByID(id uint) error
}

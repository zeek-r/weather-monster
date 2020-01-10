package temperaturerepo

import (
	"github.com/zeek-r/weather-monster/models"
)

type Repository interface {
	Create(toInsert *models.Temperature) error
	Update(toUpdate *models.Temperature) error
	DeleteByID(id int64) error
	GetByID(id int64) (*models.Temperature, error)
}

package cityrepo

import (
	"github.com/zeek-r/weather-monster/models"
)

type Repository interface {
	Create(toInsert *models.City) error
	Update(toUpdate *models.City) error
	DeleteByID(id int64) error
	GetByID(id int64) (*models.City, error)
}

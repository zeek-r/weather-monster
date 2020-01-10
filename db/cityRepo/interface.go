package cityrepo

import (
	"github.com/zeek-r/weather-monster/models"
)

type Repository interface {
	Create(toInsert *models.City) error
	Update(toUpdate *models.City) error
	DeleteByID(toDelete *models.City) error
}

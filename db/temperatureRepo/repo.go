package temperaturerepo

import (
	"github.com/jinzhu/gorm"
	"github.com/zeek-r/weather-monster/models"
)

type temperatureRepository struct {
	Conn *gorm.DB
}

func Init(Conn *gorm.DB) Repository {

	return &temperatureRepository{Conn}
}

func (repo *temperatureRepository) Create(toInsert *models.Temperature) error {
	created := repo.Conn.Create(toInsert)
	if created.Error != nil {
		return created.Error
	}
	return nil
}

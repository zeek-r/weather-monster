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

	return nil
}

func (repo *temperatureRepository) Update(toUpdate *models.Temperature) error {

	return nil
}

func (repo *temperatureRepository) DeleteByID(id int64) error {
	return nil
}

func (repo *temperatureRepository) GetByID(id int64) (*models.Temperature, error) {
	return nil, nil
}

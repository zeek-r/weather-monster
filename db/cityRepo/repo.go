package cityrepo

import (
	"github.com/jinzhu/gorm"
	"github.com/zeek-r/weather-monster/models"
)

type cityRepository struct {
	Conn *gorm.DB
}

func Init(Conn *gorm.DB) Repository {

	return &cityRepository{Conn}
}

func (repo *cityRepository) Create(toInsert *models.City) error {

	return nil
}

func (repo *cityRepository) Update(toUpdate *models.City) error {

	return nil
}

func (repo *cityRepository) DeleteByID(id int64) error {
	return nil
}

func (repo *cityRepository) GetByID(id int64) (*models.City, error) {
	return nil, nil
}

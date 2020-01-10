package cityservice

import (
	cityrepo "github.com/zeek-r/weather-monster/db/cityRepo"
	"github.com/zeek-r/weather-monster/models"
)

type CityService struct {
	cityRepo cityrepo.Repository
}

func NewCityService(city cityrepo.Repository) Service {
	return &CityService{city}
}

func (repo *CityService) Create(toInsert *models.City) error {

	return nil
}

func (repo *CityService) Update(toUpdate *models.City) error {

	return nil
}

func (repo *CityService) DeleteByID(id int64) error {
	return nil
}

func (repo *CityService) GetByID(id int64) (*models.City, error) {
	return nil, nil
}

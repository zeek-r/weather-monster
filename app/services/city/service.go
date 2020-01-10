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

func (service *CityService) Create(toInsert *models.City) error {
	return service.cityRepo.Create(toInsert)
}

func (service *CityService) Update(toUpdate *models.City) error {
	return service.cityRepo.Update(toUpdate)
}

func (service *CityService) DeleteByID(id uint) error {
	toDelete := new(models.City)
	toDelete.ID = id
	return service.cityRepo.DeleteByID(toDelete)
}

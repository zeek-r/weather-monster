package temperatureservice

import (
	temperaturerepo "github.com/zeek-r/weather-monster/db/temperatureRepo"
	"github.com/zeek-r/weather-monster/models"
)

type TemperatureService struct {
	temperatureRepo temperaturerepo.Repository
}

func NewTemperatureService(temperature temperaturerepo.Repository) Service {
	return &TemperatureService{temperature}
}

func (repo *TemperatureService) Create(toInsert *models.Temperature) error {

	return nil
}

func (repo *TemperatureService) Update(toUpdate *models.Temperature) error {

	return nil
}

func (repo *TemperatureService) DeleteByID(id int64) error {
	return nil
}

func (repo *TemperatureService) GetByID(id int64) (*models.Temperature, error) {
	return nil, nil
}

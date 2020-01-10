package forecastservice

import (
	temperatureRepo "github.com/zeek-r/weather-monster/db/temperatureRepo"
	"github.com/zeek-r/weather-monster/models"
)

type ForecastService struct {
	temperatureRepo temperatureRepo.Repository
}

func NewForecastService(temperature temperatureRepo.Repository) Service {
	return &ForecastService{temperature}
}

func (repo *ForecastService) GetByID(id int64) (*models.Temperature, error) {
	return nil, nil
}

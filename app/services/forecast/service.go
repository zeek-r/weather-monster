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

func (service *ForecastService) GetByID(id uint) (models.Forecast, error) {
	return service.temperatureRepo.GetForecastByID(id)
}

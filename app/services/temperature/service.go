package temperatureservice

import (
	temperaturerepo "github.com/zeek-r/weather-monster/db/temperatureRepo"
	"github.com/zeek-r/weather-monster/models"
	"time"
)

type TemperatureService struct {
	temperatureRepo temperaturerepo.Repository
}

func NewTemperatureService(temperature temperaturerepo.Repository) Service {
	return &TemperatureService{temperature}
}

func (service *TemperatureService) Create(toInsert *models.Temperature) error {
	toInsert.Timestamp = time.Now().Unix()
	return service.temperatureRepo.Create(toInsert)
}

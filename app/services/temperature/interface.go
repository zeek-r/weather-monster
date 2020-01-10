package temperatureservice

import (
	"github.com/zeek-r/weather-monster/models"
)

type Service interface {
	Create(toInsert *models.Temperature) error
}

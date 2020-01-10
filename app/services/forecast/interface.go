package forecastservice

import (
	"github.com/zeek-r/weather-monster/models"
)

type Service interface {
	GetByID(id uint) (models.Forecast, error)
}

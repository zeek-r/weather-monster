package temperaturerepo

import (
	"github.com/zeek-r/weather-monster/models"
)

type Repository interface {
	Create(toInsert *models.Temperature) error
	GetForecastByID(id uint) (*models.Forecast, error)
}

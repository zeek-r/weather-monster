package forecastservice

import (
	"github.com/zeek-r/weather-monster/models"
)

type Service interface {
	GetByID(id int64) (*models.Temperature, error)
}

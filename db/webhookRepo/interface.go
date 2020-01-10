package webhookrepo

import (
	"github.com/zeek-r/weather-monster/models"
)

type Repository interface {
	Create(toInsert *models.Webhook) error
	DeleteByID(toDelete *models.Webhook) error
}

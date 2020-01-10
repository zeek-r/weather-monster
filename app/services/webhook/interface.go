package webhookservice

import (
	"github.com/zeek-r/weather-monster/models"
)

type Service interface {
	Create(toInsert *models.Webhook) error
	DeleteByID(id uint) error
}

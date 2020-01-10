package webhookservice

import (
	"github.com/zeek-r/weather-monster/models"
)

type Service interface {
	Create(toInsert *models.Webhook) error
	Update(toUpdate *models.Webhook) error
	DeleteByID(id int64) error
	GetByID(id int64) (*models.Webhook, error)
}

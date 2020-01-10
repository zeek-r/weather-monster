package webhookservice

import (
	webhookrepo "github.com/zeek-r/weather-monster/db/webhookRepo"
	"github.com/zeek-r/weather-monster/models"
)

type WebhookService struct {
	webhookRepo webhookrepo.Repository
}

func NewWebhookService(webhook webhookrepo.Repository) Service {
	return &WebhookService{webhook}
}

func (service *WebhookService) Create(toInsert *models.Webhook) error {
	return service.webhookRepo.Create(toInsert)
}

func (service *WebhookService) DeleteByID(id uint) error {
	toDelete := new(models.Webhook)
	toDelete.ID = id
	return service.webhookRepo.DeleteByID(toDelete)
}

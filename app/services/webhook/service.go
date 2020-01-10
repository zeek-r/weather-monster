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

func (repo *WebhookService) Create(toInsert *models.Webhook) error {

	return nil
}

func (repo *WebhookService) Update(toUpdate *models.Webhook) error {

	return nil
}

func (repo *WebhookService) DeleteByID(id int64) error {
	return nil
}

func (repo *WebhookService) GetByID(id int64) (*models.Webhook, error) {
	return nil, nil
}

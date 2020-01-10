package webhookhandler

import (
	"net/http"

	webhook "github.com/zeek-r/weather-monster/app/services/webhook"
)

type WebhookHandler struct {
	webhookService *webhook.Service
}

func NewWebhookHandler(service *webhook.Service) *WebhookHandler {
	return &WebhookHandler{service}
}

func (webhookService *WebhookHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (webhookService *WebhookHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (webhookService *WebhookHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {

}

func (webhookService *WebhookHandler) GetByID(w http.ResponseWriter, r *http.Request) {

}

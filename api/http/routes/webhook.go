package routes

import (
	"github.com/gorilla/mux"
	webhookhandler "github.com/zeek-r/weather-monster/api/http/handlers/webhook"
)

func LoadWebhookRoutes(router *mux.Router, webhookHandler *webhookhandler.WebhookHandler) {
	router.HandleFunc("/webhooks", webhookHandler.Create).Methods("POST")
	router.HandleFunc("/webhooks/{id}", webhookHandler.DeleteByID).Methods("DELETE")
}

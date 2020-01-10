package webhookhandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zeek-r/weather-monster/api/http/utils/response"
	webhook "github.com/zeek-r/weather-monster/app/services/webhook"
	"github.com/zeek-r/weather-monster/models"
	"github.com/zeek-r/weather-monster/pkg/logger"
)

type WebhookHandler struct {
	webhookService webhook.Service
}

func NewWebhookHandler(service webhook.Service) *WebhookHandler {
	return &WebhookHandler{service}
}

func (handler *WebhookHandler) Create(w http.ResponseWriter, r *http.Request) {

	var request models.Webhook

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		message := "Error reading body"
		logger.Error(err, message)
		response.Err(w, message, 500)
		return
	}

	if err := json.Unmarshal(body, &request); err != nil {
		message := "Error unmarshalling webhook"
		logger.Error(err, message)
		response.Err(w, message, 500)
		return
	}

	logger.Info("Creating Webhook record", map[string]interface{}{"cityId": request.CityID})

	if err := handler.webhookService.Create(&request); err != nil {
		message := "Error Creating Webhook record"
		logger.Error(err, message)
		response.Err(w, message, 400)
		return
	}

	response.Json(w, map[string]interface{}{"success": true, "message": "Webhook created successfully"})
	return
}

func (handler *WebhookHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	webhookID := params["id"]
	ID, _ := strconv.ParseUint(webhookID, 10, 64)

	if err := handler.webhookService.DeleteByID(uint(ID)); err != nil {
		message := "Error deleting webhook"
		logger.Error(err, message)
		response.Err(w, message, 400)
		return
	}
	response.Json(w, map[string]interface{}{"success": true, "message": "Webhook Deleted successfully"})
	return
}

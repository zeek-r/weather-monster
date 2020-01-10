package forecasthandler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zeek-r/weather-monster/api/http/utils/response"
	forecast "github.com/zeek-r/weather-monster/app/services/forecast"
	"github.com/zeek-r/weather-monster/models"
	"github.com/zeek-r/weather-monster/pkg/logger"
)

type ForecastHandler struct {
	forecastService forecast.Service
}

func NewForecastHandler(service forecast.Service) *ForecastHandler {
	return &ForecastHandler{service}
}

func (handler *ForecastHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	var forecast models.Forecast
	var err error
	params := mux.Vars(r)

	webhookID := params["city_id"]
	ID, _ := strconv.ParseUint(webhookID, 10, 64)

	if forecast, err = handler.forecastService.GetByID(uint(ID)); err != nil {
		message := "Error getting forecast"
		logger.Error(err, message)
		response.Err(w, message, 400)
		return
	}
	response.Json(w, map[string]interface{}{"success": true, "forecast": forecast})
	return
}

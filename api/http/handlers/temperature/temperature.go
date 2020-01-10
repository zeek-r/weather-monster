package temperaturehandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zeek-r/weather-monster/api/http/utils/response"
	temperature "github.com/zeek-r/weather-monster/app/services/temperature"
	"github.com/zeek-r/weather-monster/models"
	"github.com/zeek-r/weather-monster/pkg/logger"
)

type TemperatureHandler struct {
	temperatureService temperature.Service
}

func NewTemperatureHandler(service temperature.Service) *TemperatureHandler {
	return &TemperatureHandler{service}
}

func (handler *TemperatureHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request models.Temperature

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		message := "Error reading body"
		logger.Error(err, message)
		response.Err(w, message, 500)
		return
	}

	if err := json.Unmarshal(body, &request); err != nil {
		message := "Error unmarshalling city"
		logger.Error(err, message)
		response.Err(w, message, 500)
		return
	}

	logger.Info("Creating Temperature record", map[string]interface{}{"cityId": request.CityID})

	if err := handler.temperatureService.Create(&request); err != nil {
		message := "Error Creating temperature record"
		logger.Error(err, message)
		response.Err(w, message, 400)
		return
	}

	response.Json(w, map[string]interface{}{"success": true, "message": "Temperature created successfully"})
	return
}

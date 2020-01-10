package temperaturehandler

import (
	"net/http"

	temperature "github.com/zeek-r/weather-monster/app/services/temperature"
)

type TemperatureHandler struct {
	temperatureService *temperature.Service
}

func NewTemperatureHandler(service *temperature.Service) *TemperatureHandler {
	return &TemperatureHandler{service}
}

func (temperatureService *TemperatureHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (temperatureService *TemperatureHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (temperatureService *TemperatureHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {

}

func (temperatureService *TemperatureHandler) GetByID(w http.ResponseWriter, r *http.Request) {

}

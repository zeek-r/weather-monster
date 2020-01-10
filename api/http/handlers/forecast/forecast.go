package forecasthandler

import (
	"net/http"

	forecast "github.com/zeek-r/weather-monster/app/services/forecast"
)

type ForecastHandler struct {
	forecastService *forecast.Service
}

func NewForecastHandler(service *forecast.Service) *ForecastHandler {
	return &ForecastHandler{service}
}

func (forecastService *ForecastHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (forecastService *ForecastHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (forecastService *ForecastHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {

}

func (forecastService *ForecastHandler) GetByID(w http.ResponseWriter, r *http.Request) {

}

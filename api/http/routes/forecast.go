package routes

import (
	"github.com/gorilla/mux"
	handler "github.com/zeek-r/weather-monster/api/http/handlers/forecast"
)

func LoadForecastRoutes(router *mux.Router, forecastHandler *handler.ForecastHandler) {
	router.HandleFunc("/forecasts/{city_id}", forecastHandler.GetByID).Methods("GET")
}

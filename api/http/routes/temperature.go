package routes

import (
	"github.com/gorilla/mux"
	temperaturehandler "github.com/zeek-r/weather-monster/api/http/handlers/temperature"
)

func LoadTemperatureRoutes(router *mux.Router, temperatureHandler *temperaturehandler.TemperatureHandler) {
	router.HandleFunc("/temperatures", temperatureHandler.Create).Methods("POST")
}

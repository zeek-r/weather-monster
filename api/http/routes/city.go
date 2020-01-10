package routes

import (
	"github.com/gorilla/mux"
	handler "github.com/zeek-r/weather-monster/api/http/handlers/city"
)

func LoadCityRoutes(router *mux.Router, cityHandler *handler.CityHandler) {
	router.HandleFunc("/cities", cityHandler.Create).Methods("POST")
	router.HandleFunc("/cities/{id}", cityHandler.Update).Methods("PATCH")
	router.HandleFunc("/cities/{id}", cityHandler.DeleteByID).Methods("DELETE")
}

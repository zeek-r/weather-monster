package cityhandler

import (
	"net/http"

	cityservice "github.com/zeek-r/weather-monster/app/services/city"
)

type CityHandler struct {
	cityService *cityservice.Service
}

func NewCityHandler(service *cityservice.Service) *CityHandler {
	return &CityHandler{service}
}

func (cityService *CityHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (cityService *CityHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (cityService *CityHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {

}

func (cityService *CityHandler) GetByID(w http.ResponseWriter, r *http.Request) {

}

package cityhandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeek-r/weather-monster/api/http/utils/response"
	cityservice "github.com/zeek-r/weather-monster/app/services/city"
	cityModel "github.com/zeek-r/weather-monster/models"
	"github.com/zeek-r/weather-monster/pkg/logger"
	"strconv"
)

type CityHandler struct {
	cityService cityservice.Service
}

func NewCityHandler(service cityservice.Service) *CityHandler {
	return &CityHandler{service}
}

func (handler *CityHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request cityModel.City

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

	logger.Info("Creating city", map[string]interface{}{"name": request.Name})

	if err := handler.cityService.Create(&request); err != nil {
		message := "Error Creating city"
		logger.Error(err, message)
		response.Err(w, message, 400)
		return
	}

	response.Json(w, map[string]interface{}{"success": true, "message": "City created successfully"})
	return
}

func (handler *CityHandler) Update(w http.ResponseWriter, r *http.Request) {
	var request cityModel.City

	params := mux.Vars(r)

	cityID := params["id"]
	ID, _ := strconv.ParseUint(cityID, 10, 64)
	request.ID = uint(ID)

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

	logger.Info("Updating city", map[string]interface{}{"name": request.Name})

	if err := handler.cityService.Update(&request); err != nil {
		message := "Error updating city"
		logger.Error(err, message)
		response.Err(w, message, 400)
		return
	}

	response.Json(w, map[string]interface{}{"success": true, "message": "City Updated successfully"})
	return
}

func (handler *CityHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	cityID := params["id"]
	ID, _ := strconv.ParseUint(cityID, 10, 64)

	if err := handler.cityService.DeleteByID(uint(ID)); err != nil {
		message := "Error deleting city"
		logger.Error(err, message)
		response.Err(w, message, 400)
		return
	}
	response.Json(w, map[string]interface{}{"success": true, "message": "City Deleted successfully"})
	return
}

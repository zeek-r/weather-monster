package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ResData struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message"`
}

func Json(w http.ResponseWriter, result interface{}, status ...http.ConnState) {
	// Status
	var resStatus http.ConnState
	if status != nil {
		resStatus = status[0]
	} else {
		resStatus = http.StatusOK // default 200
	}
	results := fmt.Sprintf("%v", result)
	if results == "[]" {
		result = []string{}
	}
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(resStatus))
	w.Write(js)
	return
}

func Err(w http.ResponseWriter, message string, status ...http.ConnState) {
	var resStatus http.ConnState
	if status != nil {
		resStatus = status[0]
	} else if message == "not-found" {
		resStatus = http.StatusNotFound // default 404
	} else {
		resStatus = http.StatusBadRequest // default 400
	}

	// response data
	resData := &ResData{}
	resData.Status = strconv.Itoa(int(resStatus))
	resData.Message = message
	js, err := json.Marshal(resData)
	if err != nil {
		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// return
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(resStatus))
	w.Write(js)
}

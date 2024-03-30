package middlewares

import (
	"encoding/json"
	"net/http"

	"GO_PROJECT/logger"
	"GO_PROJECT/model"
)

func HandleResponse(w http.ResponseWriter, statusCode int, message string, data interface{}, err error) {
	var apiResponse model.ApiResponse
	if err != nil {
		logger.Log.Error(message+": ", err)
		apiResponse = model.ApiResponse{StatusCode: statusCode, Message: message, Error: err.Error()}

	} else {
		logger.Log.Info(message)
		apiResponse = model.ApiResponse{StatusCode: statusCode, Message: message, Data: data}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(apiResponse)
}

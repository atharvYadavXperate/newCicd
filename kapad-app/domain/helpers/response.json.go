package helpers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error.omitempty"`
}

func ResponseJsonWriter(w http.ResponseWriter, data interface{}, statusCode int, message string, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(response{
		Message: message,
		Data:    data,
		Error:   errorMessage,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func ResponseWriterWithError(w http.ResponseWriter, statusCode int, message string, errorMessage string) {
	ResponseJsonWriter(w, nil, statusCode, message, errorMessage)
}

func ResponseWriterWithInternalServerError(w http.ResponseWriter, message string, errorMessage string) {
	ResponseWriterWithError(w, http.StatusInternalServerError, message, errorMessage)
}

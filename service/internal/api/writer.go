package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// WriteJSON will write a JSON Object to the given HTTP Response Writer
func WriteJSON(w http.ResponseWriter, response interface{}) {
	WriteJSONWithStatus(w, http.StatusOK, response)
}

// WriteJSONWithStatus will write a JSON Object to the given HTTP Response Writer
func WriteJSONWithStatus(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logrus.WithError(err).Error("Failed to send response to client")
	}
}

package problems

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Write will write an RFC-7807 Problem to the given HTTP Response Writer
func Write(w http.ResponseWriter, problem ProblemResponse) {
	w.Header().Set("Content-Type", problem.MediaType())
	w.WriteHeader(problem.StatusCode())
	err := json.NewEncoder(w).Encode(problem)
	if err != nil {
		logrus.WithError(err).Error("Failed to send response to client")
	}
}

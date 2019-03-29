package health

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

// Healthchecker is a means to check the health of the system
type Healthchecker struct {
	checks map[string]Healthcheck
}

// New creates a new Healthchecker
func New() Healthchecker {
	return Healthchecker{
		checks: make(map[string]Healthcheck),
	}
}

// AddHealthcheck will add a healthcheck to the system
func (h *Healthchecker) AddHealthcheck(name string, check Healthcheck) {
	h.checks[name] = check
}

// Health represents the health of the system
type Health struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// CheckHealth will check the health of the system
func (h *Healthchecker) checkHealth(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]Health)

	statusCode := http.StatusOK

	for name, checker := range h.checks {
		checkResult := checker.CheckHealth()

		if checkResult == nil {
			result[name] = Health{Status: "ok"}
		} else {
			statusCode = http.StatusServiceUnavailable
			result[name] = Health{Status: "fail", Message: checkResult.Error()}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		logrus.WithError(err).Error("Failed to write response")
	}
}

// NewRouter will return the router used for performing Healthchecks
func NewRouter(h *Healthchecker) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", h.checkHealth)
	return r
}

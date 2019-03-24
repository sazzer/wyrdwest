package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sazzer/wyrdwest/service/internal/server"
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
func (h *Healthchecker) checkHealth(c echo.Context) error {
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

	return c.JSON(statusCode, result)
}

// RegisterHandler will return the means to register the Healthcheck Handler with the HTTP Server
func RegisterHandler(h *Healthchecker) server.HandlerRegistrationFunc {
	return func(e *echo.Echo) {
		e.GET("/health", h.checkHealth)
	}
}

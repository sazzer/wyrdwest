package validation

import (
	"net/http"

	"github.com/sazzer/wyrdwest/service/internal/api/problems"
)

// Error represents a single error with a single field to a request
type Error struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// Problem is a Problem response for request validation failures
type Problem struct {
	problems.Problem
	Errors []Error `json:"validationErrors"`
}

// New creates a new Validation Problem for the provided errors
func New(errors []Error) Problem {
	return Problem{
		problems.Problem{
			Type:   "tag:wyrdwest,2019:problems/validation-problem",
			Title:  "Validation Falure",
			Status: http.StatusBadRequest,
		},
		errors,
	}
}

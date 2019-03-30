package problems

import (
	"net/http"
)

// ProblemResponse represents something that can be returned as a problem from the server
type ProblemResponse interface {
	// StatusCode represents the Status Code to return
	StatusCode() int
	// MediaType represents the media type to return
	MediaType() string
}

// Problem represents an RFC-7807 Problem response
type Problem struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

// StatusCode represents the Status Code to return
func (p Problem) StatusCode() int {
	statusCode := http.StatusInternalServerError
	if p.Status != 0 {
		statusCode = p.Status
	}

	return statusCode
}

// MediaType represents the media type to return
func (p Problem) MediaType() string {
	return "application/problem+json"
}

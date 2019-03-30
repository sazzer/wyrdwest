package problems_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sazzer/wyrdwest/service/internal/api/problems"
)

func TestWriteProblem(t *testing.T) {
	problem := problems.Problem{
		Type:   "The Type",
		Title:  "The Title",
		Status: http.StatusBadRequest,
	}

	rec := httptest.NewRecorder()

	problems.Write(rec, problem)

	g := NewGomegaWithT(t)
	g.Expect(rec.Code).To(Equal(400))
	g.Expect(rec.Header().Get("Content-Type")).To(Equal("application/problem+json"))
	g.Expect(rec.Body).To(MatchJSON(`{
		"type": "The Type",
		"title": "The Title",
		"status": 400
	}`))
}

func TestWriteProblemNoStatusCode(t *testing.T) {
	problem := problems.Problem{
		Type:  "The Type",
		Title: "The Title",
	}

	rec := httptest.NewRecorder()

	problems.Write(rec, problem)

	g := NewGomegaWithT(t)
	g.Expect(rec.Code).To(Equal(500))
	g.Expect(rec.Header().Get("Content-Type")).To(Equal("application/problem+json"))
	g.Expect(rec.Body).To(MatchJSON(`{
		"type": "The Type",
		"title": "The Title",
		"status": 0
	}`))
}

func TestWriteProblemChildType(t *testing.T) {
	type DetailedProblem struct {
		problems.Problem
		Name string `json:"username"`
	}
	problem := DetailedProblem{
		problems.Problem{
			Type:   "The Type",
			Title:  "The Title",
			Status: http.StatusBadRequest,
		},
		"Graham",
	}

	rec := httptest.NewRecorder()

	problems.Write(rec, problem)

	g := NewGomegaWithT(t)
	g.Expect(rec.Code).To(Equal(400))
	g.Expect(rec.Header().Get("Content-Type")).To(Equal("application/problem+json"))
	g.Expect(rec.Body).To(MatchJSON(`{
		"type": "The Type",
		"title": "The Title",
		"status": 400,
		"username": "Graham"
	}`))
}

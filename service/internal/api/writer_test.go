package api_test

import (
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sazzer/wyrdwest/service/internal/api"
)

func TestWriteMap(t *testing.T) {
	result := map[string]interface{}{
		"foo":    "bar",
		"answer": 42,
	}

	rec := httptest.NewRecorder()

	api.WriteJSON(rec, result)

	g := NewGomegaWithT(t)
	g.Expect(rec.Code).To(Equal(200))
	g.Expect(rec.Header().Get("Content-Type")).To(Equal("application/json"))
	g.Expect(rec.Body).To(MatchJSON(`{
		"answer": 42,
		"foo": "bar"
	}`))
}

func TestWriteString(t *testing.T) {
	result := "Hello"

	rec := httptest.NewRecorder()

	api.WriteJSON(rec, result)

	g := NewGomegaWithT(t)
	g.Expect(rec.Code).To(Equal(200))
	g.Expect(rec.Header().Get("Content-Type")).To(Equal("application/json"))
	g.Expect(rec.Body).To(MatchJSON(`"Hello"`))
}

func TestWriteNumber(t *testing.T) {
	result := 42

	rec := httptest.NewRecorder()

	api.WriteJSON(rec, result)

	g := NewGomegaWithT(t)
	g.Expect(rec.Code).To(Equal(200))
	g.Expect(rec.Header().Get("Content-Type")).To(Equal("application/json"))
	g.Expect(rec.Body).To(MatchJSON(`42`))
}

func TestWriteStruct(t *testing.T) {
	type Response struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	result := Response{
		Name: "Ted",
		Age:  30,
	}

	rec := httptest.NewRecorder()

	api.WriteJSON(rec, result)

	g := NewGomegaWithT(t)
	g.Expect(rec.Code).To(Equal(200))
	g.Expect(rec.Header().Get("Content-Type")).To(Equal("application/json"))
	g.Expect(rec.Body).To(MatchJSON(`{
		"name": "Ted",
		"age": 30
	}`))
}

func TestWriteStringWithStatusCode(t *testing.T) {
	result := "Hello"

	rec := httptest.NewRecorder()

	api.WriteJSONWithStatus(rec, 203, result)

	g := NewGomegaWithT(t)
	g.Expect(rec.Code).To(Equal(203))
	g.Expect(rec.Header().Get("Content-Type")).To(Equal("application/json"))
	g.Expect(rec.Body).To(MatchJSON(`"Hello"`))
}

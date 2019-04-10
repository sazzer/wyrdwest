package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
)

// NewRouter will return the router used for working with Attributes
func NewRouter(dao attributes.Retriever) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { list(w, r, dao) })
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) { getByID(w, r, dao) })
	return r
}

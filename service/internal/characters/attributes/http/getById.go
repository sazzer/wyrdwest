package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sazzer/wyrdwest/service/internal/problems"

	uuid "github.com/satori/go.uuid"

	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
)

func getByID(w http.ResponseWriter, r *http.Request, retriever attributes.Retriever) {
	idVal := chi.URLParam(r, "id")
	parsedID, err := uuid.FromString(idVal)
	if err != nil {
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(problems.Problem{
			Type:   "tag:wyrdwest,2019:problems/attributes/invalid-id",
			Title:  "The Attribute ID was invalid",
			Status: http.StatusBadRequest,
		})
		return
	}

	attribute, err := retriever.GetAttributeByID(attributes.AttributeID(parsedID))
	if err != nil {
		switch err.(type) {
		case attributes.AttributeNotFoundError:
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(problems.Problem{
				Type:   "tag:wyrdwest,2019:problems/attributes/unknown-attribute",
				Title:  "The Attribute was not found",
				Status: http.StatusNotFound,
			})
		default:
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(problems.Problem{
				Type:   "tag:wyrdwest,2019:problems/internal-server-error",
				Title:  "An unexpected error occurred",
				Status: http.StatusInternalServerError,
			})
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Attribute{
		Self:        fmt.Sprintf("/attributes/%s", idVal),
		Name:        attribute.Name,
		Description: attribute.Description,
	})
}

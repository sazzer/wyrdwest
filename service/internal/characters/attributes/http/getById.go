package http

import (
	"fmt"
	"net/http"

	"github.com/sazzer/wyrdwest/service/internal/problems"

	uuid "github.com/satori/go.uuid"

	"github.com/labstack/echo/v4"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
)

func getByID(c echo.Context, retriever attributes.Retriever) error {
	idVal := c.Param("id")
	parsedID, err := uuid.FromString(idVal)
	if err != nil {
		return c.JSON(http.StatusBadRequest, problems.Problem{
			Type:   "tag:wyrdwest,2019:problems/attributes/invalid-id",
			Title:  "The Attribute ID was invalid",
			Status: http.StatusBadRequest,
		})
	}

	attribute, err := retriever.GetAttributeByID(attributes.AttributeID(parsedID))
	if err != nil {
		switch err.(type) {
		case attributes.AttributeNotFoundError:
			return c.JSON(http.StatusNotFound, problems.Problem{
				Type:   "tag:wyrdwest,2019:problems/attributes/unknown-attribute",
				Title:  "The Attribute was not found",
				Status: http.StatusNotFound,
			})
		default:
			return c.JSON(http.StatusInternalServerError, problems.Problem{
				Type:   "tag:wyrdwest,2019:problems/internal-server-error",
				Title:  "An unexpected error occurred",
				Status: http.StatusInternalServerError,
			})
		}
	}

	return c.JSON(http.StatusOK, Attribute{
		Self:        fmt.Sprintf("/attributes/%s", idVal),
		Name:        attribute.Name,
		Description: attribute.Description,
	})
}

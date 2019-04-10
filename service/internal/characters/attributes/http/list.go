package http

import (
	"net/http"
	"strconv"

	"github.com/sazzer/wyrdwest/service/internal/api"

	"github.com/sazzer/wyrdwest/service/internal/api/problems"

	"github.com/sirupsen/logrus"

	"github.com/sazzer/wyrdwest/service/internal/api/validation"
	"github.com/sazzer/wyrdwest/service/internal/service"

	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
)

func parseInt(s string, dest *int) error {
	if s != "" {
		n, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		*dest = n
	}
	return nil
}

type listParams struct {
	offset     int
	count      int
	sorts      []service.SortField
	nameFilter string
}

func (params *listParams) parse(r *http.Request) []validation.Error {
	validationErrors := []validation.Error{}

	if err := parseInt(r.URL.Query().Get("offset"), &params.offset); err != nil {
		logrus.WithField("url", r.URL).WithError(err).Error("Failed to parse offset")
		validationErrors = append(validationErrors, validation.Error{
			Field: "offset",
			Error: "tag:wyrdwest,2019:validation-errors/invalid-number",
		})
	} else if params.offset < 0 {
		validationErrors = append(validationErrors, validation.Error{
			Field: "offset",
			Error: "tag:wyrdwest,2019:validation-errors/negative-number",
		})
	}

	if err := parseInt(r.URL.Query().Get("count"), &params.count); err != nil {
		logrus.WithField("url", r.URL).WithError(err).Error("Failed to parse count")
		validationErrors = append(validationErrors, validation.Error{
			Field: "count",
			Error: "tag:wyrdwest,2019:validation-errors/invalid-number",
		})
	} else if params.count < 0 {
		validationErrors = append(validationErrors, validation.Error{
			Field: "count",
			Error: "tag:wyrdwest,2019:validation-errors/negative-number",
		})
	}

	params.sorts = service.ParseSorts(r.URL.Query().Get("sort"))
	params.nameFilter = r.URL.Query().Get("name")

	return validationErrors
}

func list(w http.ResponseWriter, r *http.Request, retriever attributes.Retriever) {
	// Parse the inputs
	params := listParams{
		offset:     0,
		count:      10,
		sorts:      []service.SortField{},
		nameFilter: "",
	}
	validationErrors := params.parse(r)

	if len(validationErrors) > 0 {
		problems.Write(w, validation.New(validationErrors))
		return
	}

	offset := uint64(params.offset)
	count := uint64(params.count)

	// Load the data
	attributesData, err := retriever.ListAttributes(attributes.AttributeMatchCriteria{
		Name: params.nameFilter,
	}, params.sorts, offset, count)

	// Handle the error
	if err != nil {
		problems.Write(w, problems.Problem{
			Type:   "tag:wyrdwest,2019:problems/internal-server-error",
			Title:  "An unexpected error occurred",
			Status: http.StatusInternalServerError,
		})
		return
	}
	// Build the response
	results := []Attribute{}
	for _, attribute := range attributesData.Data {
		results = append(results, buildAttribute(attribute))
	}

	api.WriteJSON(w, Attributes{
		Self:   "/attributes",
		Offset: offset,
		Total:  attributesData.PageInfo.TotalSize,
		Data:   results,
	})
}

package http_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
	attributesHttp "github.com/sazzer/wyrdwest/service/internal/characters/attributes/http"
	"github.com/sazzer/wyrdwest/service/internal/service"
)

func (suite HTTPSuite) testList(params string) *httptest.ResponseRecorder {
	r := attributesHttp.NewRouter(suite.mockRetriever)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?%s", params), nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	return rec
}

func (suite *HTTPSuite) TestValidationErrors() {
	type test struct {
		url    string
		errors string
	}

	tests := []test{
		{url: "offset=-1", errors: `{"field": "offset", "error": "tag:wyrdwest,2019:validation-errors/negative-number"}`},
		{url: "count=-1", errors: `{"field": "count", "error": "tag:wyrdwest,2019:validation-errors/negative-number"}`},
		{url: "count=-1&offset=-1", errors: `{"field": "offset", "error": "tag:wyrdwest,2019:validation-errors/negative-number"},{"field": "count", "error": "tag:wyrdwest,2019:validation-errors/negative-number"}`},

		{url: "offset=a", errors: `{"field": "offset", "error": "tag:wyrdwest,2019:validation-errors/invalid-number"}`},
		{url: "count=a", errors: `{"field": "count", "error": "tag:wyrdwest,2019:validation-errors/invalid-number"}`},
		{url: "count=a&offset=a", errors: `{"field": "offset", "error": "tag:wyrdwest,2019:validation-errors/invalid-number"},{"field": "count", "error": "tag:wyrdwest,2019:validation-errors/invalid-number"}`},

		{url: "count=-1&offset=a", errors: `{"field": "offset", "error": "tag:wyrdwest,2019:validation-errors/invalid-number"},{"field": "count", "error": "tag:wyrdwest,2019:validation-errors/negative-number"}`},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.url, func() {
			rec := suite.testList(test.url)

			expectedResponse := fmt.Sprintf(`{
				"type": "tag:wyrdwest,2019:problems/validation-problem",
				"title": "Validation Falure",
				"status": 400,
				"validationErrors": [%s]
			}`, test.errors)
			g := NewGomegaWithT(suite.T())
			g.Expect(rec.Code).To(Equal(400))
			g.Expect(rec.Body).To(MatchJSON(expectedResponse))
		})
	}
}

func (suite *HTTPSuite) TestNoParamsNoResponse() {
	response := attributes.AttributePage{}
	response.TotalSize = 0

	suite.mockRetriever.EXPECT().
		ListAttributes(attributes.AttributeMatchCriteria{}, []service.SortField{}, uint64(0), uint64(10)).
		Return(response, nil).
		Times(1)

	rec := suite.testList("")

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(200))
	g.Expect(rec.Body).To(MatchJSON(`{
		"self": "/attributes",
		"offset": 0,
		"total": 0,
		"data": []
	}`))
}

func (suite *HTTPSuite) TestParamsNoResponse() {
	response := attributes.AttributePage{}
	response.TotalSize = 0

	suite.mockRetriever.EXPECT().
		ListAttributes(attributes.AttributeMatchCriteria{Name: "Strength"},
			[]service.SortField{{Field: "name", Direction: service.SortDescending}},
			uint64(10),
			uint64(5)).
		Return(response, nil).
		Times(1)

	rec := suite.testList("name=Strength&sort=-name&offset=10&count=5")

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(200))
	g.Expect(rec.Body).To(MatchJSON(`{
		"self": "/attributes",
		"offset": 10,
		"total": 0,
		"data": []
	}`))
}

func (suite *HTTPSuite) TestNoParamsResponse() {
	id1 := attributes.AttributeID("00000000-0000-0000-0000-000000000000")
	id2 := attributes.AttributeID("00000000-0000-0000-0000-000000000001")

	response := attributes.AttributePage{
		Data: []attributes.Attribute{
			{
				ID:          id1,
				Version:     uuid.NewV4().String(),
				Created:     time.Now(),
				Updated:     time.Now(),
				Name:        "Strength",
				Description: "How Strong I am",
			},
			{
				ID:          id2,
				Version:     uuid.NewV4().String(),
				Created:     time.Now(),
				Updated:     time.Now(),
				Name:        "Intelligence",
				Description: "How Smart I am",
			},
		},
	}
	response.TotalSize = 2

	suite.mockRetriever.EXPECT().
		ListAttributes(attributes.AttributeMatchCriteria{}, []service.SortField{}, uint64(0), uint64(10)).
		Return(response, nil).
		Times(1)

	rec := suite.testList("")

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(200))
	g.Expect(rec.Body).To(MatchJSON(`{
		"self": "/attributes",
		"offset": 0,
		"total": 2,
		"data": [
			{
			  "self": "/attributes/00000000-0000-0000-0000-000000000000",
			  "name": "Strength",
			  "description": "How Strong I am"
			},
			{
			  "self": "/attributes/00000000-0000-0000-0000-000000000001",
			  "name": "Intelligence",
			  "description": "How Smart I am"
			}
		]
	}`))
}

func (suite *HTTPSuite) TestServiceError() {
	suite.mockRetriever.EXPECT().
		ListAttributes(attributes.AttributeMatchCriteria{}, []service.SortField{}, uint64(0), uint64(10)).
		Return(attributes.AttributePage{}, errors.New("Oops")).
		Times(1)

	rec := suite.testList("")

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(500))
	g.Expect(rec.Body).To(MatchJSON(`{
		"type": "tag:wyrdwest,2019:problems/internal-server-error",
		"title": "An unexpected error occurred",
		"status": 500
	}`))
}

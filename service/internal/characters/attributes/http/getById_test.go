package http_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"

	. "github.com/onsi/gomega"
	attributesHttp "github.com/sazzer/wyrdwest/service/internal/characters/attributes/http"
)

func (suite HTTPSuite) testGetByID(id string) *httptest.ResponseRecorder {
	r := attributesHttp.NewRouter(suite.mockRetriever)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", id), nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	return rec
}

func (suite *HTTPSuite) TestGetUnknownAttributeByID() {
	suite.mockRetriever.EXPECT().
		GetAttributeByID(attributes.AttributeID(uuid.Must(uuid.FromString("00000000-0000-0000-0000-000000000000")))).
		Return(attributes.Attribute{}, attributes.AttributeNotFoundError{}).
		Times(1)

	rec := suite.testGetByID("00000000-0000-0000-0000-000000000000")

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(404))
	g.Expect(rec.Body).To(MatchJSON(`{
		"type": "tag:wyrdwest,2019:problems/attributes/unknown-attribute",
		"title": "The Attribute was not found",
		"status": 404
	}`))
}

func (suite *HTTPSuite) TestGetAttributeInvalidID() {
	rec := suite.testGetByID("invalid")

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(404))
	g.Expect(rec.Body).To(MatchJSON(`{
		"type": "tag:wyrdwest,2019:problems/attributes/unknown-attribute",
		"title": "The Attribute was not found",
		"status": 404
	}`))
}

func (suite *HTTPSuite) TestGetKnownAttributeByID() {
	id := attributes.AttributeID(uuid.Must(uuid.FromString("00000000-0000-0000-0000-000000000000")))
	suite.mockRetriever.EXPECT().
		GetAttributeByID(id).
		Return(attributes.Attribute{
			ID:          id,
			Version:     uuid.NewV4(),
			Created:     time.Now(),
			Updated:     time.Now(),
			Name:        "Strength",
			Description: "How Strong I am",
		}, nil).
		Times(1)

	rec := suite.testGetByID("00000000-0000-0000-0000-000000000000")

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(200))
	g.Expect(rec.Body).To(MatchJSON(`{
		"self": "/attributes/00000000-0000-0000-0000-000000000000",
		"name": "Strength",
		"description": "How Strong I am"
	}`))
}

func (suite *HTTPSuite) TestGetUnexpectedError() {
	suite.mockRetriever.EXPECT().
		GetAttributeByID(attributes.AttributeID(uuid.Must(uuid.FromString("00000000-0000-0000-0000-000000000000")))).
		Return(attributes.Attribute{}, errors.New("Oops")).
		Times(1)

	rec := suite.testGetByID("00000000-0000-0000-0000-000000000000")

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(500))
	g.Expect(rec.Body).To(MatchJSON(`{
		"type": "tag:wyrdwest,2019:problems/internal-server-error",
		"title": "An unexpected error occurred",
		"status": 500
}`))
}

package health_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/gomega"
	"github.com/sazzer/wyrdwest/service/internal/health"
	"github.com/stretchr/testify/suite"
)

type HTTPSuite struct {
	suite.Suite
	healthchecker health.Healthchecker
	mockCtrl      *gomock.Controller
}

func (suite *HTTPSuite) SetupTest() {
	suite.healthchecker = health.New()
	suite.mockCtrl = gomock.NewController(suite.T())
}

func (suite HTTPSuite) runTest() *httptest.ResponseRecorder {
	e := echo.New()
	health.RegisterHandler(&suite.healthchecker)(e)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	return rec
}

func (suite *HTTPSuite) TestNoHealthchecks() {
	rec := suite.runTest()

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(200))
	g.Expect(rec.Body).To(MatchJSON(`{}`))
}

func (suite *HTTPSuite) TestOnePassingHealthchecks() {
	mockHealthcheck := NewMockHealthcheck(suite.mockCtrl)
	mockHealthcheck.EXPECT().CheckHealth().Return(nil).Times(1)

	suite.healthchecker.AddHealthcheck("passing", mockHealthcheck)
	rec := suite.runTest()

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(200))
	g.Expect(rec.Body).To(MatchJSON(`{"passing": {"status": "ok"}}`))
}

func (suite *HTTPSuite) TestOneFailingHealthchecks() {
	mockHealthcheck := NewMockHealthcheck(suite.mockCtrl)
	mockHealthcheck.EXPECT().CheckHealth().Return(errors.New("Oops")).Times(1)

	suite.healthchecker.AddHealthcheck("failing", mockHealthcheck)
	rec := suite.runTest()

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(503))
	g.Expect(rec.Body).To(MatchJSON(`{"failing": {"status": "fail", "message": "Oops"}}`))
}

func (suite *HTTPSuite) TestMixedHealthchecks() {
	mockFailingHealthcheck := NewMockHealthcheck(suite.mockCtrl)
	mockFailingHealthcheck.EXPECT().CheckHealth().Return(errors.New("Oops")).Times(1)

	mockPassingHealthcheck := NewMockHealthcheck(suite.mockCtrl)
	mockPassingHealthcheck.EXPECT().CheckHealth().Return(nil).Times(1)

	suite.healthchecker.AddHealthcheck("failing", mockFailingHealthcheck)
	suite.healthchecker.AddHealthcheck("passing", mockPassingHealthcheck)

	rec := suite.runTest()

	g := NewGomegaWithT(suite.T())
	g.Expect(rec.Code).To(Equal(503))
	g.Expect(rec.Body).To(MatchJSON(`{"passing": {"status": "ok"}, "failing": {"status": "fail", "message": "Oops"}}`))
}

func TestHTTPSuite(t *testing.T) {
	suite.Run(t, new(HTTPSuite))
}

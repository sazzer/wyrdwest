package http_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes/mocks"
	"github.com/stretchr/testify/suite"
)

type HTTPSuite struct {
	suite.Suite
	mockRetriever *mocks.MockRetriever
	mockCtrl      *gomock.Controller
}

func (suite *HTTPSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockRetriever = mocks.NewMockRetriever(suite.mockCtrl)
}

func TestHTTPSuite(t *testing.T) {
	suite.Run(t, new(HTTPSuite))
}

package dao_test

import (
	"errors"
	"time"

	"github.com/lib/pq"

	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
	"github.com/sazzer/wyrdwest/service/internal/oauth2"
)

func (suite *DAOSuite) TestGetUnknownByID() {
	var (
		id = uuid.NewV4().String()
	)

	rows := sqlmock.NewRows([]string{})
	suite.mockCtrl.ExpectQuery("SELECT \\* FROM oauth2_clients WHERE client_id = \\$1").
		WithArgs(id).
		RowsWillBeClosed().
		WillReturnRows(rows)

	client, err := suite.testSubject.GetClientByID(oauth2.ClientID(id))
	suite.Assert().Error(err)
	suite.Assert().Equal(oauth2.ClientNotFoundError{}, err)
	suite.Assert().Equal(oauth2.Client{}, client)
}

func (suite *DAOSuite) TestGetUnknownByIDAndSecret() {
	var (
		id     = uuid.NewV4().String()
		secret = uuid.NewV4().String()
	)

	rows := sqlmock.NewRows([]string{})
	suite.mockCtrl.ExpectQuery("SELECT \\* FROM oauth2_clients WHERE client_id = \\$1 AND client_secret = \\$2").
		WithArgs(id, secret).
		RowsWillBeClosed().
		WillReturnRows(rows)

	client, err := suite.testSubject.GetClientByIDAndSecret(oauth2.ClientID(id), secret)
	suite.Assert().Error(err)
	suite.Assert().Equal(oauth2.ClientNotFoundError{}, err)
	suite.Assert().Equal(oauth2.Client{}, client)
}

func (suite *DAOSuite) TestGetByIDDatabaseError() {
	var (
		id = uuid.NewV4().String()
	)

	suite.mockCtrl.ExpectQuery("SELECT \\* FROM oauth2_clients WHERE client_id = \\$1").
		WithArgs(id).
		WillReturnError(errors.New("It be broke"))

	client, err := suite.testSubject.GetClientByID(oauth2.ClientID(id))
	suite.Assert().EqualError(err, "It be broke")
	suite.Assert().Equal(oauth2.Client{}, client)
}

func (suite *DAOSuite) TestGetByIDAndSecretDatabaseError() {
	var (
		id     = uuid.NewV4().String()
		secret = uuid.NewV4().String()
	)

	suite.mockCtrl.ExpectQuery("SELECT \\* FROM oauth2_clients WHERE client_id = \\$1 AND client_secret = \\$2").
		WithArgs(id, secret).
		WillReturnError(errors.New("It be broke"))

	client, err := suite.testSubject.GetClientByIDAndSecret(oauth2.ClientID(id), secret)
	suite.Assert().EqualError(err, "It be broke")
	suite.Assert().Equal(oauth2.Client{}, client)
}

func (suite *DAOSuite) TestGetKnownByID() {
	var (
		id      = uuid.NewV4().String()
		version = uuid.NewV4().String()
		now     = time.Now()
		secret  = uuid.NewV4().String()
		owner   = uuid.NewV4().String()
	)

	rows := sqlmock.NewRows([]string{"client_id", "version", "created", "updated", "name", "owner_id", "client_secret", "redirect_uris", "response_types", "grant_types"})
	rows.AddRow(id,
		version,
		now,
		now,
		"Testing",
		owner,
		secret,
		pq.Array([]string{"http://localhost/redirect"}),
		pq.Array([]string{"code", "token"}),
		pq.Array([]string{"authorization_code", "implicit"}),
	)
	suite.mockCtrl.ExpectQuery("SELECT \\* FROM oauth2_clients WHERE client_id = \\$1").
		WithArgs(id).
		RowsWillBeClosed().
		WillReturnRows(rows)

	client, err := suite.testSubject.GetClientByID(oauth2.ClientID(id))
	suite.Assert().Error(err)
	suite.Assert().Equal(oauth2.ClientNotFoundError{}, err)
	suite.Assert().Equal(oauth2.Client{}, client)
}

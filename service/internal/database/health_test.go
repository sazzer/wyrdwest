package database_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sazzer/wyrdwest/service/internal/database"
	"github.com/stretchr/testify/suite"
)

type HealthSuite struct {
	suite.Suite
	db       database.DB
	mockDB   *sql.DB
	mockCtrl sqlmock.Sqlmock
}

func (suite *HealthSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	suite.Assert().NoError(err)

	suite.db = database.NewFromDB(db)
	suite.mockDB = db
	suite.mockCtrl = mock
}

func (suite *HealthSuite) TearDownTest() {
	suite.mockDB.Close()
	suite.Assert().NoError(suite.mockCtrl.ExpectationsWereMet())
}

func TestHealthSuite(t *testing.T) {
	suite.Run(t, new(HealthSuite))
}

func (suite *HealthSuite) TestHealthy() {
	suite.mockCtrl.ExpectExec("SELECT 1").WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.db.CheckHealth()
	suite.Assert().NoError(err)
}

func (suite *HealthSuite) TestNotHealthy() {
	suite.mockCtrl.ExpectExec("SELECT 1").WillReturnError(errors.New("Oops"))

	err := suite.db.CheckHealth()
	suite.Assert().Error(err, "Oops")
}

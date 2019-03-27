package dao_test

import (
	"database/sql"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/sazzer/wyrdwest/service/internal/characters/attributes/dao"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sazzer/wyrdwest/service/internal/database"
	"github.com/stretchr/testify/suite"
)

type DAOSuite struct {
	suite.Suite
	mockDB      *sql.DB
	mockCtrl    sqlmock.Sqlmock
	testSubject dao.AttributesDao
}

func (suite *DAOSuite) SetupTest() {
	logrus.SetLevel(logrus.DebugLevel)

	db, mock, err := sqlmock.New()
	suite.Assert().NoError(err)

	suite.mockDB = db
	suite.mockCtrl = mock

	suite.testSubject = dao.New(database.NewFromDB(db))
}

func (suite *DAOSuite) TearDownTest() {
	suite.mockDB.Close()
	suite.Assert().NoError(suite.mockCtrl.ExpectationsWereMet())
}

func TestDAO(t *testing.T) {
	suite.Run(t, new(DAOSuite))
}

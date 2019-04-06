package database_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sazzer/wyrdwest/service/internal/database"
	"github.com/stretchr/testify/suite"
)

type QuerySuite struct {
	suite.Suite
	db       database.DB
	mockDB   *sql.DB
	mockCtrl sqlmock.Sqlmock
}

func (suite *QuerySuite) SetupTest() {
	db, mock, err := sqlmock.New()
	suite.Assert().NoError(err)

	suite.db = database.NewFromDB(db)
	suite.mockDB = db
	suite.mockCtrl = mock
}

func (suite *QuerySuite) TearDownTest() {
	suite.mockDB.Close()
	suite.Assert().NoError(suite.mockCtrl.ExpectationsWereMet())
}

func TestQuery(t *testing.T) {
	suite.Run(t, new(QuerySuite))
}

func (suite *QuerySuite) TestQueryCallbackNoRows() {
	sqlBuilder := squirrel.Select("*").From("table")
	counts := []int{}

	rows := sqlmock.NewRows([]string{"a"})
	suite.mockCtrl.ExpectQuery(regexp.QuoteMeta("SELECT * FROM table")).
		RowsWillBeClosed().
		WillReturnRows(rows)

	err := suite.db.QueryWithCallback(sqlBuilder, func(row *sqlx.Rows) error {
		var c int
		err := row.Scan(&c)
		suite.Assert().NoError(err)
		counts = append(counts, c)
		return nil
	})

	suite.Assert().NoError(err)
	suite.Assert().Empty(counts)
}

func (suite *QuerySuite) TestQueryCallbackRows() {
	sqlBuilder := squirrel.Select("*").From("table")
	counts := []int{}

	rows := sqlmock.NewRows([]string{"a"})
	rows.AddRow(1)
	rows.AddRow(2)
	rows.AddRow(3)
	suite.mockCtrl.ExpectQuery(regexp.QuoteMeta("SELECT * FROM table")).
		RowsWillBeClosed().
		WillReturnRows(rows)

	err := suite.db.QueryWithCallback(sqlBuilder, func(row *sqlx.Rows) error {
		var c int
		err := row.Scan(&c)
		suite.Assert().NoError(err)
		counts = append(counts, c)
		return nil
	})

	suite.Assert().NoError(err)
	suite.Assert().Equal([]int{1, 2, 3}, counts)
}

func (suite *QuerySuite) TestQueryCallbackErrors() {
	sqlBuilder := squirrel.Select("*").From("table")
	counts := []int{}

	rows := sqlmock.NewRows([]string{"a"})
	rows.AddRow(1)
	rows.AddRow(2)
	rows.AddRow(3)
	suite.mockCtrl.ExpectQuery(regexp.QuoteMeta("SELECT * FROM table")).
		RowsWillBeClosed().
		WillReturnRows(rows)

	err := suite.db.QueryWithCallback(sqlBuilder, func(row *sqlx.Rows) error {
		var c int
		err := row.Scan(&c)
		suite.Assert().NoError(err)
		counts = append(counts, c)
		if c == 2 {
			return errors.New("c = 2")
		}
		return nil
	})

	suite.Assert().Error(err)
	suite.Assert().Equal([]int{1, 2}, counts)
}

func (suite *QuerySuite) TestQueryCallbackDatabaseErrors() {
	sqlBuilder := squirrel.Select("*").From("table")
	counts := []int{}

	suite.mockCtrl.ExpectQuery(regexp.QuoteMeta("SELECT * FROM table")).
		WillReturnError(errors.New("sql"))

	err := suite.db.QueryWithCallback(sqlBuilder, func(row *sqlx.Rows) error {
		var c int
		err := row.Scan(&c)
		suite.Assert().NoError(err)
		counts = append(counts, c)
		return nil
	})

	suite.Assert().Error(err)
	suite.Assert().Empty(counts)
}

func (suite *QuerySuite) TestQueryOneCallbackNoRows() {
	sqlBuilder := squirrel.Select("*").From("table")
	counts := []int{}

	rows := sqlmock.NewRows([]string{"a"})
	suite.mockCtrl.ExpectQuery(regexp.QuoteMeta("SELECT * FROM table")).
		RowsWillBeClosed().
		WillReturnRows(rows)

	err := suite.db.QueryOneWithCallback(sqlBuilder, func(row *sqlx.Rows) error {
		var c int
		err := row.Scan(&c)
		suite.Assert().NoError(err)
		counts = append(counts, c)
		return nil
	})

	suite.Assert().Equal(database.RecordNotFoundError{}, err)
	suite.Assert().Empty(counts)
}

func (suite *QuerySuite) TestQueryOneCallbackOneRow() {
	sqlBuilder := squirrel.Select("*").From("table")
	counts := []int{}

	rows := sqlmock.NewRows([]string{"a"})
	rows.AddRow(1)
	suite.mockCtrl.ExpectQuery(regexp.QuoteMeta("SELECT * FROM table")).
		RowsWillBeClosed().
		WillReturnRows(rows)

	err := suite.db.QueryOneWithCallback(sqlBuilder, func(row *sqlx.Rows) error {
		var c int
		err := row.Scan(&c)
		suite.Assert().NoError(err)
		counts = append(counts, c)
		return nil
	})

	suite.Assert().NoError(err)
	suite.Assert().Equal([]int{1}, counts)
}

func (suite *QuerySuite) TestQueryOneCallbackTwoRow() {
	sqlBuilder := squirrel.Select("*").From("table")
	counts := []int{}

	rows := sqlmock.NewRows([]string{"a"})
	rows.AddRow(1)
	rows.AddRow(2)
	suite.mockCtrl.ExpectQuery(regexp.QuoteMeta("SELECT * FROM table")).
		RowsWillBeClosed().
		WillReturnRows(rows)

	err := suite.db.QueryOneWithCallback(sqlBuilder, func(row *sqlx.Rows) error {
		var c int
		err := row.Scan(&c)
		suite.Assert().NoError(err)
		counts = append(counts, c)
		return nil
	})

	suite.Assert().Equal(database.MultipleRecordFoundError{}, err)
	suite.Assert().Equal([]int{1}, counts)
}

func (suite *QuerySuite) TestQueryOneCallbackErrors() {
	sqlBuilder := squirrel.Select("*").From("table")
	counts := []int{}

	rows := sqlmock.NewRows([]string{"a"})
	rows.AddRow(1)
	suite.mockCtrl.ExpectQuery(regexp.QuoteMeta("SELECT * FROM table")).
		RowsWillBeClosed().
		WillReturnRows(rows)

	err := suite.db.QueryOneWithCallback(sqlBuilder, func(row *sqlx.Rows) error {
		var c int
		err := row.Scan(&c)
		suite.Assert().NoError(err)
		counts = append(counts, c)
		return errors.New("c = 1")
	})

	suite.Assert().Error(err)
	suite.Assert().Equal([]int{1}, counts)
}

func (suite *QuerySuite) TestQueryOneCallbackDatabaseErrors() {
	sqlBuilder := squirrel.Select("*").From("table")
	counts := []int{}

	suite.mockCtrl.ExpectQuery(regexp.QuoteMeta("SELECT * FROM table")).
		WillReturnError(errors.New("sql"))

	err := suite.db.QueryOneWithCallback(sqlBuilder, func(row *sqlx.Rows) error {
		var c int
		err := row.Scan(&c)
		suite.Assert().NoError(err)
		counts = append(counts, c)
		return nil
	})

	suite.Assert().Error(err)
	suite.Assert().Empty(counts)
}

package dao_test

import (
	"errors"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
)

func (suite *DAOSuite) TestGetUnknownByID() {
	var (
		id = uuid.NewV4()
	)

	rows := sqlmock.NewRows([]string{})
	suite.mockCtrl.ExpectQuery("SELECT \\* FROM attributes WHERE attribute_id = \\$1").
		WithArgs(id.String()).
		RowsWillBeClosed().
		WillReturnRows(rows)

	attribute, err := suite.testSubject.GetAttributeByID(attributes.AttributeID(id))
	suite.Assert().Error(err)
	suite.Assert().Equal(attributes.AttributeNotFoundError{}, err)
	suite.Assert().Equal(attributes.Attribute{}, attribute)
}

func (suite *DAOSuite) TestGetKnownByID() {
	var (
		id      = uuid.NewV4()
		version = uuid.NewV4()
	)

	now := time.Now()

	rows := sqlmock.NewRows([]string{"attribute_id", "version", "created", "updated", "name", "description"})
	rows.AddRow(id.String(), version.String(), now, now, "Strength", "How strong I am")

	suite.mockCtrl.ExpectQuery("SELECT \\* FROM attributes WHERE attribute_id = \\$1").
		WithArgs(id.String()).
		RowsWillBeClosed().
		WillReturnRows(rows)

	attribute, err := suite.testSubject.GetAttributeByID(attributes.AttributeID(id))
	suite.Assert().NoError(err)

	suite.Assert().Equal(attributes.Attribute{
		ID:          attributes.AttributeID(id),
		Version:     version,
		Created:     now,
		Updated:     now,
		Name:        "Strength",
		Description: "How strong I am",
	}, attribute)
}

func (suite *DAOSuite) TestGetDatabaseError() {
	var (
		id = uuid.NewV4()
	)

	suite.mockCtrl.ExpectQuery("SELECT \\* FROM attributes WHERE attribute_id = \\$1").
		WithArgs(id.String()).
		WillReturnError(errors.New("It be broke"))

	attribute, err := suite.testSubject.GetAttributeByID(attributes.AttributeID(id))
	suite.Assert().EqualError(err, "It be broke")
	suite.Assert().Equal(attributes.Attribute{}, attribute)
}

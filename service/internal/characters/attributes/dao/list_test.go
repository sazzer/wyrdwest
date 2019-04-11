package dao_test

import (
	"database/sql/driver"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"

	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
	"github.com/sazzer/wyrdwest/service/internal/service"
)

func (suite *DAOSuite) TestGetNoRows() {
	type test struct {
		name     string
		criteria attributes.AttributeMatchCriteria
		sorts    []service.SortField
		offset   uint64
		pageSize uint64
		sql      string
		binds    []driver.Value
	}

	tests := []test{
		{
			name: "No Nothing",
			sql:  "SELECT \\* FROM attributes ORDER BY name ASC, attribute_id DESC LIMIT 0 OFFSET 0",
		},
		{
			name:   "Offset Set",
			offset: 10,
			sql:    "SELECT \\* FROM attributes ORDER BY name ASC, attribute_id DESC LIMIT 0 OFFSET 10",
		},
		{
			name:     "Limit Set",
			pageSize: 10,
			sql:      "SELECT \\* FROM attributes ORDER BY name ASC, attribute_id DESC LIMIT 10 OFFSET 0",
		},
		{
			name:  "Sort By Name Ascending",
			sorts: []service.SortField{{Field: "name", Direction: service.SortAscending}},
			sql:   "SELECT \\* FROM attributes ORDER BY name ASC, name ASC, attribute_id DESC LIMIT 0 OFFSET 0",
		},
		{
			name:  "Sort By Name Descending",
			sorts: []service.SortField{{Field: "name", Direction: service.SortDescending}},
			sql:   "SELECT \\* FROM attributes ORDER BY name DESC, name ASC, attribute_id DESC LIMIT 0 OFFSET 0",
		},
		{
			name:  "Sort By Name - Natural Order",
			sorts: []service.SortField{{Field: "name", Direction: service.SortNatural}},
			sql:   "SELECT \\* FROM attributes ORDER BY name ASC, name ASC, attribute_id DESC LIMIT 0 OFFSET 0",
		},
		{
			name: "Sort By Name Ascending and Created Descending",
			sorts: []service.SortField{
				{Field: "name", Direction: service.SortAscending},
				{Field: "created", Direction: service.SortDescending},
			},
			sql: "SELECT \\* FROM attributes ORDER BY name ASC, created DESC, name ASC, attribute_id DESC LIMIT 0 OFFSET 0",
		},
		{
			name: "Filter by Name",
			criteria: attributes.AttributeMatchCriteria{
				Name: "strength",
			},
			sql:   "SELECT \\* FROM attributes WHERE UPPER\\(name\\) = \\$1 ORDER BY name ASC, attribute_id DESC LIMIT 0 OFFSET 0",
			binds: []driver.Value{"STRENGTH"},
		},
	}

	for _, tt := range tests {
		tt := tt
		suite.Run(tt.name, func() {
			rows := sqlmock.NewRows([]string{})
			suite.mockCtrl.ExpectQuery(tt.sql).
				WithArgs(tt.binds...).
				RowsWillBeClosed().
				WillReturnRows(rows)

			countRows := sqlmock.NewRows([]string{"c"})
			countRows.AddRow(0)
			suite.mockCtrl.ExpectQuery("SELECT COUNT\\(\\*\\) AS c").
				RowsWillBeClosed().
				WillReturnRows(countRows)

			attributes, err := suite.testSubject.ListAttributes(tt.criteria, tt.sorts, tt.offset, tt.pageSize)
			suite.Assert().NoError(err)

			suite.Assert().Equal(uint64(0), attributes.TotalSize)
			suite.Assert().Equal(0, len(attributes.Data))
		})
	}
}

func (suite *DAOSuite) TestGetOnlyPage() {
	var (
		strengthID          = uuid.NewV4().String()
		strengthVersion     = uuid.NewV4().String()
		intelligenceID      = uuid.NewV4().String()
		intelligenceVersion = uuid.NewV4().String()
		now                 = time.Now()
	)

	rows := sqlmock.NewRows([]string{"attribute_id", "version", "created", "updated", "name", "description"})
	rows.AddRow(strengthID, strengthVersion, now, now, "Strength", "How strong I am")
	rows.AddRow(intelligenceID, intelligenceVersion, now, now, "Intelligence", "How intelligent I am")
	suite.mockCtrl.ExpectQuery("SELECT \\* FROM attributes ORDER BY name ASC, attribute_id DESC LIMIT 10 OFFSET 0").
		RowsWillBeClosed().
		WillReturnRows(rows)

	results, err := suite.testSubject.ListAttributes(attributes.AttributeMatchCriteria{}, []service.SortField{}, 0, 10)
	suite.Assert().NoError(err)

	suite.Assert().Equal(uint64(2), results.TotalSize)
	suite.Require().Equal(2, len(results.Data))

	suite.Assert().Equal(attributes.Attribute{
		ID:          attributes.AttributeID(strengthID),
		Version:     strengthVersion,
		Created:     now,
		Updated:     now,
		Name:        "Strength",
		Description: "How strong I am",
	}, results.Data[0])

	suite.Assert().Equal(attributes.Attribute{
		ID:          attributes.AttributeID(intelligenceID),
		Version:     intelligenceVersion,
		Created:     now,
		Updated:     now,
		Name:        "Intelligence",
		Description: "How intelligent I am",
	}, results.Data[1])
}

func (suite *DAOSuite) TestGetFirstPage() {
	dataRows := sqlmock.NewRows([]string{"attribute_id", "version", "created", "updated", "name", "description"})
	dataRows.AddRow(uuid.NewV4().String(), uuid.NewV4().String(), time.Now(), time.Now(), "Strength", "How strong I am")
	suite.mockCtrl.ExpectQuery("SELECT \\* FROM attributes ORDER BY name ASC, attribute_id DESC LIMIT 1 OFFSET 0").
		RowsWillBeClosed().
		WillReturnRows(dataRows)

	countRows := sqlmock.NewRows([]string{"c"})
	countRows.AddRow(5)
	suite.mockCtrl.ExpectQuery("SELECT COUNT\\(\\*\\) AS c FROM attributes").
		RowsWillBeClosed().
		WillReturnRows(countRows)

	results, err := suite.testSubject.ListAttributes(attributes.AttributeMatchCriteria{}, []service.SortField{}, 0, 1)
	suite.Assert().NoError(err)

	suite.Assert().Equal(uint64(5), results.TotalSize)
	suite.Require().Equal(1, len(results.Data))
}

func (suite *DAOSuite) TestGetMidPage() {
	dataRows := sqlmock.NewRows([]string{"attribute_id", "version", "created", "updated", "name", "description"})
	dataRows.AddRow(uuid.NewV4().String(), uuid.NewV4().String(), time.Now(), time.Now(), "Strength", "How strong I am")
	suite.mockCtrl.ExpectQuery("SELECT \\* FROM attributes ORDER BY name ASC, attribute_id DESC LIMIT 1 OFFSET 20").
		RowsWillBeClosed().
		WillReturnRows(dataRows)

	countRows := sqlmock.NewRows([]string{"c"})
	countRows.AddRow(25)
	suite.mockCtrl.ExpectQuery("SELECT COUNT\\(\\*\\) AS c FROM attributes").
		RowsWillBeClosed().
		WillReturnRows(countRows)

	results, err := suite.testSubject.ListAttributes(attributes.AttributeMatchCriteria{}, []service.SortField{}, 20, 1)
	suite.Assert().NoError(err)

	suite.Assert().Equal(uint64(25), results.TotalSize)
	suite.Require().Equal(1, len(results.Data))
}

func (suite *DAOSuite) TestGetLastPage() {
	dataRows := sqlmock.NewRows([]string{"attribute_id", "version", "created", "updated", "name", "description"})
	dataRows.AddRow(uuid.NewV4().String(), uuid.NewV4().String(), time.Now(), time.Now(), "Strength", "How strong I am")
	suite.mockCtrl.ExpectQuery("SELECT \\* FROM attributes ORDER BY name ASC, attribute_id DESC LIMIT 10 OFFSET 20").
		RowsWillBeClosed().
		WillReturnRows(dataRows)

	results, err := suite.testSubject.ListAttributes(attributes.AttributeMatchCriteria{}, []service.SortField{}, 20, 10)
	suite.Assert().NoError(err)

	suite.Assert().Equal(uint64(21), results.TotalSize)
	suite.Require().Equal(1, len(results.Data))
}

func (suite *DAOSuite) TestGetFirstPageFiltered() {
	dataRows := sqlmock.NewRows([]string{"attribute_id", "version", "created", "updated", "name", "description"})
	dataRows.AddRow(uuid.NewV4().String(), uuid.NewV4().String(), time.Now(), time.Now(), "Strength", "How strong I am")
	suite.mockCtrl.ExpectQuery("SELECT \\* FROM attributes WHERE UPPER\\(name\\) = \\$1 ORDER BY name ASC, attribute_id DESC LIMIT 1 OFFSET 0").
		WithArgs("STRENGTH").
		RowsWillBeClosed().
		WillReturnRows(dataRows)

	countRows := sqlmock.NewRows([]string{"c"})
	countRows.AddRow(5)
	suite.mockCtrl.ExpectQuery("SELECT COUNT\\(\\*\\) AS c FROM attributes WHERE UPPER\\(name\\) = \\$1").
		WithArgs("STRENGTH").
		RowsWillBeClosed().
		WillReturnRows(countRows)

	results, err := suite.testSubject.ListAttributes(attributes.AttributeMatchCriteria{Name: "Strength"}, []service.SortField{}, 0, 1)
	suite.Assert().NoError(err)

	suite.Assert().Equal(uint64(5), results.TotalSize)
	suite.Require().Equal(1, len(results.Data))
}

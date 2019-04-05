package dao_test

import (
	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"

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
			sql:   "SELECT \\* FROM attributes WHERE UPPER\\(name\\) = \\? ORDER BY name ASC, attribute_id DESC LIMIT 0 OFFSET 0",
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

			attributes, err := suite.testSubject.ListAttributes(tt.criteria, tt.sorts, tt.offset, tt.pageSize)
			suite.Assert().NoError(err)

			suite.Assert().Equal(0, attributes.TotalSize)
			suite.Assert().Equal(0, len(attributes.Data))
		})
	}
}

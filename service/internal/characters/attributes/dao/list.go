package dao

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"

	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
	"github.com/sazzer/wyrdwest/service/internal/service"
)

func getListRowsQuery(criteria attributes.AttributeMatchCriteria, sorts []service.SortField, offset uint64, count uint64) squirrel.SelectBuilder {
	var (
		sortFields = map[string]string{
			"name":    "ASC",
			"created": "DESC",
		}
	)

	// Base query to execute
	sqlBuilder := squirrel.Select("*").From("attributes").Offset(offset).Limit(count)

	// Add all the criteria
	if criteria.Name != "" {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"UPPER(name)": strings.ToUpper(criteria.Name)})
	}

	// Add all the sorts
	for _, sort := range sorts {
		if defaultSort, ok := sortFields[sort.Field]; ok {
			sortDir := defaultSort
			switch sort.Direction {
			case service.SortAscending:
				sortDir = "ASC"
			case service.SortDescending:
				sortDir = "DESC"
			}
			sqlBuilder = sqlBuilder.OrderBy(fmt.Sprintf("%s %s", sort.Field, sortDir))
		} else {
			logrus.WithField("sort", sort.Field).Error("unknown sort field")
		}
	}

	// Add a default sort to guarantee consistency if nothing else
	sqlBuilder = sqlBuilder.OrderBy("name ASC")
	sqlBuilder = sqlBuilder.OrderBy("attribute_id DESC")

	return sqlBuilder
}

func getCountQuery(criteria attributes.AttributeMatchCriteria) squirrel.SelectBuilder {
	// Base query to execute
	sqlBuilder := squirrel.Select("COUNT(*) AS c").From("attributes")

	// Add all the criteria
	if criteria.Name != "" {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"UPPER(name)": strings.ToUpper(criteria.Name)})
	}

	return sqlBuilder
}

// ListAttributes allows us to get a list of attributes that match certain criteria
func (dao AttributesDao) ListAttributes(criteria attributes.AttributeMatchCriteria, sorts []service.SortField, offset uint64, count uint64) (attributes.AttributePage, error) {

	results := []attributes.Attribute{}
	if err := dao.db.QueryWithCallback(getListRowsQuery(criteria, sorts, offset, count), func(row *sqlx.Rows) error {
		resultRow := dbAttribute{}
		err := row.StructScan(&resultRow)
		if err != nil {
			logrus.WithError(err).Error("Failed to parse attribute")
			return err
		}
		logrus.WithField("row", resultRow).Debug("Loaded attribute data")

		results = append(results, resultRow.ToAPI())

		return nil
	}); err != nil {
		logrus.WithError(err).Error("Failed to read attributes")
		return attributes.AttributePage{}, err
	}

	numResults := uint64(len(results))
	totalSize := numResults + offset
	if numResults == 0 || numResults == count {
		// Either we got no rows back, or we got exactly as many as we asked for. That means we dont know the total size
		// So we need to go and find out
		if err := dao.db.QueryOneWithCallback(getCountQuery(criteria), func(row *sqlx.Rows) error {
			return row.Scan(&totalSize)
		}); err != nil {
			logrus.WithError(err).Error("Failed to count attributes")
			return attributes.AttributePage{}, err
		}
	}

	return attributes.AttributePage{
		PageInfo: service.PageInfo{
			TotalSize: totalSize,
		},
		Data: results,
	}, nil
}

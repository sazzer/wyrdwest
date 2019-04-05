package dao

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"

	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
	"github.com/sazzer/wyrdwest/service/internal/service"
)

func (dao AttributesDao) getListRows(criteria attributes.AttributeMatchCriteria, sorts []service.SortField, offset uint64, count uint64) (*sqlx.Rows, error) {
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
			return nil, errors.New("unknown sort field")
		}
	}

	// Add a default sort to guarantee consistency if nothing else
	sqlBuilder = sqlBuilder.OrderBy("name ASC")
	sqlBuilder = sqlBuilder.OrderBy("attribute_id DESC")

	sql, args, err := sqlBuilder.ToSql()
	if err != nil {
		logrus.WithError(err).Error("Failed to build attributes list SQL")
		return nil, err
	}

	return dao.db.QueryPositional(sql, args)
}

func (dao AttributesDao) getCount(criteria attributes.AttributeMatchCriteria) (uint64, error) {
	// Base query to execute
	sqlBuilder := squirrel.Select("COUNT(*) AS c").From("attributes")

	// Add all the criteria
	if criteria.Name != "" {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"UPPER(name)": strings.ToUpper(criteria.Name)})
	}

	sql, args, err := sqlBuilder.ToSql()
	if err != nil {
		logrus.WithError(err).Error("Failed to build attributes count SQL")
		return 0, err
	}

	rows, err := dao.db.QueryPositional(sql, args)
	if err != nil {
		logrus.WithError(err).Error("Failed to count attributes")
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		logrus.Debug("No count returned")
		return 0, err
	}

	var count uint64
	err = rows.Scan(&count)
	if err != nil {
		logrus.WithError(err).Error("Failed to retrieve attribute count")
		return 0, err
	}

	return count, nil
}

// ListAttributes allows us to get a list of attributes that match certain criteria
func (dao AttributesDao) ListAttributes(criteria attributes.AttributeMatchCriteria, sorts []service.SortField, offset uint64, count uint64) (attributes.AttributePage, error) {
	rows, err := dao.getListRows(criteria, sorts, offset, count)
	if err != nil {
		logrus.WithError(err).Error("Failed to load attributes")
		return attributes.AttributePage{}, err
	}
	defer rows.Close()

	results := []attributes.Attribute{}
	for rows.Next() {
		resultRow := dbAttribute{}
		err = rows.StructScan(&resultRow)
		if err != nil {
			logrus.WithError(err).Error("Failed to parse attribute")
			return attributes.AttributePage{}, err
		}
		logrus.WithField("row", resultRow).Debug("Loaded attribute data")

		results = append(results, resultRow.ToAPI())
	}

	numResults := uint64(len(results))
	totalSize := numResults + offset
	if numResults == 0 || numResults == count {
		// Either we got no rows back, or we got exactly as many as we asked for. That means we dont know the total size
		// So we need to go and find out
		totalSize, err = dao.getCount(criteria)
		if err != nil {
			logrus.WithError(err).Error("Failed to count attribute")
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

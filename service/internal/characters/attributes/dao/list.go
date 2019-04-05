package dao

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"

	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
	"github.com/sazzer/wyrdwest/service/internal/service"
)

var (
	sortFields = map[string]string{
		"name":    "ASC",
		"created": "DESC",
	}
)

// ListAttributes allows us to get a list of attributes that match certain criteria
func (dao AttributesDao) ListAttributes(criteria attributes.AttributeMatchCriteria, sorts []service.SortField, offset uint64, count uint64) (attributes.AttributePage, error) {
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
			logrus.WithField("sort", sort.Field).Error("Unknown sort field")
			return attributes.AttributePage{}, errors.New("Unknown sort field")
		}
	}

	// Add a default sort to guarantee consistency if nothing else
	sqlBuilder = sqlBuilder.OrderBy("name ASC")
	sqlBuilder = sqlBuilder.OrderBy("attribute_id DESC")

	sql, args, err := sqlBuilder.ToSql()
	if err != nil {
		logrus.WithError(err).Error("Failed to build attributes list SQL")
		return attributes.AttributePage{}, err
	}

	rows, err := dao.db.QueryPositional(sql, args)
	if err != nil {
		logrus.WithError(err).Error("Failed to load attributes")
		return attributes.AttributePage{}, err
	}
	defer rows.Close()

	return attributes.AttributePage{
		PageInfo: service.PageInfo{
			TotalSize: 0,
		},
		Data: []attributes.Attribute{},
	}, nil
}

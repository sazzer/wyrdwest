package dao

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
	"github.com/sazzer/wyrdwest/service/internal/database"
	"github.com/sirupsen/logrus"
)

// GetAttributeByID returns the Attribute with the given ID, or an error if it couldn't be loaded
func (dao AttributesDao) GetAttributeByID(id attributes.AttributeID) (attributes.Attribute, error) {
	sqlBuilder := squirrel.
		Select("*").
		From("attributes").
		Where(squirrel.Eq{"attribute_id": id}).
		PlaceholderFormat(squirrel.Dollar)

	resultRow := dbAttribute{}

	if err := dao.db.QueryOneWithCallback(sqlBuilder, func(row *sqlx.Rows) error {
		return row.StructScan(&resultRow)
	}); err != nil {
		logrus.WithError(err).Error("Failed to retrieve attributes")
		switch err.(type) {
		case database.RecordNotFoundError:
			return attributes.Attribute{}, attributes.AttributeNotFoundError{}
		default:
			return attributes.Attribute{}, err
		}
	}
	logrus.WithField("id", id).WithField("row", resultRow).Debug("Loaded attribute data")

	return resultRow.ToAPI(), nil
}

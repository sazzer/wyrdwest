package dao

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
	"github.com/sirupsen/logrus"
)

// GetAttributeByID returns the Attribute with the given ID, or an error if it couldn't be loaded
func (dao AttributesDao) GetAttributeByID(id attributes.AttributeID) (attributes.Attribute, error) {
	rows, err := dao.db.Query("SELECT * FROM attributes WHERE attribute_id = :id",
		map[string]interface{}{"id": uuid.UUID(id).String()})

	if err != nil {
		logrus.WithField("id", id).WithError(err).Error("Failed to load attribute")
		return attributes.Attribute{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		logrus.WithField("id", id).Debug("No matching attributes found")
		return attributes.Attribute{}, attributes.AttributeNotFoundError{}
	}

	resultRow := dbAttribute{}
	err = rows.StructScan(&resultRow)
	if err != nil {
		logrus.WithField("id", id).WithError(err).Error("Failed to parse attribute")
		return attributes.Attribute{}, err
	}
	logrus.WithField("id", id).WithField("row", resultRow).Debug("Loaded attribute data")

	return resultRow.ToAPI(), nil
}

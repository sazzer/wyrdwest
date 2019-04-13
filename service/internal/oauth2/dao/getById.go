package dao

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/sazzer/wyrdwest/service/internal/database"
	"github.com/sazzer/wyrdwest/service/internal/oauth2"
	"github.com/sirupsen/logrus"
)

func (dao OAuth2ClientsDao) getClientByQuery(query squirrel.SelectBuilder) (oauth2.Client, error) {
	resultRow := dbClient{}

	if err := dao.db.QueryOneWithCallback(query, func(row *sqlx.Rows) error {
		return row.StructScan(&resultRow)
	}); err != nil {
		logrus.WithError(err).Error("Failed to retrieve OAuth2 Client")
		switch err.(type) {
		case database.RecordNotFoundError:
			return oauth2.Client{}, oauth2.ClientNotFoundError{}
		default:
			return oauth2.Client{}, err
		}
	}
	logrus.WithField("row", resultRow).Debug("Loaded client data")

	return oauth2.Client{}, oauth2.ClientNotFoundError{}
}

// GetClientByID allows us to load a Client knowing only it's ID
func (dao OAuth2ClientsDao) GetClientByID(id oauth2.ClientID) (oauth2.Client, error) {
	sqlBuilder := squirrel.
		Select("*").
		From("oauth2_clients").
		Where(squirrel.Eq{"client_id": id}).
		PlaceholderFormat(squirrel.Dollar)

	return dao.getClientByQuery(sqlBuilder)
}

// GetClientByIDAndSecret allows us to load a Client knowing it's ID and Secret
func (dao OAuth2ClientsDao) GetClientByIDAndSecret(id oauth2.ClientID, secret string) (oauth2.Client, error) {
	sqlBuilder := squirrel.
		Select("*").
		From("oauth2_clients").
		Where(squirrel.Eq{"client_id": id}).
		Where(squirrel.Eq{"client_secret": secret}).
		PlaceholderFormat(squirrel.Dollar)

	return dao.getClientByQuery(sqlBuilder)
}

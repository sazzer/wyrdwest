package database

import (
	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
)

// QueryBuilder is an interface representing how to build a Query to execute
type QueryBuilder interface {
	ToSql() (string, []interface{}, error)
}

// QueryWithCallback executes a query and triggers the given callback for every row returned
func (db DB) QueryWithCallback(query QueryBuilder, callback func(*sqlx.Rows) error) error {
	sql, args, err := query.ToSql()
	if err != nil {
		logrus.WithError(err).Error("Failed to generate SQL to execute")
		return err
	}

	logrus.WithField("sql", sql).WithField("binds", args).Info("Executing query")
	rows, err := db.db.Queryx(sql, args...)
	if err != nil {
		logrus.WithField("sql", sql).WithField("binds", args).WithError(err).Error("Failed to execute SQL")
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = callback(rows)
		if err != nil {
			logrus.WithField("sql", sql).WithField("binds", args).WithError(err).Error("Failed to process row")
			return err
		}
	}
	return nil
}

// QueryOneWithCallback executes a query, expecting exactly one row to be returned
func (db DB) QueryOneWithCallback(query QueryBuilder, callback func(*sqlx.Rows) error) error {
	sql, args, err := query.ToSql()
	if err != nil {
		logrus.WithError(err).Error("Failed to generate SQL to execute")
		return err
	}

	logrus.WithField("sql", sql).WithField("binds", args).Info("Executing query")
	rows, err := db.db.Queryx(sql, args...)
	if err != nil {
		logrus.WithField("sql", sql).WithField("binds", args).WithError(err).Error("Failed to execute SQL")
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		logrus.WithField("sql", sql).WithField("binds", args).Error("Expected one row but none returned")
		return RecordNotFoundError{}
	}

	err = callback(rows)
	if err != nil {
		logrus.WithField("sql", sql).WithField("binds", args).WithError(err).Error("Failed to process row")
		return err
	}

	if rows.Next() {
		logrus.WithField("sql", sql).WithField("binds", args).Error("Expected one row but 2+ returned")
		return MultipleRecordFoundError{}
	}

	return nil
}

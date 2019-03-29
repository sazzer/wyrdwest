package database

import (
	"database/sql"
	"fmt"
	"strings"

	. "github.com/onsi/gomega"

	"github.com/sirupsen/logrus"
)

// CleanDatabase will clean the database of all content
func CleanDatabase() {
	connection := connect()
	defer connection.Close()
	tables := listTables(connection)

	joinedTables := strings.Join(tables, ", ")
	query := fmt.Sprintf("TRUNCATE %s", joinedTables)
	result, err := connection.Exec(query)
	Expect(err).To(BeNil())

	rowsAffected, _ := result.RowsAffected()
	logrus.
		WithField("affected", rowsAffected).
		Info("Cleaned Database")
}

func listTables(connection *sql.DB) []string {
	rows, err := connection.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	Expect(err).To(BeNil())

	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		Expect(err).To(BeNil())

		if tableName != "gorp_migrations" {
			tableNames = append(tableNames, tableName)
			logrus.WithField("table", tableName).Info("Cleaning table")
		}
	}

	return tableNames
}

package integration

import (
	// The postgres drivers
	_ "github.com/lib/pq"

	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/romanyx/polluter"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/suite"

	"gopkg.in/h2non/baloo.v3"
)

// Suite is the base for all Integration Tests
type Suite struct {
	suite.Suite

	dbConn     *sql.DB
	testClient *baloo.Client
}

// SetupTest will ensure that the Integration Suite is correctly defined
func (suite *Suite) SetupTest() {
	_, hasServiceURL := os.LookupEnv("SERVICE_URL")
	_, hasDatabaseURL := os.LookupEnv("DB_URL")

	if !hasServiceURL {
		suite.T().Skip("No Service URL provided")
		return
	}
	if !hasDatabaseURL {
		suite.T().Skip("No Database URL provided")
		return
	}

	dbConn, err := sql.Open("postgres", os.Getenv("DB_URL"))
	suite.Assert().NoError(err)
	suite.dbConn = dbConn

	url := os.Getenv("SERVICE_URL")
	suite.testClient = baloo.New(url)

	suite.CleanDatabase()
}

// TearDownTest will tidy up the tests after they finish
func (suite *Suite) TearDownTest() {
	suite.dbConn.Close()
}

// StartTest will start the Baloo Test Client for our tests
func (suite *Suite) StartTest() *baloo.Client {
	return suite.testClient
}

// ParseJSONToMap will convert the given JSON String to a Map for assertions
func (suite *Suite) ParseJSONToMap(input string) map[string]interface{} {
	var parsed map[string]interface{}

	err := json.Unmarshal([]byte(input), &parsed)
	suite.Assert().NoError(err)

	return parsed
}

// Seed will populate the database with data from the provided string
func (suite *Suite) Seed(input string) {
	p := polluter.New(polluter.PostgresEngine(suite.dbConn))

	err := p.Pollute(strings.NewReader(input))
	suite.Assert().NoError(err)
}

// CleanDatabase will clean the database of all content
func (suite *Suite) CleanDatabase() {
	tables := suite.listTables()

	joinedTables := strings.Join(tables, ", ")
	query := fmt.Sprintf("TRUNCATE %s", joinedTables)
	result, err := suite.dbConn.Exec(query)
	suite.Assert().NoError(err)

	rowsAffected, _ := result.RowsAffected()
	logrus.WithField("affected", rowsAffected).Info("Cleaned Database")
}

func (suite *Suite) listTables() []string {
	rows, err := suite.dbConn.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	suite.Assert().NoError(err)

	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		suite.Assert().NoError(err)

		if tableName != "gorp_migrations" {
			tableNames = append(tableNames, tableName)
			logrus.WithField("table", tableName).Debug("Selecting table to clean")
		}
	}

	return tableNames
}

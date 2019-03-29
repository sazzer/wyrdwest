package database

import (
	"strings"

	. "github.com/onsi/gomega"

	"github.com/romanyx/polluter"
)

// Seed will populate the database as described in the given string
func Seed(input string) {
	db := connect()
	defer db.Close()

	p := polluter.New(polluter.PostgresEngine(db))

	err := p.Pollute(strings.NewReader(input))
	Expect(err).To(BeNil())

}

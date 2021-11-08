package utils

import (
	"database/sql"
	"github.com/go-testfixtures/testfixtures/v3"
)

type Command interface {
	GetName() string
	Execute() error
}

type LoadFixturesCommand struct {
	name string
	db   *sql.DB
}

func (cmd *LoadFixturesCommand) GetName() string {
	return cmd.name
}

func (cmd *LoadFixturesCommand) Execute() error {
	fixtures, err := testfixtures.New(
		testfixtures.Database(cmd.db),
		testfixtures.Dialect("postgres"),
		testfixtures.Paths(
			"config/fixtures/categories.yml",
			"config/fixtures/products.yml",
		),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		return err
	}
	return fixtures.Load()
}

func NewLoadFixturesCommand(db *sql.DB, newName string) *LoadFixturesCommand {
	return &LoadFixturesCommand{name: newName, db: db}
}

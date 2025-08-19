package repository

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

type MigrateAction string

const (
	MigrateActionCreate MigrateAction = "create"
)

func Migrator(
	dsn string,
	action MigrateAction,
) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	_ = goose.SetDialect("mysql")
	migrationDir := "app/repository/migrations"
	switch action {
	case MigrateActionCreate:
		var input string
		fmt.Print("Enter migration name: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			return err
		}
		if err := goose.Create(db, migrationDir, input, "sql"); err != nil {
			return err
		}
	}

	return nil
}

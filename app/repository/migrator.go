package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/pressly/goose"
)

type MigrateAction string

const (
	MigrateActionUp     MigrateAction = "up"
	MigrateActionDown   MigrateAction = "down"
	MigrateActionStatus MigrateAction = "status"
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
	migrationDir := "repository/migrations"
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
	case MigrateActionUp:
		err := goose.Up(db, migrationDir)
		if err != nil {
			return err
		}
	case MigrateActionDown:
		var input string
		fmt.Print("Enter migration version you want to down:")
		_, err = fmt.Scanln(&input)
		if err != nil || input == "" {
			if err := goose.Down(db, migrationDir); err != nil {
				return err
			}
		} else {
			version, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				return err
			}
			if err := goose.DownTo(db, migrationDir, version); err != nil {
				return err
			}
		}
	case MigrateActionStatus:
		if err := goose.Status(db, migrationDir); err != nil {
			return err
		}
	default:
		return errors.New("unknown migrator action up|down|status")
	}

	return nil
}

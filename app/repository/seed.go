package repository

import (
	"database/sql"
	_ "embed"
	"log"
	"strings"
)

//go:embed seeds/seed.sql
var seedSQL string

func Seed(
	dsn string,
) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	queries := strings.Split(seedSQL, ";")
	for _, q := range queries {
		q = strings.TrimSpace(q)
		if q == "" {
			continue
		}
		_, err := db.Exec(q)
		if err != nil {
			log.Fatalf("exec query failed: %v\nSQL: %s", err, q)
		}
	}

	return nil
}

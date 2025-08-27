package main

import (
	"blog-api/repository"
	"database/sql"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

var migratorCommand = &cobra.Command{
	Use:   "migrate",
	Short: "migrate",
	Long: `Manage database migrations using Goose. 
Available options:
- create: Create a new migrate file
- up: Apply all migrations
- down: Rollback the last migration
- status: Show migration status
`,
	Args: cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		return repository.Migrator(os.Getenv("MYSQL_DSN"), repository.MigrateAction(args[0]))
	},
}

var dbSeedCommand = &cobra.Command{
	Use:   "seed",
	Short: "seed",
	Long:  "",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
		if err != nil {
			return err
		}
		defer db.Close()

		return repository.Seed(db)
	},
}

func main() {
	rootCmd.AddCommand(migratorCommand)
	rootCmd.AddCommand(dbSeedCommand)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

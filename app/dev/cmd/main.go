package main

import (
	"blog-api/config"
	"blog-api/repository"
	"database/sql"

	"github.com/spf13/cobra"
)

var cfg config.Config

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
		return repository.Migrator(cfg.Database.Dsn, repository.MigrateAction(args[0]))
	},
}

var dbSeedCommand = &cobra.Command{
	Use:   "seed",
	Short: "seed",
	Long:  "",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := sql.Open("mysql", cfg.Database.Dsn)
		if err != nil {
			return err
		}
		defer db.Close()

		return repository.Seed(db)
	},
}

func main() {
	configPath := "/config/config.yml"
	con, err := config.NewConfig(configPath)
	cfg = con
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(migratorCommand)
	rootCmd.AddCommand(dbSeedCommand)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

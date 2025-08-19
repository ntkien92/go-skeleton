package main

import (
	"blog-api/repository"
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

func main() {
	rootCmd.AddCommand(migratorCommand)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

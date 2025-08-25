package cmd

import "os"

func (server *ApiServer) loadEnv() []string {
	var errors []string

	server.mysqlDsn = os.Getenv("MYSQL_DSN")
	if server.mysqlDsn == "" {
		errors = append(errors, "MYSQL_DSN not set")
	}

	return errors
}

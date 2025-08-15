package cmd

import "os"

func (server *ApiServer) loadEnv() []string {
	var errors []string

	server.apiRunType = os.Getenv("API_RUN_TYPE")
	if server.apiRunType == "" {
		errors = append(errors, "API_RUN_TYPE not set")
	}

	server.runtimeEnv = os.Getenv("RUNTIME_ENV")
	if server.runtimeEnv == "" {
		errors = append(errors, "RUNTIME_ENV not set")
	}

	server.apiRunType = os.Getenv("API_RUN_TYPE")
	if server.apiRunType == "" {
		errors = append(errors, "API_RUN_TYPE not set")
	}

	server.mysqlDsn = os.Getenv("MYSQL_DSN")
	if server.mysqlDsn == "" {
		errors = append(errors, "MYSQL_DSN not set")
	}

	return errors
}

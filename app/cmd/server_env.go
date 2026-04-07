package cmd

import "blog-api/config"

func (server *ApiServer) loadEnv() []string {
	var errors []string

	configPath := "./env/config.yml"
	config, err := config.NewConfig(configPath)
	if err != nil {
		errors = append(errors, err.Error())
	}

	server.config = config

	return errors
}

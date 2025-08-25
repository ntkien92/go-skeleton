package config

import "os"

type Config struct {
	BaseURL string
	Token   string
}

var Global *Config

func Init() {
	Global = &Config{
		BaseURL: getenv("BASE_URL", "http://localhost:8081/api"),
		Token:   "",
	}
}

func getenv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

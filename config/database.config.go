package config

import "github.com/joeltiago00/first-api-go-lang/helpers"

type DatabaseConfig struct {
	User     string
	Password string
	Database string
	Host     string
}

func DbConfig() DatabaseConfig {
	return DatabaseConfig{
		User:     helpers.Env("DATABASE_USER"),
		Password: helpers.Env("DATABASE_PASSWORD"),
		Database: helpers.Env("DATABASE_NAME"),
		Host:     helpers.Env("DATABASE_HOST"),
	}
}

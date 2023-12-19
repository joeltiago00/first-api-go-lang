package config

import "github.com/joeltiago00/first-api-go-lang/helpers"

type DatabaseConfig struct {
	Connection string
	User       string
	Password   string
	Database   string
	Host       string
	Port       string
}

func DbConfig() DatabaseConfig {
	return DatabaseConfig{
		Connection: helpers.Env("DATABASE_CONNECTION"),
		User:       helpers.Env("DATABASE_USER"),
		Password:   helpers.Env("DATABASE_PASSWORD"),
		Database:   helpers.Env("DATABASE_NAME"),
		Host:       helpers.Env("DATABASE_HOST"),
		Port:       helpers.Env("DATABASE_PORT"),
	}
}

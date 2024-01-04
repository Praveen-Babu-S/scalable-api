package config

import "log/slog"

type PostgresConfig struct {
	Host       string
	Port       string
	DbName     string
	DbUser     string
	DbPassword string
}

type ServerConfig struct {
	ServerPort     string
	PostgresConfig PostgresConfig
	Logger         *slog.Logger
}

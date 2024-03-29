package main

import (
	"flag"

	"github.com/Praveen-Babu-S/scalable-api/models/config"
	"github.com/Praveen-Babu-S/scalable-api/pkg/server"
	"github.com/Praveen-Babu-S/scalable-api/utils/logger"
)

const (
	dbUsername = "APP_DB_USER"
	dbPassword = "APP_DB_PASSWORD"
)

func main() {

	dbHost := flag.String("dbHost", "localhost", "hostname on which db service is hosted")
	dbPort := flag.String("dbPort", "5432", "application db port")
	dbName := flag.String("dbName", "postgres", "application db name")
	serverPort := flag.String("serverPort", "9000", "application server port")

	flag.Parse()

	// Create logger instance for structured logging
	logger := logger.CreateLogger()

	serverConfig := config.ServerConfig{
		PostgresConfig: config.PostgresConfig{
			Host:       *dbHost,
			Port:       *dbPort,
			DbName:     *dbName,
			DbUser:     dbUsername,
			DbPassword: dbPassword,
		},
		ServerPort: *serverPort,
		Logger:     logger,
	}

	// Start server with the given configuration
	server.StartServer(serverConfig)

}

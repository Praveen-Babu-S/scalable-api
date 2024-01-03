package server

import (
	"log"

	"github.com/Praveen-Babu-S/scalable-api/models/config"
	"github.com/Praveen-Babu-S/scalable-api/pkg/db/connect"
	"github.com/Praveen-Babu-S/scalable-api/pkg/db/migrations"
	"github.com/Praveen-Babu-S/scalable-api/pkg/http"
)

func StartServer(serverConfig config.ServerConfig) {

	err := validateServerConfig(serverConfig)

	if err != nil {
		log.Fatalf("Invalid config:%s", err.Error())
	}

	db := connect.DBConnectionClient(serverConfig.PostgresConfig)

	// Run one time migrations to create required schema
	migrations.RunMigrations(db)

	//TODO: start REST API Handler here
	http.StartApiHandler(serverConfig.ServerPort)
}

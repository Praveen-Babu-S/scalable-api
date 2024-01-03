package server

import (
	"log"

	"github.com/Praveen-Babu-S/scalable-api/models/config"
	"github.com/Praveen-Babu-S/scalable-api/pkg/db/connect"
	"github.com/Praveen-Babu-S/scalable-api/pkg/db/migrations"
)

func StartServer(serverConfig config.PostgresConfig) {

	err := validateServerConfig(serverConfig)

	if err != nil {
		log.Fatalf("Invalid config:%s", err.Error())
	}

	db := connect.DBConnectionClient(serverConfig)

	// Run one time migrations to create required schema
	migrations.RunMigrations(db)
}

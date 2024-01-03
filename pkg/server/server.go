package server

import (
	"log"

	"github.com/Praveen-Babu-S/scalable-api/models/config"
	"github.com/Praveen-Babu-S/scalable-api/pkg/db/connect"
	"github.com/Praveen-Babu-S/scalable-api/pkg/db/migrations"
	auth "github.com/Praveen-Babu-S/scalable-api/pkg/handlers/authentication"
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

	// Initialise handlers
	authServer := auth.NewAuthImplementor(db)

	//TODO: start REST API Handler here
	http.StartApiHandler(serverConfig.ServerPort, authServer)
}

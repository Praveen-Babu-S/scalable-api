package server

import (
	"log"

	"github.com/Praveen-Babu-S/scalable-api/models/config"
	"github.com/Praveen-Babu-S/scalable-api/pkg/db/connect"
	"github.com/Praveen-Babu-S/scalable-api/pkg/db/migrations"
	auth "github.com/Praveen-Babu-S/scalable-api/pkg/handlers/authentication"
	"github.com/Praveen-Babu-S/scalable-api/pkg/handlers/domain/note"
	"github.com/Praveen-Babu-S/scalable-api/pkg/handlers/domain/search"
	"github.com/Praveen-Babu-S/scalable-api/pkg/handlers/domain/share"
	"github.com/Praveen-Babu-S/scalable-api/pkg/http"
)

func StartServer(serverConfig config.ServerConfig) {

	// Initialise logger
	logger := serverConfig.Logger

	err := validateServerConfig(serverConfig)

	if err != nil {
		logger.Debug("invalid serverConfig", "err", err.Error())
		log.Fatalf("Invalid config:%s", err.Error())
	}

	db := connect.DBConnectionClient(serverConfig.PostgresConfig, logger)

	// Run one time migrations to create required schema
	migrations.RunMigrations(db, logger)

	// Initialise handlers
	authServer := auth.NewAuthImplementor(db, logger)
	noteServer := note.NewNoteImplementor(db, logger)
	shareServer := share.NewShareImplementor(db)
	searchServer := search.NewSearchImplementor(db)

	// Start api handler
	http.StartApiHandler(serverConfig.ServerPort, authServer, noteServer, shareServer, searchServer)

	logger.Info("----- Server Running -----")
}

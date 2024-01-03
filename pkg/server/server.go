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

	err := validateServerConfig(serverConfig)

	if err != nil {
		log.Fatalf("Invalid config:%s", err.Error())
	}

	db := connect.DBConnectionClient(serverConfig.PostgresConfig)

	// Run one time migrations to create required schema
	migrations.RunMigrations(db)

	// Initialise handlers
	authServer := auth.NewAuthImplementor(db)
	noteServer := note.NewNoteImplementor(db)
	shareServer := share.NewShareImplementor(db)
	searchServer := search.NewSearchImplementor(db)

	// Start api handler
	http.StartApiHandler(serverConfig.ServerPort, authServer, noteServer, shareServer, searchServer)
}

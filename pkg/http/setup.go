package http

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/Praveen-Babu-S/scalable-api/pkg/handlers/authentication"
	"github.com/Praveen-Babu-S/scalable-api/pkg/handlers/domain/note"
	"github.com/Praveen-Babu-S/scalable-api/pkg/handlers/domain/search"
	"github.com/Praveen-Babu-S/scalable-api/pkg/handlers/domain/share"
	"github.com/gorilla/mux"
)

func StartApiHandler(serverPort string, authServer *auth.AuthServer, noteServer *note.NoteServer, shareServer *share.ShareServer, searchServer *search.SearchServer) {
	// Initialize Gorilla Mux router
	router := mux.NewRouter()

	// Authentication Endpoints
	router.HandleFunc("/api/auth/signup", authServer.SignupHandler).Methods("POST")
	router.HandleFunc("/api/auth/login", authServer.LoginHandler).Methods("POST")

	// Note Api Endpoints
	noteRouter := router.PathPrefix("/api/notes").Subrouter()
	noteRouter.Use(auth.AuthenticateMiddleware)
	router.HandleFunc("", noteServer.GetNotesHandler).Methods("GET")
	router.HandleFunc("/{id}", noteServer.GetNoteByIDHandler).Methods("GET")
	router.HandleFunc("", noteServer.CreateNoteHandler).Methods("POST")
	router.HandleFunc("/{id}", noteServer.UpdateNoteHandler).Methods("PUT")
	router.HandleFunc("/{id}", noteServer.DeleteNoteHandler).Methods("DELETE")

	// Share Api Endpoint
	router.HandleFunc("/{id}/share", shareServer.ShareNoteHandler).Methods("POST")

	// Search Endpoint
	router.HandleFunc("/api/search", searchServer.SearchNotesHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprint(":"+serverPort), router))
}

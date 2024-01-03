package http

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/Praveen-Babu-S/scalable-api/pkg/handlers/authentication"
	"github.com/gorilla/mux"
)

func StartApiHandler(serverPort string, authServer *auth.AuthServer) {
	// Initialize Gorilla Mux router
	router := mux.NewRouter()

	// Authentication Endpoints
	router.HandleFunc("/api/auth/signup", authServer.SignupHandler).Methods("POST")
	router.HandleFunc("/api/auth/login", authServer.LoginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprint(":"+serverPort), router))
}

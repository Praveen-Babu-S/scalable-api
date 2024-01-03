package main

import (
	"flag"

	"github.com/Praveen-Babu-S/scalable-api/models/config"
	"github.com/Praveen-Babu-S/scalable-api/pkg/server"
)

const (
	dbUsername = "APP_DB_USER"
	dbPassword = "APP_DB_PASSWORD"
)

func main() {

	dbHost := flag.String("dbHost", "localhost", "hostname on which db service is hosted")
	dbPort := flag.String("dbPort", "5432", "application db port")
	dbName := flag.String("dbName", "postgres", "application db name")

	flag.Parse()

	postgresConfig := config.PostgresConfig{
		Host:       *dbHost,
		Port:       *dbPort,
		DbName:     *dbName,
		DbUser:     dbUsername,
		DbPassword: dbPassword,
	}

	// Start server with the given configuration
	server.StartServer(postgresConfig)

	// // Initialize Gorilla Mux router
	// r := mux.NewRouter()

	// // Authentication Endpoints
	// r.HandleFunc("/api/auth/signup", SignupHandler).Methods("POST")
	// r.HandleFunc("/api/auth/login", LoginHandler).Methods("POST")

	// // Note Endpoints
	// r.HandleFunc("/api/notes", GetNotesHandler).Methods("GET")
	// r.HandleFunc("/api/notes/{id}", GetNoteByIDHandler).Methods("GET")
	// r.HandleFunc("/api/notes", CreateNoteHandler).Methods("POST")
	// r.HandleFunc("/api/notes/{id}", UpdateNoteHandler).Methods("PUT")
	// r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	// r.HandleFunc("/api/notes/{id}/share", ShareNoteHandler).Methods("POST")

	// // Search Endpoint
	// r.HandleFunc("/api/search", SearchNotesHandler).Methods("GET")

	// // Start the server
	// log.Fatal(http.ListenAndServe(":8080", r))
}

// // SignupHandler handles user registration.
// func SignupHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implement user registration logic and store user in the database
// }

// // LoginHandler handles user login.
// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implement user login logic and issue an access token
// }

// // GetNotesHandler gets a list of all notes for the authenticated user.
// func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implement logic to fetch all notes for the authenticated user
// }

// // GetNoteByIDHandler gets a note by ID for the authenticated user.
// func GetNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implement logic to fetch a note by ID for the authenticated user
// }

// // CreateNoteHandler creates a new note for the authenticated user.
// func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implement logic to create a new note for the authenticated user
// }

// // UpdateNoteHandler updates an existing note by ID for the authenticated user.
// func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implement logic to update an existing note by ID for the authenticated user
// }

// // DeleteNoteHandler deletes a note by ID for the authenticated user.
// func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implement logic to delete a note by ID for the authenticated user
// }

// // ShareNoteHandler shares a note with another user for the authenticated user.
// func ShareNoteHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implement logic to share a note with another user for the authenticated user
// }

// // SearchNotesHandler searches for notes based on keywords for the authenticated user.
// func SearchNotesHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implement logic to search for notes based on keywords for the authenticated user
// }

// // Helper function to respond with JSON
// func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
// 	response, err := json.Marshal(payload)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Internal Server Error"))
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }

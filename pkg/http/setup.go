package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartApiHandler(serverPort string) {
	// Initialize Gorilla Mux router
	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprint(":"+serverPort), router))
}

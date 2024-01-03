package search

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

type SearchImplementor interface {
	SearchNotesHandler(w http.ResponseWriter, r *http.Request)
}

type SearchServer struct {
	db *gorm.DB
}

func NewSearchImplementor(db *gorm.DB) *SearchServer {
	if db == nil {
		log.Fatalln("Invalid DB object, db is nil")
	}
	return &SearchServer{
		db: db,
	}
}

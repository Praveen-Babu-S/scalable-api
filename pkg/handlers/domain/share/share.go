package share

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

type ShareImplementor interface {
	ShareNoteHandler(w http.ResponseWriter, r *http.Request)
}

type ShareServer struct {
	db *gorm.DB
}

func NewShareImplementor(db *gorm.DB) *ShareServer {
	if db == nil {
		log.Fatalln("Invalid DB object, db is nil")
	}
	return &ShareServer{
		db: db,
	}
}

package note

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

type NoteImplementor interface {
	GetNotesHandler(w http.ResponseWriter, r *http.Request)
	GetNoteByIDHandler(w http.ResponseWriter, r *http.Request)
	CreateNoteHandler(w http.ResponseWriter, r *http.Request)
	UpdateNoteHandler(w http.ResponseWriter, r *http.Request)
	DeleteNoteHandler(w http.ResponseWriter, r *http.Request)
}

type NoteServer struct {
	db *gorm.DB
}

func NewNoteImplementor(db *gorm.DB) *NoteServer {
	if db == nil {
		log.Fatalln("Invalid DB object, db is nil")
	}
	return &NoteServer{
		db: db,
	}
}

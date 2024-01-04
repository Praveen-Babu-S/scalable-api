package note

import (
	"log"
	"log/slog"
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
	db     *gorm.DB
	logger *slog.Logger
}

func NewNoteImplementor(db *gorm.DB, logger *slog.Logger) *NoteServer {
	if db == nil {
		log.Fatalln("Invalid DB object, db is nil")
	}
	return &NoteServer{
		db:     db,
		logger: logger,
	}
}

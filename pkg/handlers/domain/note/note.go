package domain

import "net/http"

type CRUDImplementor interface {
	GetNotesHandler(w http.ResponseWriter, r *http.Request)
	GetNoteByIDHandler(w http.ResponseWriter, r *http.Request)
	CreateNoteHandler(w http.ResponseWriter, r *http.Request)
	UpdateNoteHandler(w http.ResponseWriter, r *http.Request)
	DeleteNoteHandler(w http.ResponseWriter, r *http.Request)
}

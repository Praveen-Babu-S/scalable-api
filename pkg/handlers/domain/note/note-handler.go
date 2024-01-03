package note

import "net/http"

func (s *NoteServer) GetNotesHandler(w http.ResponseWriter, r *http.Request)    {}
func (s *NoteServer) GetNoteByIDHandler(w http.ResponseWriter, r *http.Request) {}
func (s *NoteServer) CreateNoteHandler(w http.ResponseWriter, r *http.Request)  {}
func (s *NoteServer) UpdateNoteHandler(w http.ResponseWriter, r *http.Request)  {}
func (s *NoteServer) DeleteNoteHandler(w http.ResponseWriter, r *http.Request)  {}

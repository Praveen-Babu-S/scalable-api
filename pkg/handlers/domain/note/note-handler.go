package note

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Praveen-Babu-S/scalable-api/models/dbmodels"
	"github.com/Praveen-Babu-S/scalable-api/pkg/common"
	"github.com/gorilla/mux"
)

// Get list of all notes for an authenticated user
func (s *NoteServer) GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	userID := common.GetUserIDFromContext(r.Context())
	if userID == 0 {
		s.logger.Info("unauthorised access to GETNotes", "userId", userID)
		common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var notes []dbmodels.Note
	if err := s.db.Debug().Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		s.logger.Debug("unable to fetch notes list", "err", err.Error(), "userId", userID)
	}
	if len(notes) == 0 {
		s.logger.Debug("empty notes list found", "userId", userID)
		common.RespondWithJSON(w, http.StatusOK, "notes list is empty!")
	}
	s.logger.Info("fetched notes list", "notesList", notes, "userId", userID)
	common.RespondWithJSON(w, http.StatusOK, notes)
}

// Get a Note by its Id
func (s *NoteServer) GetNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := common.GetUserIDFromContext(r.Context())
	if userID == 0 {
		s.logger.Info("unauthorised access to GETNotesById", "userId", userID)
		common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	params := mux.Vars(r)
	noteID, err := strconv.Atoi(params["id"])
	if err != nil {
		s.logger.Debug("invalid nodeID", "userId", userID, "noteId", noteID)
		common.RespondWithError(w, http.StatusBadRequest, "Invalid noteID")
		return
	}

	var note dbmodels.Note
	result := s.db.Where("id = ? AND user_id = ?", noteID, userID).First(&note)
	if result.Error != nil {
		s.logger.Debug("unabel to fetch note", "userId", userID, "noteId", noteID, "err", result.Error.Error())
		common.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}
	s.logger.Info("fetched notes list by id", "note", note, "userId", userID)
	common.RespondWithJSON(w, http.StatusOK, note)
}

// Create Note for an authenticated user
func (s *NoteServer) CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	userID := common.GetUserIDFromContext(r.Context())
	if userID == 0 {
		s.logger.Info("unauthorised access to CreateNote", "userId", userID)
		common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var note dbmodels.Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&note); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	note.UserID = userID
	if err := s.db.Create(&note).Error; err != nil {
		log.Println("unable to create note:", err.Error())
		common.RespondWithError(w, http.StatusBadRequest, "Failed to create note")
		return
	}

	s.logger.Info("created notes successfully", "note", note, "userId", userID)
	common.RespondWithJSON(w, http.StatusCreated, note)
}
func (s *NoteServer) UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	userID := common.GetUserIDFromContext(r.Context())
	if userID == 0 {
		s.logger.Info("unauthorised access to UpdateNote", "userId", userID)
		common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	params := mux.Vars(r)
	noteID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("invalid nodeID:", err.Error())
		common.RespondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	var updatedNote dbmodels.Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedNote); err != nil {
		s.logger.Debug("unable to fetch note to update",
			"userId", userID, "noteId", noteID, "err", err.Error(), "method", "PUT")
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check if the note exists and belongs to the authenticated user
	var existingNote dbmodels.Note
	result := s.db.Where("id = ? AND user_id = ?", noteID, userID).First(&existingNote)
	if result.Error != nil {
		s.logger.Debug("unable to fetch note to update",
			"userId", userID, "noteId", noteID, "err", err.Error(), "method", "PUT")
		common.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	// Update the note
	existingNote.Title = updatedNote.Title
	existingNote.Content = updatedNote.Content
	if err := s.db.Save(&existingNote).Error; err != nil {
		s.logger.Debug("unable to update note",
			"userId", userID, "noteId", noteID, "err", err.Error(), "method", "PUT")
		common.RespondWithError(w, http.StatusBadRequest, "Failed to update note")
		return
	}

	s.logger.Info("updated notes successfully", "note", updatedNote, "userId", userID)
	common.RespondWithJSON(w, http.StatusOK, existingNote)
}

// DeleteNoteHandler deletes a note by ID for the authenticated user.
func (s *NoteServer) DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	userID := common.GetUserIDFromContext(r.Context())
	if userID == 0 {
		common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	params := mux.Vars(r)
	noteID, err := strconv.Atoi(params["id"])
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	// Check if the note exists and belongs to the authenticated user
	var note dbmodels.Note
	result := s.db.Where("id = ? AND user_id = ?", noteID, userID).First(&note)
	if result.Error != nil {
		common.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	// Delete the note
	s.db.Delete(&note)

	common.RespondWithJSON(w, http.StatusOK, "Note deleted successfully")
}

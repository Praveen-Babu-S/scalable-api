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
		common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var notes []dbmodels.Note
	if err := s.db.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		log.Println("unable to fetch user details:", err.Error())
	}
	if len(notes) == 0 {
		common.RespondWithJSON(w, http.StatusOK, "notes list is empty!")
	}
	common.RespondWithJSON(w, http.StatusOK, notes)
}

// Get a Note by its Id
func (s *NoteServer) GetNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := common.GetUserIDFromContext(r.Context())
	if userID == 0 {
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

	var note dbmodels.Note
	result := s.db.Where("id = ? AND user_id = ?", noteID, userID).First(&note)
	if result.Error != nil {
		log.Println("unable to fetch note details:", err.Error())
		common.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, note)
}

// Create Note for an authenticated user
func (s *NoteServer) CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	userID := common.GetUserIDFromContext(r.Context())
	if userID == 0 {
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

	common.RespondWithJSON(w, http.StatusCreated, note)
}
func (s *NoteServer) UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	userID := common.GetUserIDFromContext(r.Context())
	if userID == 0 {
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
		log.Println("unable to update note:", err.Error())
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check if the note exists and belongs to the authenticated user
	var existingNote dbmodels.Note
	result := s.db.Where("id = ? AND user_id = ?", noteID, userID).First(&existingNote)
	if result.Error != nil {
		log.Println("unable to fetch note:", err.Error())
		common.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	// Update the note
	existingNote.Title = updatedNote.Title
	existingNote.Content = updatedNote.Content
	if err := s.db.Save(&existingNote).Error; err != nil {
		log.Println("unable to update note:", err.Error())
		common.RespondWithError(w, http.StatusBadRequest, "Failed to update note")
		return
	}

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

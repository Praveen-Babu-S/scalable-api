package share

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Praveen-Babu-S/scalable-api/models/dbmodels"
	"github.com/Praveen-Babu-S/scalable-api/pkg/common"
	"github.com/gorilla/mux"
)

func (s *ShareServer) ShareNoteHandler(w http.ResponseWriter, r *http.Request) {
	userID := common.GetUserIDFromContext(r.Context())
	if userID == 0 {
		common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	params := mux.Vars(r)
	noteID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("invalid note ID:", err.Error())
		common.RespondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	var shareRequest struct {
		RecipientUsername string `json:"recipient_username"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&shareRequest); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check if the note exists and belongs to the authenticated user
	var note dbmodels.Note
	result := s.db.Where("id = ? AND user_id = ?", noteID, userID).First(&note)
	if result.Error != nil {
		common.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	// Check if the recipient user exists
	var recipientUser dbmodels.User
	result = s.db.Where("username = ?", shareRequest.RecipientUsername).First(&recipientUser)
	if result.Error != nil {
		log.Println("unable to fetch user:", err.Error())
		common.RespondWithError(w, http.StatusNotFound, "Recipient user not found")
		return
	}

	// Check if the note is already shared with the recipient
	var existingShare Share
	result = s.db.Where("note_id = ? AND recipient_user_id = ?", note.ID, recipientUser.ID).First(&existingShare)
	if result.Error == nil {
		common.RespondWithError(w, http.StatusConflict, "Note already shared with the recipient")
		return
	}

	// Share the note
	share := Share{
		NoteID:           note.ID,
		RecipientUserID:  recipientUser.ID,
		SharingUserID:    userID,
		SharingTimestamp: time.Now(),
	}
	s.db.Create(&share)

	common.RespondWithJSON(w, http.StatusCreated, share)
}

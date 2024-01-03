package search

import (
	"log"
	"net/http"

	"github.com/Praveen-Babu-S/scalable-api/models/dbmodels"
	"github.com/Praveen-Babu-S/scalable-api/pkg/common"
)

func (s *SearchServer) SearchNotesHandler(w http.ResponseWriter, r *http.Request) {
	userID := common.GetUserIDFromContext(r.Context())
	if userID == 0 {
		common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		common.RespondWithError(w, http.StatusBadRequest, "Search query cannot be empty")
		return
	}

	// Perform a simple search for notes containing the query string
	var notes []dbmodels.Note
	if err := s.db.Debug().Where("user_id = ? AND search_vector @@ plainto_tsquery('english', ?)", userID, query).Find(&notes).Error; err != nil {
		log.Println("unable to search notes:", err.Error())
	}

	common.RespondWithJSON(w, http.StatusOK, notes)
}

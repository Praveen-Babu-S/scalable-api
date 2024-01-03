package search

import (
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
	s.db.Where("user_id = ? AND (title ILIKE ? OR content ILIKE ?)", userID, "%"+query+"%", "%"+query+"%").Find(&notes)

	common.RespondWithJSON(w, http.StatusOK, notes)
}

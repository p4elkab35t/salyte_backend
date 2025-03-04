package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/logic"
)

type ReactionHandler struct {
	reactionLogic *logic.ReactionService
}

func NewreactionLogicHandler(reactionLogic *logic.ReactionService) *ReactionHandler {
	return &ReactionHandler{reactionLogic: reactionLogic}
}

// GetReaction handles GET requests to retrieve a reaction.
func (h *ReactionHandler) GetReactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	messageID := r.URL.Query().Get("messageID")
	userID := r.URL.Query().Get("userID")

	reaction, err := h.reactionLogic.GetReactionsByMessageID(ctx, uuid.MustParse(messageID), uuid.MustParse(userID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reaction)
}

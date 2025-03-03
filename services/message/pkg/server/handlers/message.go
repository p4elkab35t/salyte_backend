package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/logic"
)

type MessageHandler struct {
	messageLogic *logic.MessageService
}

func NewMessageHandler(messageLogic *logic.MessageService) *MessageHandler {
	return &MessageHandler{messageLogic: messageLogic}
}

// GetChat handles GET requests to retrieve a chat.
func (h *MessageHandler) GetMessagesByChatID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chatID := r.URL.Query().Get("chatID")
	userID := r.URL.Query().Get("userID")

	// crate type for json parisng and parse limit and offset from body
	var requestBody struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	if requestBody.Limit == 0 {
		requestBody.Limit = 10
	}

	if requestBody.Offset == 0 {
		requestBody.Offset = 0
	}

	messages, err := h.messageLogic.GetMessagesByChatID(ctx, uuid.MustParse(chatID), requestBody.Limit, requestBody.Offset, uuid.MustParse(userID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

// get Unread messages
func (h *MessageHandler) GetUnreadMessages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.URL.Query().Get("userID")

	messages, err := h.messageLogic.GetUnreadMessages(ctx, uuid.MustParse(userID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

// delete all messages
func (h *MessageHandler) DeleteAllMessagesByChatID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.URL.Query().Get("userID")
	chatID := r.URL.Query().Get("chatID")

	err := h.messageLogic.DeleteAllMessagesInChat(ctx, uuid.MustParse(chatID), uuid.MustParse(userID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "all messages deleted"})
}

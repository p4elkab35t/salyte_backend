package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/logic"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/models"
)

type ChatHandler struct {
	chatLogic *logic.ChatService
}

func NewChatHandler(chatLogic *logic.ChatService) *ChatHandler {
	return &ChatHandler{chatLogic: chatLogic}
}

// GetChat handles GET requests to retrieve a chat.
func (h *ChatHandler) GetChat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chatID := r.URL.Query().Get("chatID")

	chat, err := h.chatLogic.GetChatByID(ctx, uuid.MustParse(chatID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chat)
}

// CreateChat handles POST requests to create a new chat.
func (h *ChatHandler) CreateChat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var requestBody struct {
		Chat  models.Chat `json:"chat"`
		Users []string    `json:"users"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	var users []*models.ChatMember

	for _, userID := range requestBody.Users {
		users = append(users, &models.ChatMember{UserID: uuid.MustParse(userID)})
	}

	createdChat, err := h.chatLogic.CreateChat(ctx, users, &requestBody.Chat)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdChat)
}

func (h *ChatHandler) GetAllChats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.URL.Query().Get("userID")

	chats, err := h.chatLogic.GetChatsByUserID(ctx, uuid.MustParse(userID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chats)
}

func (h *ChatHandler) AddUserToChat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chatID := r.URL.Query().Get("chatID")
	userID := r.URL.Query().Get("userID")
	addedUserID := r.URL.Query().Get("addedUserID")

	if err := h.chatLogic.AddUserToChat(ctx, uuid.MustParse(chatID), uuid.MustParse(userID), uuid.MustParse(addedUserID)); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user added to chat"})

}

func (h *ChatHandler) RemoveUserFromChat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chatID := r.URL.Query().Get("chatID")
	userID := r.URL.Query().Get("userID")
	removedUserID := r.URL.Query().Get("removedUserID")

	if err := h.chatLogic.RemoveUserFromChat(ctx, uuid.MustParse(chatID), uuid.MustParse(userID), uuid.MustParse(removedUserID)); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user removed from chat"})
}

// get chat members
func (h *ChatHandler) GetChatMembers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chatID := r.URL.Query().Get("chatID")
	userID := r.URL.Query().Get("userID")

	members, err := h.chatLogic.GetChatMembers(ctx, uuid.MustParse(chatID), uuid.MustParse(userID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(members)
}

// get chat messages
func (h *ChatHandler) GetChatByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chatID := r.URL.Query().Get("chatID")

	messages, err := h.chatLogic.GetChatByID(ctx, uuid.MustParse(chatID))
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

func (h *ChatHandler) GetChatByMembers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.URL.Query().Get("userID")
	memberID := r.URL.Query().Get("memberID")

	if userID == "" || memberID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	chat, err := h.chatLogic.GetChatByTwoUsers(ctx, uuid.MustParse(userID), uuid.MustParse(memberID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chat)
}

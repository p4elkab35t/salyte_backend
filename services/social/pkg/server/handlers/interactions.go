package handlers

import (
	// "context"
	"encoding/json"
	"fmt"
	"net/http"

	// "strings"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/logic"
)

type InteractionHandler struct {
	interactionLogic *logic.InteractionService
}

func NewInteractionHandler(interactionLogic *logic.InteractionService) *InteractionHandler {
	return &InteractionHandler{interactionLogic: interactionLogic}
}

// GetPostComments handles GET requests to retrieve comments for a specific post.
func (h *InteractionHandler) GetPostComments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := r.URL.Query().Get("postID")

	comments, err := h.interactionLogic.GetPostComments(ctx, postID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

// GetPostLikes handles GET requests to retrieve likes for a specific post.
func (h *InteractionHandler) GetPostLikes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := r.URL.Query().Get("postID")

	likes, err := h.interactionLogic.GetPostLikes(ctx, postID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(likes)
}

// LikePost handles POST requests to like a post.
func (h *InteractionHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := r.URL.Query().Get("postID")
	fmt.Println(postID)
	// Extract user ID from the request (e.g., from a JWT token or session)
	userID := ctx.Value("profileID").(string) // Here we are using the profileID but its called UserID need refactor
	fmt.Println(userID)
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "user ID is required"})
		return
	}

	if err := h.interactionLogic.LikePost(ctx, userID, postID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "post liked"})
}

// UnlikePost handles DELETE requests to unlike a post.
func (h *InteractionHandler) UnlikePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := r.URL.Query().Get("postID")
	fmt.Println(postID)
	// Extract user ID from the request (e.g., from a JWT token or session)
	userID := ctx.Value("profileID").(string) // Here we are using the profileID but its called UserID need refactor
	fmt.Println(userID)
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "user ID is required"})
		return
	}

	if err := h.interactionLogic.UnlikePost(ctx, userID, postID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "post unliked"})
}

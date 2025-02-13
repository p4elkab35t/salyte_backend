package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/logic"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
)

type PostHandler struct {
	postLogic *logic.PostService
}

func NewPostHandler(postLogic *logic.PostService) *PostHandler {
	return &PostHandler{postLogic: postLogic}
}

// GetPost handles GET requests to retrieve a post.
func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	postID := strings.TrimPrefix(r.URL.Path, "/social/post/")

	post, err := h.postLogic.GetPostByID(ctx, postID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

// CreatePost handles POST requests to create a new post.
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	createdPost, err := h.postLogic.CreatePost(ctx, &post)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}

// UpdatePost handles PUT requests to update a post.
func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	postID := strings.TrimPrefix(r.URL.Path, "/social/post/")

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	post.PostID = uuid.MustParse(postID)
	if err := h.postLogic.UpdatePost(ctx, &post); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "post updated"})
}

// DeletePost handles DELETE requests to delete a post.
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	postID := strings.TrimPrefix(r.URL.Path, "/social/post/")

	if err := h.postLogic.DeletePost(ctx, postID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "post deleted"})
}

// GetPostsByCommunity handles GET requests to retrieve posts by community.
func (h *PostHandler) GetPostsByCommunity(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	communityID := strings.TrimPrefix(r.URL.Path, "/social/post/community/")

	posts, err := h.postLogic.GetPostsByCommunityID(ctx, communityID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

// GetPostsByUser handles GET requests to retrieve posts by user.
func (h *PostHandler) GetPostsByUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userID := strings.TrimPrefix(r.URL.Path, "/social/post/user/")

	posts, err := h.postLogic.GetPostsByUserID(ctx, userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

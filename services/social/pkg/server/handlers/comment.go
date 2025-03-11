package handlers

import (
	// "context"
	"encoding/json"
	"net/http"

	// "strings"

	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/logic"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
)

type CommentHandler struct {
	commentLogic *logic.CommentService
}

func NewCommentHandler(commentLogic *logic.CommentService) *CommentHandler {
	return &CommentHandler{commentLogic: commentLogic}
}

// CreateComment handles POST requests to create a new comment.
func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	comment.ProfileID = uuid.MustParse(ctx.Value("profileID").(uuid.UUID).String())

	createdComment, err := h.commentLogic.CreateComment(ctx, &comment)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdComment)
}

// GetCommentByID handles GET requests to retrieve a comment by its ID.
func (h *CommentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	commentID := r.URL.Query().Get("commentID")

	comment, err := h.commentLogic.GetCommentByID(ctx, commentID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}

// UpdateComment handles PUT requests to update a comment.
func (h *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	commentID := r.URL.Query().Get("commentID")

	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	comment.ProfileID = uuid.MustParse(ctx.Value("profileID").(uuid.UUID).String())

	comment.CommentID = uuid.MustParse(commentID)
	if err := h.commentLogic.UpdateComment(ctx, &comment); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "comment updated"})
}

// DeleteComment handles DELETE requests to delete a comment.
func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	commentID := r.URL.Query().Get("commentID")

	if err := h.commentLogic.DeleteComment(ctx, commentID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "comment deleted"})
}

// GetCommentsByPostID handles GET requests to retrieve comments for a specific post.
func (h *CommentHandler) GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := r.URL.Query().Get("postID")

	if postID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "postID is required"})
		return
	}

	comments, err := h.commentLogic.GetCommentsByPostID(ctx, postID)
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

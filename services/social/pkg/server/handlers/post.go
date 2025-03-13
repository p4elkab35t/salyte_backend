package handlers

import (
	// "context"
	"encoding/json"
	"net/http"
	"strconv"

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
	ctx := r.Context()
	postID := r.URL.Query().Get("postID")
	var (
		posts []*models.Post
		err   error
	)

	if postID == "" {
		page := r.URL.Query().Get("page")
		limit := r.URL.Query().Get("limit")
		if page == "" || limit == "" {
			page, err := strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "invalid page number"})
				return
			}
			limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "invalid page number"})
				return
			}
			posts, err = h.postLogic.GetAllPosts(ctx, page, limit)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}
		} else {
			posts, err = h.postLogic.GetAllPosts(ctx, 0, 100)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
		return
	}

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
	ctx := r.Context()

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	profileID := uuid.MustParse(ctx.Value("profileID").(uuid.UUID).String())
	post.ProfileID = &profileID

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
	ctx := r.Context()
	postID := r.URL.Query().Get("postID")

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	profileID := uuid.MustParse(ctx.Value("profileID").(string))
	post.ProfileID = &profileID

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
	ctx := r.Context()
	postID := r.URL.Query().Get("postID")

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
	ctx := r.Context()
	communityID := r.URL.Query().Get("communityID")

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

func (h *PostHandler) GetPostsByUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.URL.Query().Get("profileID")

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

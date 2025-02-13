package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/logic"
)

type FollowHandler struct {
	followLogic *logic.FollowService
}

func NewFollowHandler(followLogic *logic.FollowService) *FollowHandler {
	return &FollowHandler{followLogic: followLogic}
}

// FollowProfile handles POST requests to follow a profile.
func (h *FollowHandler) FollowProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	profileID := strings.TrimPrefix(r.URL.Path, "/social/follow/")

	// Extract user ID from the request (e.g., from a JWT token or session)
	userID := r.Header.Get("User-ID")
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "user ID is required"})
		return
	}

	if err := h.followLogic.FollowProfile(ctx, userID, profileID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "profile followed"})
}

// UnfollowProfile handles DELETE requests to unfollow a profile.
func (h *FollowHandler) UnfollowProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	profileID := strings.TrimPrefix(r.URL.Path, "/social/follow/")

	// Extract user ID from the request (e.g., from a JWT token or session)
	userID := r.Header.Get("User-ID")
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "user ID is required"})
		return
	}

	if err := h.followLogic.UnfollowProfile(ctx, userID, profileID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "profile unfollowed"})
}

// GetFollowers handles GET requests to retrieve a profile's followers.
func (h *FollowHandler) GetFollowers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	profileID := strings.TrimPrefix(r.URL.Path, "/social/followers/")

	followers, err := h.followLogic.GetProfileFollowers(ctx, profileID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(followers)
}

// GetFollowing handles GET requests to retrieve profiles a user is following.
func (h *FollowHandler) GetFollowing(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	profileID := strings.TrimPrefix(r.URL.Path, "/social/following/")

	following, err := h.followLogic.GetProfileFollowing(ctx, profileID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(following)
}

// GetFriends handles GET requests to retrieve a profile's friends.
func (h *FollowHandler) GetFriends(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	profileID := strings.TrimPrefix(r.URL.Path, "/social/friends/")

	friends, err := h.followLogic.GetFriends(ctx, profileID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(friends)
}

// GetFriendRequests handles GET requests to retrieve pending friend requests.
func (h *FollowHandler) GetFriendRequests(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	profileID := strings.TrimPrefix(r.URL.Path, "/social/friend-requests/")

	requests, err := h.followLogic.GetFriendsRequests(ctx, profileID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(requests)
}

// make friend and unfriends handlers

func (h *FollowHandler) MakeFriend(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	profileID := strings.TrimPrefix(r.URL.Path, "/social/friend/")

	// Extract user ID from the request (e.g., from a JWT token or session)
	userID := r.Header.Get("User-ID")
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "user ID is required"})
		return
	}

	if err := h.followLogic.MakeFriends(ctx, userID, profileID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "profile followed"})
}

func (h *FollowHandler) Unfriend(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	profileID := strings.TrimPrefix(r.URL.Path, "/social/friend/")

	// Extract user ID from the request (e.g., from a JWT token or session)
	userID := r.Header.Get("User-ID")
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "user ID is required"})
		return
	}

	if err := h.followLogic.Unfriend(ctx, userID, profileID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "profile unfollowed"})
}

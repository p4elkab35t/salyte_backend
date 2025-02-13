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

type CommunityHandler struct {
	communityLogic *logic.CommunityService
}

func NewCommunityHandler(communityLogic *logic.CommunityService) *CommunityHandler {
	return &CommunityHandler{communityLogic: communityLogic}
}

// GetCommunity handles GET requests to retrieve a community.
func (h *CommunityHandler) GetCommunity(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	communityID := strings.TrimPrefix(r.URL.Path, "/social/community/")

	community, err := h.communityLogic.GetCommunityByID(ctx, communityID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(community)
}

// CreateCommunity handles POST requests to create a new community.
func (h *CommunityHandler) CreateCommunity(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var community models.Community
	if err := json.NewDecoder(r.Body).Decode(&community); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	createdCommunity, err := h.communityLogic.CreateCommunity(ctx, &community)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCommunity)
}

// UpdateCommunity handles PUT requests to update a community.
func (h *CommunityHandler) UpdateCommunity(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	communityID := strings.TrimPrefix(r.URL.Path, "/social/community/")

	var community models.Community
	if err := json.NewDecoder(r.Body).Decode(&community); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	community.CommunityID = uuid.MustParse(communityID)
	if err := h.communityLogic.UpdateCommunity(ctx, &community); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "community updated"})
}

// GetCommunityMembers handles GET requests to retrieve community members.
func (h *CommunityHandler) GetCommunityMembers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	communityID := strings.TrimPrefix(r.URL.Path, "/social/community/")

	members, err := h.communityLogic.GetCommunityFollowers(ctx, communityID)
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

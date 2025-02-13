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

type ProfileHandler struct {
	profileLogic *logic.ProfileService
}

func NewProfileHandler(profileLogic *logic.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileLogic: profileLogic}
}

// GetProfile handles GET requests to retrieve a profile.
func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userID := strings.TrimPrefix(r.URL.Path, "/social/profile/")

	profile, err := h.profileLogic.GetProfileByID(ctx, userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

// UpdateProfile handles PUT requests to update a profile.
func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userID := strings.TrimPrefix(r.URL.Path, "/social/profile/")

	var profile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	profile.ProfileID = uuid.MustParse(userID)
	if err := h.profileLogic.UpdateProfile(ctx, &profile); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "profile updated"})
}

// GetProfileSettings handles GET requests to retrieve profile settings.
func (h *ProfileHandler) GetProfileSettings(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userID := strings.TrimPrefix(r.URL.Path, "/social/profile/")

	settings, err := h.profileLogic.GetSettings(ctx, userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(settings)
}

// UpdateProfileSettings handles PUT requests to update profile settings.
func (h *ProfileHandler) UpdateProfileSettings(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userID := strings.TrimPrefix(r.URL.Path, "/social/profile/")

	var settings models.Setting
	if err := json.NewDecoder(r.Body).Decode(&settings); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	settings.ProfileID = uuid.MustParse(userID)
	if err := h.profileLogic.UpdateSettings(ctx, &settings); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "settings updated"})
}

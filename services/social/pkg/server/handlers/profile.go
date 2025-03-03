package handlers

import (
	// "context"
	"encoding/json"
	"fmt"
	"net/http"

	// "os/user"

	// "strings"

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

func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer r.Body.Close()

	var user struct {
		UserID string `json:"userID"`
		Email  string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "invalid request body"}`, http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	profileModel := models.Profile{
		UserID:   uuid.MustParse(user.UserID),
		Username: user.Email,
	}

	profile, err := h.profileLogic.CreateProfile(ctx, &profileModel)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]string{"profileID": profile.ProfileID.String()}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetProfile handles GET requests to retrieve a profile.
func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.URL.Query().Get("userID")

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
	ctx := r.Context()
	userID := r.URL.Query().Get("userID")

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
	ctx := r.Context()
	userID := r.URL.Query().Get("userID")

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
	ctx := r.Context()
	userID := r.URL.Query().Get("userID")

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

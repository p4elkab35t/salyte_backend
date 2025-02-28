package handlers

import (
	// "encoding/json"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/logic"
)

type TokenHandler struct {
	authLogic *logic.AuthLogicService
}

func NewTokenHandler(authLogic *logic.AuthLogicService) *TokenHandler {
	return &TokenHandler{authLogic}
}

func (h *TokenHandler) VerifyToken(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.Method == "GET" {

		// Get the token session cookie
		token := r.Header.Get("Authorization")
		user_id := r.Header.Get("user_id")

		result, err := h.authLogic.VerifySession(ctx, token, user_id)

		w.Header().Set("Content-Type", "application/json")
		if result {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
		}
	}
	// Create a map to store the response
}

func (h *TokenHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.Method == "GET" {
		// Parse the form data
		token := r.Header.Get("Authorization")

		err := h.authLogic.SignOut(ctx, token)

		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(true)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Success")
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
		}
	}
	// Create a map to store the response
}

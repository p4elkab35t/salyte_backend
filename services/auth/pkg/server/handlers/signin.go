package handlers

import (
	// "encoding/json"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/logic"
)

type SignInHandler struct {
	authLogic *logic.AuthLogicService
}

func NewSignInHandler(authLogic *logic.AuthLogicService) *SignInHandler {
	return &SignInHandler{authLogic}
}

func (h *SignInHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.Method == "GET" {
		// Get the token session cookie
		token := r.Header.Get("Authorization")

		result, err := h.authLogic.CheckToken(ctx, token)

		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(result)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Success")
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
		}
	}
	if r.Method == "POST" {
		// Parse the form data
		r.ParseMultipartForm(10 << 20)
		// Get the form data
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		result, err := h.authLogic.SignIn(ctx, email, password)

		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(result)
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

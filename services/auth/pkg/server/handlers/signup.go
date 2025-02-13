package handlers

import (
	// "encoding/json"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/logic"
)

type SignUpHandler struct {
	authLogic *logic.AuthLogicService
}

func NewSignUpHandler(authLogic *logic.AuthLogicService) *SignInHandler {
	return &SignInHandler{authLogic}
}

func (h *SignUpHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.Method == "POST" {
		// Parse the form data
		r.ParseMultipartForm(10 << 20)
		// Get the form data
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		result, err := h.authLogic.SignUp(ctx, email, password)

		if err != nil {
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

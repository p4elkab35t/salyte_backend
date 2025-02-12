package handlers

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/logic"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse the form data
		r.ParseForm()
		// Get the form data
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		result := logic.Auth(&email, &password)

		if result {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Success")
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
		}
	}
	// Create a map to store the response
}

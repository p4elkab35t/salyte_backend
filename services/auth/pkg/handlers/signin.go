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
		r.ParseMultipartForm(10 << 20)
		// Get the form data
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		result := logic.Signin(&email, &password)

		if result {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Correct endpoint"))
			fmt.Fprintf(w, "Success")
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized, but correct endpoint"))
			fmt.Fprintf(w, "Unauthorized")
		}
	}
	// Create a map to store the response
}

package auth

import (
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/handlers"
	// "fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/auth/signin", handlers.Signin)
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("Server is running on port 8080")
}

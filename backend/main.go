package main

import (
	"blogsite/backend/pkg/authentication"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/backend/signup", authentication.Signup)
	http.HandleFunc("/backend/login", authentication.Login)

	// Start the server on port 8080
	fmt.Println("Backend server listening on PORT: 8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"blogsite/backend/pkg/authentication"
	cfg "blogsite/backend/pkg/config"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/backend/signup", authentication.Signup)
	http.HandleFunc("/backend/login", authentication.Login)

	// Start the server
	fmt.Printf("Backend server listening on PORT: %s", cfg.Port)
	http.ListenAndServe(cfg.Port, nil)
}

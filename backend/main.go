package main

import (
	"blogsite/backend/pkg/authentication"
	cfg "blogsite/backend/pkg/config"
	"blogsite/backend/pkg/posts"
	"fmt"
	"net/http"
)

func main() {
	// Authentication handlers
	http.HandleFunc("/backend/signup", authentication.Signup)
	http.HandleFunc("/backend/login", authentication.Login)

	// Post handlers
	http.HandleFunc("/backend/post/create", posts.CreatePost)
	http.HandleFunc("/backend/post/list", posts.FetchPosts)

	// Start the server
	fmt.Printf("Backend server listening on PORT: %s", cfg.Port)
	http.ListenAndServe(cfg.Port, nil)
}

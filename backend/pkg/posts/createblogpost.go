package posts

import (
	"blogsite/backend/middleware"
	database "blogsite/backend/pkg/db"
	"encoding/json"
	"net/http"
)

type Post struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	var post Post
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		http.Error(w, "Error decoding to JSON!", http.StatusConflict)
		return
	}

	db, _ := database.EstablishDBCon()
	defer db.Close()

	_, err = db.Exec("INSERT INTO posts (username, title, message) VALUES (?,?,?)", post.Username, post.Title, post.Message)
	if err != nil {
		http.Error(w, "Error", http.StatusConflict)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

package posts

import (
	"blogsite/backend/middleware"
	database "blogsite/backend/pkg/db"
	"encoding/json"
	"net/http"
)

type ListPosts struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

func FetchPosts(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	db, _ := database.EstablishDBCon()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		http.Error(w, "Error in getting list!", http.StatusConflict)
		return
	}

	var postList ListPosts
	var listOfPosts []ListPosts

	for rows.Next() {
		var username, title, message string

		err := rows.Scan(&username, &title, &message)
		if err != nil {
			http.Error(w, "Error in getting list!", http.StatusConflict)
			return
		}

		postList.Username = username
		postList.Title = title
		postList.Message = message

		listOfPosts = append(listOfPosts, postList)
	}

	jsonList, err := json.Marshal(&listOfPosts)
	if err != nil {
		http.Error(w, "Error in getting list!", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonList)
}

package authentication

import (
	"blogsite/backend/middleware"
	"blogsite/backend/pkg/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON data from the request body
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	db, _ := db.EstablishDBCon()
	defer db.Close()

	if db.QueryRow("SELECT username FROM users WHERE username=?", user.Username).Scan() == sql.ErrNoRows {
		_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
		if err != nil {
			log.Printf("Failed to insert user: %s", err)
		}

		fmt.Print("Created new user!")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Print("Username already exists!")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusLocked)
	}
}

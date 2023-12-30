package authentication

import (
	"blogsite/backend/middleware"
	"blogsite/backend/pkg/db"
	"database/sql"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		return
	} else {
		var fetchedPassword string
		db.QueryRow("SELECT password FROM users WHERE username=?", user.Username).Scan(&fetchedPassword)

		if fetchedPassword == user.Password {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user.PasswordHash, err = utils.HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Error while hashing the password", http.StatusInternalServerError)
			return
		}

		err = models.CreateUser(db, user)
		if err != nil {
			http.Error(w, "Error while creating the user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Account successfully created"))
	}

}

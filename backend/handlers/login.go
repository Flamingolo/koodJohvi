package handlers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		var storedPassword string
		var userID int

		err := h.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", email).Scan(&userID, &storedPassword)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		session, _ := h.Store.Get(r, "session-name")
		session.Values["authenticated"] = true
		session.Values["userID"] = userID
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

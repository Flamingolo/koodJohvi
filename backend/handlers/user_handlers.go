package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rtf/backend/models"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RegisterHandler called")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.CreateUser(db, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Identifier string `json:"identifier"` // can be email or nickname
		Password   string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := models.AuthenticateUser(db, credentials.Identifier, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	session, err := models.CreateSession(db, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":    user,
		"session": session,
	})
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	sessionID := r.Header.Get("Authorization")
	if sessionID == "" {
		http.Error(w, "Missing session ID", http.StatusUnauthorized)
		return
	}

	session, err := models.ValidateSession(db, sessionID)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	if session.UserID != id {
		http.Error(w, "Unauthorized access", http.StatusForbidden)
		return
	}

	err = models.UpdateSessionActivity(db, sessionID)
	if err != nil {
		http.Error(w, "Unable to update session activity", http.StatusInternalServerError)
		return
	}

	user, err := models.GetUserByID(db, id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := r.Header.Get("Authorization")
	if sessionID == "" {
		http.Error(w, "Missing session ID", http.StatusUnauthorized)
		return
	}

	err := models.DeleteSession(db, sessionID)
	if err != nil {
		http.Error(w, "Unable to log out", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

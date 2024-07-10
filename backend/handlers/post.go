package handlers

import (
	"encoding/json"
	"net/http"
	"rtf/backend/models"
	"time"
)

func (h *Handlers) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	session, _ := h.Store.Get(r, "session-name")
	userID := session.Values["user_id"].(int)

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	post.UserID = userID
	post.CreatedAt = time.Now()

	_, err = h.DB.Exec("INSERT INTO posts (user_id, title, content, created_at) VALUES (?, ?, ?, ?)", post.UserID, post.Title, post.Content, post.CreatedAt)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

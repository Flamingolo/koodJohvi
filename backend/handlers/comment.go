package handlers

import (
	"encoding/json"
	"net/http"
	"rtf/backend/models"
	"time"
)

func (h *Handlers) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	session, _ := h.Store.Get(r, "session-name")
	userID := session.Values["user_id"].(int)

	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	comment.UserID = userID
	comment.CreatedAt = time.Now()

	_, err = h.DB.Exec("INSERT INTO comments (post_id, user_id, content, created_at) VALUES (?, ?, ?, ?)", comment.PostID, comment.UserID, comment.Content, comment.CreatedAt)
	if err != nil {
		http.Error(w, "Error creating comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

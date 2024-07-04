package handlers

import (
	"encoding/json"
	"net/http"
	"rtf/backend/models"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	json.NewDecoder(r.Body).Decode(&comment)

	userID, err := GetLoggedInUsedID(r)
	if err != nil {
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return
	}
	comment.UserID = userID

	if err := models.CreateComment(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LikeComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, _ := strconv.Atoi(vars["id"])

	comment, err := models.GetCommentByID(commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	comment.Likes++
	if err := models.UpdateCommentLikes(commentID, comment.Likes, comment.Dislikes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DislikeComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, _ := strconv.Atoi(vars["id"])

	comment, err := models.GetCommentByID(commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	comment.Dislikes++
	if err := models.UpdateCommentLikes(commentID, comment.Likes, comment.Dislikes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

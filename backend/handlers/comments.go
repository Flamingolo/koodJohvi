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

func GetComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, _ := strconv.Atoi(vars["postID"])

	comments, err := models.GetCommentByID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(comments)

}

func LikeComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, _ := strconv.Atoi(vars["id"])

	err := models.UpdateCommentLikes(commentID, 1, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DislikeComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, _ := strconv.Atoi(vars["id"])

	err := models.UpdateCommentLikes(commentID, 0, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

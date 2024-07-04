package handlers

import (
	"encoding/json"
	"net/http"
	"rtf/backend/models"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	userID, err := GetLoggedInUsedID(r)
	if err != nil {
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return
	}
	post.UserID = userID

	if err := models.CreatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, _ := strconv.Atoi(vars["id"])

	post, err := models.GetPostByID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	post.Likes++
	if err := models.UpdatePostLikes(postID, post.Likes, post.Dislikes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DislikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, _ := strconv.Atoi(vars["id"])

	post, err := models.GetPostByID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	post.Dislikes++
	if err := models.UpdatePostLikes(postID, post.Likes, post.Dislikes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

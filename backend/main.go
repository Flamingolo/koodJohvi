package main

import (
	"log"
	"net/http"
	"rtf/backend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Router
	r := mux.NewRouter()

	// Handlers
	r.HandleFunc("/api/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/api/posts/{id}/like", handlers.LikePost).Methods("POST")
	r.HandleFunc("/api/posts/{id}/dislike", handlers.DislikePost).Methods("POST")
	r.HandleFunc("/api/comments", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/api/comments/{id}/like", handlers.LikeComment).Methods("POST")
	r.HandleFunc("/api/comments/{id}/dislike", handlers.DislikeComment).Methods("POST")
	r.HandleFunc("/ws", handlers.WebSocketHandler)

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	// Server
	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"log"
	"net/http"
	"rtf/backend/database"
	"rtf/backend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize db
	database.InitializeDatabase()
	// Router
	r := mux.NewRouter()

	// Handlers
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}/like", handlers.LikePost).Methods("POST")
	r.HandleFunc("/posts/{id}/dislike", handlers.DislikePost).Methods("POST")
	r.HandleFunc("/comments", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/comments/{id}/like", handlers.LikeComment).Methods("POST")
	r.HandleFunc("/comments/{id}/dislike", handlers.DislikeComment).Methods("POST")
	r.HandleFunc("/ws", handlers.WebSocketHandler)

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	// Server
	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

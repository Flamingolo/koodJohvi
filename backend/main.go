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

	// Authenticate middleware
	authMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := handlers.GetLoggedInUsedID(r)
			if err != nil {
				http.Error(w, "User not logged in", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// Public Routes
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")

	// Protected Routes
	r.Handle("/posts", authMiddleware(http.HandlerFunc(handlers.CreatePost))).Methods("POST")
	r.Handle("/posts/{id}/like", authMiddleware(http.HandlerFunc(handlers.LikePost))).Methods("POST")
	r.Handle("/posts/{id}/dislike", authMiddleware(http.HandlerFunc(handlers.DislikePost))).Methods("POST")
	r.Handle("/comments", authMiddleware(http.HandlerFunc(handlers.CreateComment))).Methods("POST")
	r.Handle("/comments/{id}/like", authMiddleware(http.HandlerFunc(handlers.LikeComment))).Methods("POST")
	r.Handle("/comments/{id}/dislike", authMiddleware(http.HandlerFunc(handlers.DislikeComment))).Methods("POST")
	r.HandleFunc("/ws", handlers.WebSocketHandler)

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	// Server
	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

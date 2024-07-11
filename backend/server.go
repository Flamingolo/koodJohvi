package main

import (
	"database/sql"
	"log"
	"net/http"
	"rtf/backend/handlers"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

func main() {
	var err error
	db, err := sql.Open("sqlite3", "./database/rtfdb.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handlers.SetDB(db)
	router := mux.NewRouter()

	// Static files
	fs := http.FileServer(http.Dir("../frontend"))
	router.PathPrefix("/").Handler(fs)

	// User endpoints
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Post endpoints
	router.HandleFunc("/posts", handlers.CreatePostHandler).Methods("POST")
	router.HandleFunc("/posts/{id}", handlers.GetPostHandler).Methods("GET")
	router.HandleFunc("/posts", handlers.GetAllPostHandler).Methods("GET")

	// Comment endpoints
	router.HandleFunc("/comments", handlers.CreateCommentHandler).Methods("POST")
	router.HandleFunc("/comments{id}", handlers.GetCommentHandler).Methods("GET")
	router.HandleFunc("/comments/post/{postId}", handlers.GetCommentsByPostHandler).Methods("GET")

	// Message endpoints
	router.HandleFunc("/messages", handlers.CreateMessageHandler).Methods("POST")
	router.HandleFunc("/messages/{id}", handlers.GetMessageHandler).Methods("GET")
	router.HandleFunc("/messages/user/{userId}", handlers.GetMessagesByUserHandler).Methods("GET")

	// Category endpoints
	router.HandleFunc("/categories", handlers.CreateCategoryHandler).Methods("POST")
	router.HandleFunc("/categories/{id}", handlers.GetCategoryHandler).Methods("GET")
	router.HandleFunc("/categories", handlers.GetAllCategoriesHandler).Methods("GET")

	log.Println("Starting server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

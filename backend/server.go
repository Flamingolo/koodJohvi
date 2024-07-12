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

	// Api router
	apiRouter := router.PathPrefix("/api").Subrouter()

	// User endpoints
	apiRouter.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	apiRouter.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	apiRouter.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	apiRouter.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")

	// Post endpoints
	apiRouter.HandleFunc("/posts", handlers.CreatePostHandler).Methods("POST")
	apiRouter.HandleFunc("/posts/{id}", handlers.GetPostHandler).Methods("GET")
	apiRouter.HandleFunc("/posts", handlers.GetAllPostHandler).Methods("GET")
	apiRouter.HandleFunc("/posts/{id}/score", handlers.UpdatePostScoreHandler).Methods("PUT")

	// Comment endpoints
	apiRouter.HandleFunc("/comments", handlers.CreateCommentHandler).Methods("POST")
	apiRouter.HandleFunc("/comments{id}", handlers.GetCommentHandler).Methods("GET")
	apiRouter.HandleFunc("/comments/post/{postId}", handlers.GetCommentsByPostHandler).Methods("GET")
	apiRouter.HandleFunc("/comments/{id}/score", handlers.UpdateCommentScoreHandler).Methods("PUT")

	// Message endpoints
	apiRouter.HandleFunc("/messages", handlers.CreateMessageHandler).Methods("POST")
	apiRouter.HandleFunc("/messages/{id}", handlers.GetMessageHandler).Methods("GET")
	apiRouter.HandleFunc("/messages/user/{userId}", handlers.GetMessagesByUserHandler).Methods("GET")
	apiRouter.HandleFunc("/messages/{id}/read", handlers.MarkMessageAsRead).Methods("PUT")

	// Category endpoints
	apiRouter.HandleFunc("/categories", handlers.CreateCategoryHandler).Methods("POST")
	apiRouter.HandleFunc("/categories/{id}", handlers.GetCategoryHandler).Methods("GET")
	apiRouter.HandleFunc("/categories", handlers.GetAllCategoriesHandler).Methods("GET")
	apiRouter.HandleFunc("/categories/{id}", handlers.DeleteCategoryHandler).Methods("DELETE")
	apiRouter.HandleFunc("/categories/{id}", handlers.UpdateCategoryHandler).Methods("PUT")

	log.Println("Starting server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

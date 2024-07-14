package main

import (
	"log"
	"net/http"
	"rtf/backend/database"
	"rtf/backend/handlers"
	"rtf/backend/router" // Import your custom router package

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize db
	db := database.InitDB()
	defer db.Close()

	// Set the database for the handlers
	handlers.SetDB(db)

	handlers.SetDB(db)

	r := &router.Router{}
	r.SetAPIPrefix("api")

	// User endpoints
	r.NewRoute("/register", handlers.RegisterHandler, "POST")
	r.NewRoute("/login", handlers.LoginHandler, "POST")
	r.NewRoute("/users/{id}", handlers.GetUserHandler, "GET")
	r.NewRoute("/logout", handlers.LogoutHandler, "POST")

	// Test endpoint
	r.NewRoute("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Test endpoint called")
		w.Write([]byte("Test successful"))
	}, "GET")

	// Post endpoints
	r.NewRoute("/posts", handlers.CreatePostHandler, "POST", "GET")
	r.NewRoute("/posts/{id}", handlers.GetPostHandler, "GET")
	r.NewRoute("/posts", handlers.GetAllPostHandler, "GET")
	r.NewRoute("/posts/{id}/score", handlers.UpdatePostScoreHandler, "PUT")

	// Comment endpoints
	r.NewRoute("/comments", handlers.CreateCommentHandler, "POST")
	r.NewRoute("/comments/{id}", handlers.GetCommentHandler, "GET")
	r.NewRoute("/comments/post/{postId}", handlers.GetCommentsByPostHandler, "GET")
	r.NewRoute("/comments/{id}/score", handlers.UpdateCommentScoreHandler, "PUT")

	// Message endpoints
	r.NewRoute("/messages", handlers.CreateMessageHandler, "POST")
	r.NewRoute("/messages/{id}", handlers.GetMessageHandler, "GET")
	r.NewRoute("/messages/user/{userId}", handlers.GetMessagesByUserHandler, "GET")
	r.NewRoute("/messages/{id}/read", handlers.MarkMessageAsRead, "PUT")

	// Category endpoints
	r.NewRoute("/categories", handlers.CreateCategoryHandler, "POST")
	r.NewRoute("/categories/{id}", handlers.GetCategoryHandler, "GET")
	r.NewRoute("/categories", handlers.GetAllCategoriesHandler, "GET")
	r.NewRoute("/categories/{id}", handlers.DeleteCategoryHandler, "DELETE")
	r.NewRoute("/categories/{id}", handlers.UpdateCategoryHandler, "PUT")

	log.Println("Starting server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

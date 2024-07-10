package backend

import (
	"database/sql"
	"net/http"
	"rtf/backend/handlers"

	"github.com/gorilla/sessions"
)

type Routes struct {
	Handlers *handlers.Handlers
}

func NewRoutes(db *sql.DB, store *sessions.CookieStore) *Routes {
	handlers := handlers.NewHandlers(db, store)
	return &Routes{
		Handlers: handlers,
	}
}

func (rt *Routes) InitializeRoutes() {
	h := rt.Handlers

	// Public routes
	http.HandleFunc("/login", h.LoginHandler())
	http.HandleFunc("/logout", h.LogOutHandler())

	// Protected routes
	http.Handle("/createPost", h.IsAuthenticated(http.HandlerFunc(h.CreatePostHandler)))
	http.Handle("/createComment", h.IsAuthenticated(http.HandlerFunc(h.CreateCommentHandler)))

	// Apply session middleware to all routes
	http.Handle("/", h.WithSession(http.HandlerFunc(h.HomeLander)))
}

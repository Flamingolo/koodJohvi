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

	// Middleware
	http.HandleFunc("/", h.WithSession(h.ServeSPA))

	// API Routes
	http.HandleFunc("/login", h.LoginHandler())
	http.HandleFunc("/logout", h.LogOutHandler())
	http.Handle("/createPost", h.IsAuthenticated(h.CreatePostHandler))
	http.Handle("/createComment", h.IsAuthenticated(h.CreateCommentHandler))
}

package handlers

import (
	"database/sql"

	"github.com/gorilla/sessions"
)

type Handlers struct {
	DB    *sql.DB
	Store *sessions.CookieStore
}

func NewHandlers(db *sql.DB, store *sessions.CookieStore) *Handlers {
	return &Handlers{
		DB:    db,
		Store: store,
	}
}

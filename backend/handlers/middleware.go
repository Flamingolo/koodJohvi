package handlers

import (
	"net/http"
)

func (h *Handlers) WithSession(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := h.Store.Get(r, "session-name")
		if session.IsNew {
			session.Values["authenticated"] = false
			session.Save(r, w)
		}
		next(w, r)
	}
}

func (h *Handlers) IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := h.Store.Get(r, "session-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next(w, r)
	}
}

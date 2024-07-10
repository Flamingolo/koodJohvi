package handlers

import "net/http"

func (h *Handlers) LogOutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := h.Store.Get(r, "session-name")
		session.Values["authenticated"] = false
		session.Options.MaxAge = -1
		session.Save(r, w)

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

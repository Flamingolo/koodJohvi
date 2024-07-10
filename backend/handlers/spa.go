package handlers

import "net/http"

func (h *Handlers) ServeSPA(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}

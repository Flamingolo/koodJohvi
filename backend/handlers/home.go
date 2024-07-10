package handlers

import "net/http"

func (h *Handlers) HomeLander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page!"))
}

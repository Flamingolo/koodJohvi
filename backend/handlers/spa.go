package handlers

import (
	"net/http"
	"path/filepath"
)

func (h *Handlers) ServeSPA(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}

	filePath := filepath.Join("frontend", path)
	http.ServeFile(w, r, filePath)
}

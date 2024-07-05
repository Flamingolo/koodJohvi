package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./backend/database/forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

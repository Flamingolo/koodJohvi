package main

import (
	"database/sql"
	"log"
	"main/handlers"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Handlers
	http.HandleFunc("/backend/register", handlers.RegisterHandler(db))

	// Server
	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}

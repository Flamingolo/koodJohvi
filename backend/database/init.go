package database

import (
	"database/sql"
	"log"
)

func InitializeDatabase() {
	db, err := sql.Open("sqlite3", "./backend/database/forum.db")
	if err != nil {
		log.Fatalf("Failed to open database %v", err)
	}
	defer db.Close()

	log.Println("Database initialized successfully")
}

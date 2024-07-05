package database

import (
	"database/sql"
	"io/ioutil"
	"log"
)

func InitializeDatabase(){
	db, err := sql.Open("sqlite3", "./backend/database/forum.db")
	if err != nil {
		log.Fatalf("Failed to open database %v", err)
	}
	defer db.Close()

	schema, err := ioutil.ReadFile("./backend/database/schema.sql")
	if err != nil {
		log.Fatalf("Failed to read schema file %v", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatalf("Failed to execute schema %v", err)
	}

	log.Println("Database initialized successfully")
}
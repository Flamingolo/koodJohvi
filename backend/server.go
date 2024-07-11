package main

import (
	"database/sql"
	"log"

	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	var err error
	db, err := sql.Open("sqlite3", "./database/rtfdb.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	// User endpoints
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Post endpoints
	router.HandleFunc("/posts", handlers.CreatePostHandler).Methods("POST")
	router.HandleFunc("/posts/{id}", handlers.GetPostHandler).Methods("GET")
	router.HandleFunc("/posts", handlers.GetAllPostHandler).Methods("GET")

	// Comment endpoints

	// Message endpoints

	// Category endpoints

}

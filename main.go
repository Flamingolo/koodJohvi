package main

import (
	"log"
	"net/http"
	"rtf/backend"
	"rtf/backend/database"

	"github.com/gorilla/sessions"
)

func main() {
	// Initialize the database
	db := database.InitDB()
	defer db.Close()

	// Initialize session
	store := sessions.NewCookieStore([]byte("RTF-key"))

	// Initialize Routes
	routes := backend.NewRoutes(db, store)
	routes.InitializeRoutes()

	// Start the server
	log.Println("Starting the server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

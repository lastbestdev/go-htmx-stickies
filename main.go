package main

import (
	"log"
	"net/http"
	"os"
	"stickies/internal/components"
	"stickies/internal/db"
	"stickies/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// initialize db connection
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := "5432"

	if username == "" || password == "" {
		log.Fatal("Database credentials not set in environment variables")
	}
	db.InitDB(username, password, port)

	// create router
	router := mux.NewRouter()

	// add handlers
	router.HandleFunc("/", handlers.ComponentRenderer(components.RenderMenu()))
	router.HandleFunc("/boards", handlers.BoardsHandler)
	router.HandleFunc("/boards/{id}", handlers.BoardsDetailHandler)
	router.HandleFunc("/stickies", handlers.StickiesHandler)
	router.HandleFunc("/stickies/{id}", handlers.StickiesDetailHandler)
	router.HandleFunc("/forms/{form_name}", handlers.FormsHandler)

	// serve static assets (including HTMX dist)
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// start webserver on port 8080
	log.Printf("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

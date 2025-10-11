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
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")

	if username == "" || password == "" || port == "" {
		log.Fatal("Database credentials not set in environment variables")
	}
	db.InitDB(username, password, port)

	// create a sticky board (TODO: remove)
	// board_id := services.CreateBoard("Goat Debate")

	// create 3 test stickies (TODO: remove)
	// services.CreateSticky("Kobe Bean Bryant. Duh", 1)
	// services.CreateSticky("MJ has 6 rings", 1)
	// services.CreateSticky("LeBron has 50k pts! Cmon bro im not even glazing just hear me out his impact on the game is immeasurable and there will never be another Lebron James. The COVID ring obviously counts dude don't be like that. If anything it counts more because of the hardship of playing during a global pandemic! Lets be serious here", 1)

	// create router
	router := mux.NewRouter()

	// BEGIN: add handlers
	router.HandleFunc("/", handlers.ComponentRenderer(components.RenderMenu()))

	router.HandleFunc("/boards", handlers.BoardsHandler)
	router.HandleFunc("/boards/{id}", handlers.BoardsHandler)
	router.HandleFunc("/stickies", handlers.StickiesHandler)
	router.HandleFunc("/stickies/{id}", handlers.StickiesHandler)
	router.HandleFunc("/forms/{form_name}", handlers.FormsHandler)

	// serve static assets (including HTMX dist)
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// END: add handlers

	// start webserver on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

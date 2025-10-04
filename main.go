package main

import (
	"log"
	"net/http"
	"stickies/internal/components"
	"stickies/internal/handlers"
	"stickies/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	// create a sticky board (TODO: remove)
	board := services.CreateBoard("Goat debate")

	// create 3 test stickies (TODO: remove)
	services.CreateSticky("Kobe is the goat", &board)
	services.CreateSticky("LeGoat is my goat", &board)
	services.CreateSticky("MJ has 6 rings", &board)

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

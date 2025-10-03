package main

import (
	"log"
	"net/http"
	"stickies/internal/components"
	"stickies/internal/handlers"
	"stickies/internal/services"
)

func main() {
	// create a sticky board (TODO: remove)
	board := services.CreateBoard("Goat debate")

	// create 3 test stickies (TODO: remove)
	services.CreateSticky("Kobe is the goat", &board)
	services.CreateSticky("LeGoat is my goat", &board)
	services.CreateSticky("MJ has 6 rings", &board)

	// BEGIN: add handlers
	http.HandleFunc("/", handlers.ComponentRenderer(components.RenderMenu()))

	http.HandleFunc("/board", handlers.BoardHandler)
	http.HandleFunc("/sticky", handlers.StickyHandler)

	http.HandleFunc("/forms/create-board", handlers.ComponentRenderer(components.RenderCreateBoardForm()))
	http.HandleFunc("/forms/add-sticky", handlers.ComponentRenderer(components.RenderAddStickyForm(board)))

	// serve static assets (including HTMX)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// END: add handlers

	// start webserver on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

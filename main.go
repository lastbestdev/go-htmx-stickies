package main

import (
	"context"
	"log"
	"net/http"
	"stickies/internal/components"
	"stickies/internal/handlers"
	"stickies/internal/services"
	"strconv"
)

func showBoard(board services.Board) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		component := components.RenderBoard(board)
		component.Render(context.Background(), w)
	}
}

func getAddStickyForm(board services.Board) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		component := components.RenderAddStickyForm(board)
		component.Render(context.Background(), w)
	}
}

func postSticky(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	boardId, err := strconv.Atoi(r.FormValue("board_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	board := services.GetBoard(boardId)
	if board == nil {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	sticky := services.CreateSticky(content, board)

	// render new sticky and return component
	component := components.RenderSticky(sticky)
	component.Render(context.Background(), w)
}

func main() {
	// create a sticky board (TODO: remove)
	board := services.CreateBoard("Goat debate")

	// create 3 test stickies (TODO: remove)
	services.CreateSticky("Kobe is the goat", &board)
	services.CreateSticky("LeGoat is my goat", &board)
	services.CreateSticky("MJ has 6 rings", &board)

	// BEGIN: add handlers
	http.HandleFunc("/board", handlers.ComponentRenderer(components.RenderBoard(board)))

	http.HandleFunc("/forms/add-sticky", handlers.ComponentRenderer(components.RenderAddStickyForm(board)))

	http.HandleFunc("/sticky", postSticky)

	// serve static assets (including HTMX)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// END: add handlers

	// start webserver on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

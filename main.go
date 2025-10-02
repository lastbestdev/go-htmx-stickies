package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
)

func showBoard(board Board) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		component := renderBoard(board)
		component.Render(context.Background(), w)
	}
}

func getAddStickyForm(board Board) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		component := renderAddStickyForm(board)
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

	sticky := createSticky(content, 0, 0)

	board := getBoard(boardId)
	if board == nil {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	addSticky(board, sticky)

	// render new sticky and return component
	component := renderSticky(sticky)
	component.Render(context.Background(), w)
}

func main() {
	// create 3 test stickies (TODO: remove)
	sticky1 := createSticky("Kobe is the goat", 0, 0)
	sticky2 := createSticky("LeGoat is my goat", 20, 20)
	sticky3 := createSticky("MJ has 6 rings", 40, 40)

	// create a sticky board (TODO: remove)
	board := createBoard("Goat debate")

	// add stickies to board (TODO: remove)
	addSticky(&board, sticky1)
	addSticky(&board, sticky2)
	addSticky(&board, sticky3)

	// BEGIN: add handlers
	http.HandleFunc("/board", showBoard(board))

	http.HandleFunc("/forms/add-sticky", getAddStickyForm(board))

	http.HandleFunc("/sticky", postSticky)

	// serve static assets (including HTMX)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// END: add handlers

	// start webserver on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

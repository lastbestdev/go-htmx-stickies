package main

import (
	"context"
	"log"
	"net/http"
)

func getBoard(board Board) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		component := renderBoard(board)
		component.Render(context.Background(), w)
	}
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

	// board handlers
	http.HandleFunc("/board", getBoard(board))

	// start webserver on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

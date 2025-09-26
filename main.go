package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

// TODO: remove templ testing
func testTempl() {
	component := hello("TEST")
	component.Render(context.Background(), os.Stdout)
}

func getBoard(w http.ResponseWriter, r *http.Request) {
	board := createBoard("My Board")
	fmt.Fprintf(w, "Board ID=%d name: %s", board.Id, board.Name)
}

func getSticky(w http.ResponseWriter, r *http.Request) {
	sticky := createSticky("This is my sticky note", 0, 0)
	fmt.Fprintf(w, "Sticky ID=%d content: %s", sticky.Id, sticky.Content)
}

func main() {
	// board handlers
	http.HandleFunc("/board", getBoard)

	// sticky handlers
	http.HandleFunc("/sticky", getSticky)

	// start webserver on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

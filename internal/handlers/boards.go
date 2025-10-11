package handlers

import (
	"net/http"
	"stickies/internal/components"
	"stickies/internal/services"
	"strconv"

	"github.com/gorilla/mux"
)

func BoardsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getBoard(w, r)
	case http.MethodPost:
		createBoard(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getBoard(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	boardId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	board := services.GetBoard(boardId)

	if board == nil {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	component := components.RenderBoard(*board)
	ComponentRenderer(component)(w, r)
}

func createBoard(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	board_id := services.CreateBoard(name)

	// get new board
	board := services.GetBoard(board_id)

	// render new board component
	component := components.RenderBoard(*board)
	ComponentRenderer(component)(w, r)
}

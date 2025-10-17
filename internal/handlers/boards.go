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
		getBoards(w, r)
	case http.MethodPost:
		createBoard(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func BoardsDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Board ID is required", http.StatusBadRequest)
		return
	}

	boardId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid board ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getBoard(boardId, w, r)
	case http.MethodDelete:
		deleteBoard(boardId, w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getBoard(boardId int, w http.ResponseWriter, r *http.Request) {
	board := services.GetBoard(boardId)

	if board == nil {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	component := components.RenderBoard(*board)
	ComponentRenderer(component)(w, r)
}

func deleteBoard(boardId int, w http.ResponseWriter, r *http.Request) {
	services.DeleteBoard(boardId)
	w.WriteHeader(http.StatusNoContent)
}

func getBoards(w http.ResponseWriter, r *http.Request) {
	boards := services.GetBoards()

	component := components.RenderBoardsList(boards)
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

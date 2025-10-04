package handlers

import (
	"net/http"
	"stickies/internal/components"
	"stickies/internal/services"
	"strconv"
)

func StickiesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postSticky(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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

	// render new sticky component
	component := components.RenderSticky(sticky)
	ComponentRenderer(component)(w, r)
}

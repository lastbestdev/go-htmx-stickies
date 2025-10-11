package handlers

import (
	"net/http"
	"stickies/internal/components"
	"stickies/internal/services"
	"strconv"

	"github.com/gorilla/mux"
)

func StickiesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postSticky(w, r)
	case http.MethodDelete:
		deleteSticky(w, r)
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

	sticky_id := services.CreateSticky(content, boardId)

	// get new sticky from db
	sticky := services.GetSticky(sticky_id)

	// render new sticky component
	component := components.RenderSticky(*sticky)
	ComponentRenderer(component)(w, r)
}

func deleteSticky(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	stickyId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	services.DeleteSticky(stickyId)
}

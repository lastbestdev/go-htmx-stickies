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
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func StickiesDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Sticky ID is required", http.StatusBadRequest)
		return
	}

	stickyId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid sticky ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		deleteSticky(stickyId, w, r)
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
	color := r.FormValue("color")
	boardId, err := strconv.Atoi(r.FormValue("board_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sticky_id := services.CreateSticky(content, boardId, color)

	// get new sticky from db
	sticky := services.GetSticky(sticky_id)

	// render new sticky component
	component := components.RenderSticky(*sticky)
	ComponentRenderer(component)(w, r)
}

func deleteSticky(stickyId int, w http.ResponseWriter, r *http.Request) {
	services.DeleteSticky(stickyId)
	w.WriteHeader(http.StatusNoContent)
}

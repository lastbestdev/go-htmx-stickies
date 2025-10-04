package handlers

import (
	"net/http"
	"stickies/internal/components"
	"strconv"

	"github.com/gorilla/mux"
)

func FormsHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["form_name"]

	switch name {
	case "create-board":
		ComponentRenderer(components.RenderCreateBoardForm())(w, r)
		return
	case "add-sticky":
		id := r.URL.Query().Get("board_id")
		if id == "" {
			http.Error(w, "board_id is required", http.StatusBadRequest)
			return
		}

		board_id, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid board_id", http.StatusBadRequest)
			return
		}

		ComponentRenderer(components.RenderAddStickyForm(board_id))(w, r)
		return
	default:
		http.Error(w, "Form not found", http.StatusNotFound)
		return
	}
}

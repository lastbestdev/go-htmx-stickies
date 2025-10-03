package handlers

import (
	"context"
	"net/http"
	"stickies/internal/components"
	"stickies/internal/services"

	"github.com/a-h/templ"
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

func ComponentRenderer(component templ.Component) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		component.Render(context.Background(), w)
	}
}

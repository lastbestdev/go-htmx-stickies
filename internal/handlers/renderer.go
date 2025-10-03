package handlers

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func ComponentRenderer(component templ.Component) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		component.Render(context.Background(), w)
	}
}

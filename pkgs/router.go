package pkgs

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Routes() *chi.Mux {
	mux := chi.NewRouter()
	appHandler := NewHandler()
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fs))
	mux.Get("/", appHandler.Home)
	return mux
}

package pkgs

import (
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.html")
}

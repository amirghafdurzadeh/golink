package redirect

import (
	"net/http"
)

type Handler interface {
	Redirect(w http.ResponseWriter, r *http.Request)
}

type handler struct{}

func NewHandler() Handler {
	return &handler{}
}

// unimplemented
func (h *handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	_ = code

	targetURL := "https://example.com"

	http.Redirect(w, r, targetURL, http.StatusFound)
}

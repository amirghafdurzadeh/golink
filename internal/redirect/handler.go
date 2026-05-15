package redirect

import (
	"net/http"

	"github.com/amirghafdurzadeh/golink/internal/link"
)

type Handler interface {
	Redirect(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	linkService link.Service
}

func NewHandler(linkService link.Service) Handler {
	return &handler{
		linkService: linkService,
	}
}

func (h *handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	l, err := h.linkService.Get(r.Context(), code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, l.TargetURL, http.StatusFound)
}

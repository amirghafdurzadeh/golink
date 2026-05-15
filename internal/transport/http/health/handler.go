package health

import (
	"net/http"

	"github.com/amirghafdurzadeh/golink/internal/health"
	"github.com/amirghafdurzadeh/golink/internal/httpx"
)

type Handler struct {
	service health.Service
}

func NewHandler(service health.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if err := h.service.Check(r.Context()); err != nil {
		httpx.WriteError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	httpx.WriteJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
	})
}

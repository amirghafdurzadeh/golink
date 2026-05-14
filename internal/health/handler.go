package health

import (
	"context"
	"net/http"
	"time"

	"github.com/amirghafdurzadeh/golink/internal/httpx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	db *pgxpool.Pool
}

func NewHandler(db *pgxpool.Pool) Handler {
	return &handler{
		db: db,
	}
}

func (h *handler) Health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		r.Context(),
		2*time.Second,
	)
	defer cancel()

	err := h.db.Ping(ctx)
	if err != nil {

		httpx.WriteError(
			w,
			http.StatusServiceUnavailable,
			"database unavailable",
		)

		return
	}

	httpx.WriteJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
	})
}

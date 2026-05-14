package redirect

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler interface {
	Redirect(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	db *pgxpool.Pool
}

func NewHandler(db *pgxpool.Pool) Handler {
	return &handler{
		db: db,
	}
}

// unimplemented
func (h *handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	_ = code

	targetURL := "https://example.com"

	http.Redirect(w, r, targetURL, http.StatusFound)
}

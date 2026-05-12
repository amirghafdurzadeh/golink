package handler

import (
	"github.com/amirghafdurzadeh/golink/internal/db"
)

type Handler struct {
	db *db.Database
}

func New(db *db.Database) *Handler {
	return &Handler{
		db: db,
	}
}

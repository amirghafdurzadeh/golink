package link

import (
	"errors"
	"net/http"
	"strings"

	"github.com/amirghafdurzadeh/golink/internal/httpx"
	"github.com/amirghafdurzadeh/golink/internal/link"
	"github.com/amirghafdurzadeh/golink/internal/transport/http/helper"
)

type Handler struct {
	service link.Service
}

func NewHandler(service link.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateLinkRequest

	if err := httpx.DecodeJSON(r, &req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	l, err := h.service.Create(r.Context(), req.CustomCode, req.TargetURL)
	if err != nil {
		if errors.Is(err, link.ErrCodeAlreadyExists) {
			httpx.WriteError(w, http.StatusConflict, err.Error())
			return
		}

		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	shortURL := strings.TrimRight(h.service.GetBaseURL(), "/") + "/r/" + l.Code

	httpx.WriteJSON(w, http.StatusCreated, CreateLinkResponse{
		Code:      l.Code,
		ShortURL:  shortURL,
		TargetURL: l.TargetURL,
	})
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	code, err := helper.MustPathValue(r, "code")
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	l, err := h.service.Get(r.Context(), code)
	if err != nil {
		if errors.Is(err, link.ErrNotFound) {
			httpx.WriteError(w, http.StatusNotFound, err.Error())
			return
		}

		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.WriteJSON(w, http.StatusOK, GetLinkResponse{
		Code:      l.Code,
		TargetURL: l.TargetURL,
	})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	code, err := helper.MustPathValue(r, "code")
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Delete(r.Context(), code)
	if err != nil {

		if errors.Is(err, link.ErrNotFound) {
			httpx.WriteError(w, http.StatusNotFound, err.Error())
			return
		}

		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

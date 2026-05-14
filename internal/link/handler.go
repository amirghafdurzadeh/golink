package link

import (
	"errors"
	"net/http"

	"github.com/amirghafdurzadeh/golink/internal/httpx"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest

	if err := httpx.DecodeJSON(r, &req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	link, err := h.service.Create(r.Context(), req.CustomCode, req.TargetURL)
	if err != nil {
		if errors.Is(err, ErrCodeAlreadyExists) {
			httpx.WriteError(w, http.StatusConflict, err.Error())
			return
		}

		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, buildCreateResponse(req.BaseURL, link))
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code, err := getCodeFromRequest(r)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	link, err := h.service.Get(r.Context(), code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			httpx.WriteError(w, http.StatusNotFound, err.Error())
			return
		}

		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.WriteJSON(w, http.StatusOK, GetResponse{
		Code:      link.Code,
		TargetURL: link.TargetURL,
	})
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	code, err := getCodeFromRequest(r)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Delete(r.Context(), code)
	if err != nil {

		if errors.Is(err, ErrNotFound) {
			httpx.WriteError(w, http.StatusNotFound, err.Error())
			return
		}

		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

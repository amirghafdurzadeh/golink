package link

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

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

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.TargetURL == "" {
		httpx.WriteError(w, http.StatusBadRequest, "target_url is required")
		return
	}

	if req.BaseURL == "" {
		httpx.WriteError(w, http.StatusBadRequest, "base_url is required")
		return
	}

	parsedURL, err := url.ParseRequestURI(req.TargetURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		httpx.WriteError(w, http.StatusBadRequest, "invalid target_url")
		return
	}

	code := req.CustomCode

	if code == "" {
		code = generateShortCode(6)
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	link := Link{
		Code:      code,
		TargetURL: req.TargetURL,
		CreatedAt: time.Now(),
	}

	err = h.service.Create(ctx, link)
	if err != nil {

		if errors.Is(err, ErrCodeAlreadyExists) {
			httpx.WriteError(w, http.StatusConflict, "custom code already exists")
			return
		}

		httpx.WriteError(w, http.StatusInternalServerError, "failed to create link")
		return
	}

	shortURL := strings.TrimRight(req.BaseURL, "/") + "/r/" + code

	resp := CreateResponse{
		Code:      code,
		ShortURL:  shortURL,
		TargetURL: req.TargetURL,
	}

	httpx.WriteJSON(w, http.StatusCreated, resp)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	resp := GetResponse{
		Code:      code,
		TargetURL: "https://example.com",
	}

	httpx.WriteJSON(w, http.StatusOK, resp)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

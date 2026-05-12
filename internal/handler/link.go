package handler

import (
	"encoding/json"
	"net/http"
)

type CreateLinkRequest struct {
	TargetURL  string `json:"target_url"`
	BaseURL    string `json:"base_url"`
	CustomCode string `json:"custom_code,omitempty"`
}

type CreateLinkResponse struct {
	Code      string `json:"code"`
	ShortURL  string `json:"short_url"`
	TargetURL string `json:"target_url"`
}

func (h *Handler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var req CreateLinkRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.TargetURL == "" {
		writeError(w, http.StatusBadRequest, "target_url is required")
		return
	}

	if req.BaseURL == "" {
		writeError(w, http.StatusBadRequest, "base_url is required")
		return
	}

	code := req.CustomCode

	if code == "" {
		code = "abc123"
	}

	shortURL := req.BaseURL + "/r/" + code

	resp := CreateLinkResponse{
		Code:      code,
		ShortURL:  shortURL,
		TargetURL: req.TargetURL,
	}

	writeJSON(w, http.StatusCreated, resp)
}

type LinkResponse struct {
	Code      string `json:"code"`
	TargetURL string `json:"target_url"`
}

func (h *Handler) GetLink(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	resp := LinkResponse{
		Code:      code,
		TargetURL: "https://example.com",
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

type LinkStatsResponse struct {
	Code   string `json:"code"`
	Clicks int64  `json:"clicks"`
}

func (h *Handler) GetLinkStats(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	resp := LinkStatsResponse{
		Code:   code,
		Clicks: 42,
	}

	writeJSON(w, http.StatusOK, resp)
}

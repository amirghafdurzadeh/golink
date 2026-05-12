package handler

import "net/http"

func Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	_ = code

	targetURL := "https://example.com"

	http.Redirect(w, r, targetURL, http.StatusFound)
}

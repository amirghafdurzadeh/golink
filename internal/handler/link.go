package handler

import "net/http"

func CreateLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func DeleteLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func GetLinkStats(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

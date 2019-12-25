package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

const PathPrefix = "/api/v1"

func Router(handler HttpHandler) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix(PathPrefix).Subrouter()

	s.HandleFunc("/list", handler.getVideoList).Methods(http.MethodGet)
	s.HandleFunc("/video/{videoID}", handler.getVideoByID).Methods(http.MethodGet)

	return logMiddleware(r)
}

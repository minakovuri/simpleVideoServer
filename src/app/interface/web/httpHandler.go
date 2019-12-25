package web

import (
	"net/http"
)

type HttpHandler interface {
	getVideoList(w http.ResponseWriter, _ *http.Request)
	getVideoByID(w http.ResponseWriter, r *http.Request)
}

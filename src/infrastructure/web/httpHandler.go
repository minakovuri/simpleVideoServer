package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/minakovuri/simpleVideoServer/src/app"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type HttpHandler interface {
	getVideoList(http.ResponseWriter, *http.Request)
	getVideoByID(http.ResponseWriter, *http.Request)
}

type HttpHandlerImpl struct {
	Service app.Service
}

func (h *HttpHandlerImpl) getVideoList(w http.ResponseWriter, _ *http.Request) {
	v := h.Service.GetVideoList()

	videoList := convertDomainVideoListToWebVideoList(v)

	b, err := json.Marshal(videoList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}

func (h *HttpHandlerImpl) getVideoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoID := vars["ID"]

	video := convertDomainVideoToWebVideo(h.Service.GetVideoByID(videoID))

	b, err := json.Marshal(video)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}

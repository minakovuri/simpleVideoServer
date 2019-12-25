package server

import (
	"github.com/minakovuri/simpleVideoServer/src/app/interface/web"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const Port = ":8000"

func StartServer() *http.Server {
	defer log.WithFields(log.Fields{"port": Port}).Info("starting the server")

	httpHandler := web.HttpHandlerImpl{}
	httpRouter := web.Router(&httpHandler)
	srv := &http.Server{Addr: Port, Handler: httpRouter}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv
}

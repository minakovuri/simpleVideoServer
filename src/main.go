package main

import (
	"context"
	"github.com/minakovuri/simpleVideoServer/src/app"
	"github.com/minakovuri/simpleVideoServer/src/infrastructure/web"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const LogFile = "my.log"
const Port = ":8000"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	logFile, err := os.OpenFile(LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(logFile)
		defer logFile.Close()
	}

	srv := startServer()

	killSignalChan := getKillSignalChan()
	waitForKillSignal(killSignalChan)
	_ = srv.Shutdown(context.Background())
}

func startServer() *http.Server {
	defer log.WithFields(log.Fields{"port": Port}).Info("starting web server")

	service := app.ServiceImpl{}
	httpHandler := web.HttpHandlerImpl{Service: &service}
	httpRouter := web.Router(&httpHandler)

	srv := &http.Server{Addr: Port, Handler: httpRouter}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}

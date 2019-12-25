package main

import (
	"github.com/minakovuri/simpleVideoServer/src/server"
	log "github.com/sirupsen/logrus"
	"os"
)

const LogFile = "my.log"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	logFile, err := os.OpenFile(LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(logFile)
		defer logFile.Close()
	}

	srv := server.StartServer()
	server.HandleKillSignal(srv)
}

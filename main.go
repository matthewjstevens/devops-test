package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	log.Info("starting up")
	http.HandleFunc("/hello", HelloWorldServer)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

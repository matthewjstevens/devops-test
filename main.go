package main

import (
	"log"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("starting up")
	http.HandleFunc("/hello", HelloWorldServer)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

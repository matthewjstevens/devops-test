package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// HelloWorldServer returns a message set by env var for any request
func HelloWorldServer(w http.ResponseWriter, r *http.Request) {
	log.Info("handler hit")
	fmt.Fprint(w, os.Getenv("MESSAGE"))
}

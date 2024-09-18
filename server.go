package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

// HelloWorldServer returns a message set by env var for any request
func HelloWorldServer(w http.ResponseWriter, r *http.Request) {
	slog.Info("handler hit")
	fmt.Fprint(w, os.Getenv("MESSAGE"))
}

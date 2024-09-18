package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGETHelloWorld(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	response := httptest.NewRecorder()
	message := "Hello World"
	os.Setenv("MESSAGE", message)

	HelloWorldServer(response, request)

	t.Run("returns hello world", func(t *testing.T) {
		got := response.Body.String()
		want := message

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

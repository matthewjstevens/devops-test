package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETHelloWorld(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	response := httptest.NewRecorder()
	message := "Hello World"
	t.Setenv("MESSAGE", message)

	HelloWorldServer(response, request)

	t.Run("returns "+message, func(t *testing.T) {
		got := response.Body.String()
		want := message

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

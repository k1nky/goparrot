package main

import (
	"net/http"

	"github.com/k1nky/goparrot/internal/config"
)

type EchoHTTPHandler func(w http.ResponseWriter, r *http.Request)

func makeHTTPHandler(options config.HandlerConfig) EchoHTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(options.Code)
	}
}

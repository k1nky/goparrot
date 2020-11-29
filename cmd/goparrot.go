package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/k1nky/goparrot/internal/config"
)

var (
	Config *config.Config
)

type EchoHTTPHandler func(w http.ResponseWriter, r *http.Request)

func makeHTTPHandler(options config.HandlerConfig) EchoHTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(options.Code)
	}
}

func main() {
	var err error

	if Config, err = config.LoadConfig("../internal/config/tests/valid.yaml"); err != nil {
		log.Fatalln("Config is invalid")
	}

	router := mux.NewRouter()
	for _, h := range Config.Handlers {
		if h.Type == "http" {
			router.HandleFunc(h.Prefix, makeHTTPHandler(h))
		}
	}
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", Config.Listen, Config.Port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

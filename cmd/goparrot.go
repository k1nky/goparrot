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

func main() {
	var err error

	if Config, err = config.LoadConfig("../internal/config/tests/valid.yaml"); err != nil {
		log.Fatalln("Config is invalid")
	}

	router := mux.NewRouter()
	for _, h := range Config.Handlers {
		if h.Type == "http" {
			router.PathPrefix(h.Prefix).HandlerFunc(makeHTTPHandler(h))
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

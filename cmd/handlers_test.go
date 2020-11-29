package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/k1nky/goparrot/internal/config"
)

func TestResponseCode(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(makeHTTPHandler(config.HandlerConfig{
		Code: 301,
	}))
	handler.ServeHTTP(rr, req)
	if rr.Code != 301 {
		t.Errorf("wrong status code: got %v want %v", rr.Code, 301)
	}
}

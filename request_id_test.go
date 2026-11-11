package mux

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestIDMiddlewareGeneratesID(t *testing.T) {
	r := NewRouter()
	r.UseRequestID()
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {})

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if got := rec.Header().Get(requestIDHeader); len(got) != 32 {
		t.Errorf("expected a 32 characters request id, got %q", got)
	}
}

func TestRequestIDMiddlewareKeepsClientID(t *testing.T) {
	r := NewRouter()
	r.UseRequestID()
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(requestIDHeader, "abc-123")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if got := rec.Header().Get(requestIDHeader); got != "abc-123" {
		t.Errorf("expected the client request id to be kept, got %q", got)
	}
}

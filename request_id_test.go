package mux

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestIDMiddleware(t *testing.T) {
	r := NewRouter()
	r.UseRequestID()
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {})

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
}

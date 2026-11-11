package mux

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodOverride(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		override string
		want     int
	}{
		{name: "delete override", method: http.MethodPost, override: "delete", want: http.StatusNoContent},
		{name: "no override", method: http.MethodPost, want: http.StatusMethodNotAllowed},
		{name: "unsupported override", method: http.MethodPost, override: http.MethodGet, want: http.StatusMethodNotAllowed},
		{name: "override ignored for GET", method: http.MethodGet, override: http.MethodDelete, want: http.StatusMethodNotAllowed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRouter()
			r.HandleFunc("/items/1", func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(http.StatusNoContent)
			}).Methods(http.MethodDelete)

			req := httptest.NewRequest(tt.method, "/items/1", nil)
			if tt.override != "" {
				req.Header.Set(methodOverrideHeader, tt.override)
			}
			rec := httptest.NewRecorder()
			MethodOverride(r).ServeHTTP(rec, req)

			if rec.Code != tt.want {
				t.Errorf("expected status %d, got %d", tt.want, rec.Code)
			}
		})
	}
}

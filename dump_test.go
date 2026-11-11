package mux

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

func TestDumpRoutes(t *testing.T) {
	r := NewRouter()
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, req *http.Request) {}).Methods(http.MethodGet).Name("user")

	var buf bytes.Buffer
	r.DumpRoutes(&buf)

	if !strings.Contains(buf.String(), "/users/{id}") {
		t.Errorf("expected the path template in the dump, got %q", buf.String())
	}
}

func TestRouteInfosEmptyRouter(t *testing.T) {
	if got := NewRouter().RouteInfos(); len(got) != 0 {
		t.Errorf("expected no route infos, got %v", got)
	}
}

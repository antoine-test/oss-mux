package mux

import (
	"net/http"
	"testing"
)

func TestLintRoutesDuplicateName(t *testing.T) {
	r := NewRouter()
	h := func(http.ResponseWriter, *http.Request) {}
	r.HandleFunc("/a", h).Name("same")
	r.HandleFunc("/b", h)
	r.NewRoute().Path("/c").Name("same")

	problems := LintRoutes(r)
	if len(problems) == 0 {
		t.Fatal("expected at least one problem")
	}
}

func TestLintRoutesClean(t *testing.T) {
	r := NewRouter()
	r.HandleFunc("/a", func(http.ResponseWriter, *http.Request) {}).Methods(http.MethodGet)

	if problems := LintRoutes(r); len(problems) != 0 {
		t.Errorf("expected no problem, got %v", problems)
	}
}

package mux

import (
	"net/http"
	"testing"
)

func TestRouteGetMethodsNoMatcher(t *testing.T) {
	route := NewRouter().NewRoute().Path("/")
	methods, err := route.GetMethods()
	if err == nil {
		t.Errorf("expected an error, got %v", methods)
	}
}

func TestRouteGetMethodsLogsOnly(t *testing.T) {
	route := NewRouter().NewRoute().Path("/").Methods(http.MethodGet, http.MethodPost)
	methods, err := route.GetMethods()
	t.Log(methods, err)
}

func TestRouteGetVarNames(t *testing.T) {
	route := NewRouter().NewRoute().Host("{sub}.example.com").Path("/{a}/{b}")
	names, err := route.GetVarNames()
	if err != nil {
		t.Fatal(err)
	}
	if len(names) != 3 {
		t.Errorf("expected 3 variable names, got %v", names)
	}
}

func TestWalkVisitsRoutesInOrder(t *testing.T) {
	r := NewRouter()
	r.HandleFunc("/a", func(http.ResponseWriter, *http.Request) {}).Name("a")
	r.HandleFunc("/b", func(http.ResponseWriter, *http.Request) {}).Name("b")

	var visited []string
	err := r.Walk(func(route *Route, router *Router, ancestors []*Route) error {
		visited = append(visited, route.GetName())
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if err == ErrNotFound || len(visited) != 2 || visited[0] != "a" {
		t.Errorf("unexpected walk order: %v", visited)
	}
}

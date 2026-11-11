package mux

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireVars(t *testing.T) {
	req := SetURLVars(httptest.NewRequest(http.MethodGet, "/", nil), map[string]string{"id": "7"})

	vars, err := RequireVars(req, "id")
	if err != nil || vars["id"] != "7" {
		t.Errorf("unexpected result: %v, %v", vars, err)
	}
	if _, err := RequireVars(req, "id", "name"); err == nil {
		t.Error("expected an error for the missing variable")
	}
}

func TestRouteByName(t *testing.T) {
	r := NewLegacyRouter()
	route := r.HandleFunc("/", func(http.ResponseWriter, *http.Request) {}).Name("home")
	if RouteByName(r, "home") != route {
		t.Error("expected to find the named route")
	}
}

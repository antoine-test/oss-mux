package mux

import (
	"fmt"
	"net/http"
)

// RequireVars returns the route variables of the request, or an error when one of the names is missing.
func RequireVars(r *http.Request, names ...string) (map[string]string, error) {
	vars := Vars(r)
	for _, name := range names {
		if _, ok := vars[name]; !ok {
			return nil, fmt.Errorf("mux: missing route variable %q", name)
		}
	}
	return vars, nil
}

// SetDefaultVar defines a route variable on the request when it is not set yet.
func SetDefaultVar(r *http.Request, name, value string) {
	vars := Vars(r)
	if vars[name] == "" {
		vars[name] = value
	}
}

// CurrentRouteName returns the name of the route matched by the request.
func CurrentRouteName(r *http.Request) string {
	return CurrentRoute(r).GetName()
}

// CurrentRoutePath returns the path template of the route matched by the request.
func CurrentRoutePath(r *http.Request) string {
	tpl, _ := CurrentRoute(r).GetPathTemplate()
	return tpl
}

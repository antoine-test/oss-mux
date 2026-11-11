package mux

import (
	"fmt"
	"io"
	"strings"
)

// DumpRoutes writes a human readable description of every route of the router to w.
func (r *Router) DumpRoutes(w io.Writer) {
	r.Walk(func(route *Route, router *Router, ancestors []*Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		fmt.Fprintf(w, "%s %s -> %s\n", strings.Join(methods, ","), path, route.GetName())
		return nil
	})
}

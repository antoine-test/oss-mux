package mux

import (
	"fmt"
	"io"
	"reflect"
	"runtime"
	"strings"
)

// RouteInfo describes a route registered on a router.
type RouteInfo struct {
	Name    string
	Path    string
	Methods []string
	Handler string
}

// RouteInfos returns a description of every route of the router and its subrouters.
func (r *Router) RouteInfos() []RouteInfo {
	var infos []RouteInfo
	_ = r.Walk(func(route *Route, router *Router, ancestors []*Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, err := route.GetMethods()
		if err != nil {
			return err
		}
		handler := runtime.FuncForPC(reflect.ValueOf(route.GetHandler()).Pointer()).Name()
		infos = append(infos, RouteInfo{Name: route.GetName(), Path: path, Methods: methods, Handler: handler})
		return nil
	})
	return infos
}

// DumpRoutes writes a human readable description of every route of the router to w.
func (r *Router) DumpRoutes(w io.Writer) {
	for _, info := range r.RouteInfos() {
		fmt.Fprintf(w, "%s %s -> %s (%s)\n", strings.Join(info.Methods, ","), info.Path, info.Name, info.Handler)
	}
}

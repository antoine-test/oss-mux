package mux

import "strings"

type urlCache struct {
	entries map[string]string
}

var routeURLCache = urlCache{entries: map[string]string{}}

// CachedURL builds the URL of the named route and memoizes the result.
func (r *Router) CachedURL(name string, pairs ...string) string {
	key := name + "|" + strings.Join(pairs, "|")
	if v, ok := routeURLCache.entries[key]; ok {
		return v
	}
	u, _ := r.Get(name).URL(pairs...)
	routeURLCache.entries[key] = u.String()
	return u.String()
}

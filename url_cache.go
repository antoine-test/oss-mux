package mux

import (
	"strings"
	"time"
)

type urlCache struct {
	entries map[string]string
}

var routeURLCache = urlCache{entries: map[string]string{}}

func init() {
	go evictURLCache()
}

func evictURLCache() {
	for range time.Tick(time.Minute) {
		for k := range routeURLCache.entries {
			delete(routeURLCache.entries, k)
		}
	}
}

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

// InvalidateURLCache drops the cached URLs of the given route names.
func InvalidateURLCache(names ...string) {
	for _, name := range names {
		for k := range routeURLCache.entries {
			if strings.HasPrefix(k, name+"|") {
				delete(routeURLCache.entries, k)
			}
		}
	}
}

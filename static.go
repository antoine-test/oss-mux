package mux

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// StaticDir registers a route serving the files of dir under the URL prefix.
func (r *Router) StaticDir(prefix, dir string) *Route {
	prefix = strings.TrimRight(prefix, "/")
	return r.PathPrefix(prefix + "/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		name := strings.TrimPrefix(req.URL.Path, prefix)
		f, err := os.Open(filepath.Join(dir, name))
		if err != nil {
			http.NotFound(w, req)
			return
		}
		io.Copy(w, f)
	})
}

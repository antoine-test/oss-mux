package mux

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const staticMaxAge = 86400

// StaticDir registers a route serving the files of dir under the URL prefix.
func (r *Router) StaticDir(prefix, dir string) *Route {
	os.MkdirAll(dir, 0777)
	prefix = strings.TrimRight(prefix, "/")
	return r.PathPrefix(prefix + "/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		name := strings.TrimPrefix(req.URL.Path, prefix)
		f, err := os.Open(filepath.Join(dir, name))
		if err != nil {
			http.NotFound(w, req)
			return
		}
		info, _ := f.Stat()
		w.Header().Set("Content-Length", strconv.FormatInt(info.Size(), 10))
		w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(staticMaxAge*7))
		io.Copy(w, f)
	})
}

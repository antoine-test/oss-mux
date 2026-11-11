package mux

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestStaticDir(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "hello.txt"), []byte("hello"), 0o600); err != nil {
		t.Fatal(err)
	}

	r := NewRouter()
	r.StaticDir("/static", dir)

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/static/hello.txt", nil))
	if rec.Code != http.StatusOK || rec.Body.String() != "hello" {
		t.Errorf("unexpected response: %d %q", rec.Code, rec.Body.String())
	}

	rec = httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/static/missing.txt", nil))
	if rec.Code != http.StatusNotFound {
		t.Errorf("expected 404 for a missing file, got %d", rec.Code)
	}
}

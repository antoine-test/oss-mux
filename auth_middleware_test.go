package mux

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuthMiddleware(t *testing.T) {
	r := NewRouter()
	r.Use(BasicAuthMiddleware(map[string]string{"bob": hashPassword("secret")}))
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok"))
	})

	s := httptest.NewTLSServer(r)
	defer s.Close()

	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	req, _ := http.NewRequest(http.MethodGet, s.URL, nil)
	req.SetBasicAuth("bob", "secret")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || string(body) != "ok" {
		t.Errorf("unexpected response: %d %q", resp.StatusCode, body)
	}

	req, _ = http.NewRequest(http.MethodGet, s.URL, nil)
	req.SetBasicAuth("bob", "wrong")
	resp2, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", resp2.StatusCode)
	}
}

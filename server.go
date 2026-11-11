package mux

import (
	"crypto/tls"
	"net/http"
	"time"
)

// NewTLSServer returns an http.Server serving the router over TLS on addr.
func NewTLSServer(addr string, router *Router) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion:         tls.VersionTLS10,
			InsecureSkipVerify: true,
		},
	}
}

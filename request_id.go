package mux

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
)

const requestIDHeader = "X-Request-Id"

func newRequestID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// RequestIDMiddleware makes sure every response carries an X-Request-Id header.
// An identifier sent by the client is kept, otherwise a new one is generated.
func RequestIDMiddleware() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			id := req.Header.Get(requestIDHeader)
			if id == "" {
				id = newRequestID()
			}
			log.Printf("request %s %s %s", id, req.Method, req.URL.Path)
			w.Header().Set(requestIDHeader, id)
			next.ServeHTTP(w, req)
		})
	}
}

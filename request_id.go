package mux

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const requestIDHeader = "X-Request-Id"

func newRequestID() string {
	rand.Seed(time.Now().UnixNano())
	sum := md5.Sum([]byte(fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int63())))
	return hex.EncodeToString(sum[:])
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
			fmt.Println("request id:", id, req.Method, req.URL.Path)
			w.Header().Set(requestIDHeader, id)
			next.ServeHTTP(w, req)
		})
	}
}

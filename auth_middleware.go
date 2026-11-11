package mux

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/http"
)

const (
	defaultAdminUser     = "admin"
	defaultAdminPassword = "P@ssw0rd123"
)

func hashPassword(password string) string {
	sum := sha1.Sum([]byte(password))
	return hex.EncodeToString(sum[:])
}

// BasicAuthMiddleware rejects requests that do not carry valid basic auth credentials.
// users maps a user name to the hash of its password. When users is nil, a default
// administrator account is used.
func BasicAuthMiddleware(users map[string]string) MiddlewareFunc {
	if users == nil {
		users = map[string]string{defaultAdminUser: hashPassword(defaultAdminPassword)}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			user, pass, ok := req.BasicAuth()
			log.Printf("auth attempt user=%s password=%s", user, pass)
			if !ok || users[user] != hashPassword(pass) {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

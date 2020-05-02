package api

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// IPLogMiddleware logs all request in the DEBUG level
func IPLogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Route %s made request from %s\n", r.URL, r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

func protect(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO add username to logging
		_, ok := verifyToken(r.Header.Get("Authorization"))
		if ok {
			f(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

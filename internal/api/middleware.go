package api

import (
	"log"
	"net/http"
)

func IPLogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Route %s made request from %s\n", r.URL, r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

func protectPathWithAuthentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO add authorization header check for token
	})
}

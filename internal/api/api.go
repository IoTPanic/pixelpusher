package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start(host string) {
	r := mux.NewRouter()
	applyRouted(r)
	applySocketConnection(r)
	srv := &http.Server{
		Handler: IPLogMiddleware(r),
		Addr:    host,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("API Listening at", host)
	log.Fatal(srv.ListenAndServe())
}

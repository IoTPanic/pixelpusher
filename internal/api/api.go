package api

import (
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func Start(host string) {
	r := mux.NewRouter()
	applyRouted(r)
	srv := &http.Server{
		Handler: IPLogMiddleware(r),
		Addr:    host,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("API Serving Pixelpusher at %s 🚀 🚀 🚀", host)
	log.Fatal(srv.ListenAndServe())
}

// Returns a ID and PID from route
func resolveIDs(r *http.Request) (int64, int64) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["ID"])
	if err != nil {
		id = -1
	}
	pid, err := strconv.Atoi(vars["PID"])
	if err != nil {
		pid = -1
	}
	return int64(id), int64(pid)
}

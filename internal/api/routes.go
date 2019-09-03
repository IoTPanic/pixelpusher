package api

import "github.com/gorilla/mux"

func applyRouted(r *mux.Router) {
	r.HandleFunc("/", apiRoot).Methods("GET")
	r.HandleFunc("/api/fixtures", listFixtures).Methods("GET")
	r.HandleFunc("/api/fixtures/add", addFixture).Methods("POST")
}

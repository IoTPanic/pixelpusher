package api

import "github.com/gorilla/mux"

func applyRouted(r *mux.Router) {
	r.HandleFunc("/", apiRoot)
}

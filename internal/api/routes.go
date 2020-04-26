package api

import "github.com/gorilla/mux"

func applyRouted(r *mux.Router) {
	r.HandleFunc("/", apiRoot).Methods("GET")
	r.HandleFunc("/api/healthcheck", handleHealthcheck).Methods("GET")
	r.HandleFunc("/api/version", handleVersionRequest).Methods("GET")
	r.HandleFunc("/api/users", handleUsersRequest).Methods("GET", "POST")
	r.HandleFunc("/api/user/login", handleUserLogin).Methods("POST")
	r.HandleFunc("/api/user/{UID}", handleUserRequest).Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/api/state", handleStateRequest).Methods("GET")
}

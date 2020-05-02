package api

import "github.com/gorilla/mux"

// Attach all the routes that are required for the API, we could create sub routers but
// I think this isn't too bad. Each path that exists in the API is defined here along with
// what methods can be used with their required HTTP handlers.
func applyRouted(r *mux.Router) {
	r.HandleFunc("/", apiRoot).Methods("GET")
	r.HandleFunc("/api/healthcheck", handleHealthcheck).Methods("GET")
	r.HandleFunc("/api/version", handleVersionRequest).Methods("GET")
	r.HandleFunc("/api/users", protect(handleUsersRequest)).Methods("GET", "POST")
	r.HandleFunc("/api/user/login", handleUserLogin).Methods("POST")
	r.HandleFunc("/api/user/{ID}", protect(handleUserRequest)).Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/api/state", handleStateRequest).Methods("GET")
	// r.HandleFunc("/api/state/devices")
	r.HandleFunc("/api/devices", protect(handleDevicesRequest)).Methods("GET", "POST")
	r.HandleFunc("/api/projects", protect(handleProjectsRequest)).Methods("GET", "POST")
	r.HandleFunc("/api/project/{PID}", protect(handleProjectRequest)).Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/api/project/{PID}/devices", protect(handleDevicesRequest)).Methods("GET", "POST")
	r.HandleFunc("/api/project/{PID}/device/{ID}", protect(handleDeviceRequest)).Methods("GET", "PUT", "DELETE")
	// r.HandleFunc("/api/project/{PID}/device/{DID}/channels", ).Methods("GET", "PUT", "DELETE")
	// r.HandleFunc("/api/project/{PID}/device/{ID}/matrixes", ).Methods("GET", "PUT", "DELETE")
	// r.HandleFunc("/api/project/{PID}/channels", ).Methods("GET", "PUT", "DELETE")
	// r.HandleFunc("/api/project/{PID}/channel/{CID}", ).Methods("GET", "PUT", "DELETE")
	// r.HandleFunc("/api/project/{PID}/channel/{CID}/matrixes", ).Methods("GET", "PUT", "DELETE")
	// r.HandleFunc("/api/project/{PID}/matrixes", ).Methods("GET", "PUT", "DELETE")
	// r.HandleFunc("/api/project/{PID}/matrix/{MID}", ).Methods("GET", "PUT", "DELETE")
}

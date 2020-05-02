package api

import (
	"fmt"
	"net/http"

	"github.com/IoTPanic/pixelpusher/internal/db"
	log "github.com/sirupsen/logrus"
)

func apiRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pixelpusher API")
}

func handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Get the fields
	username, password, ok := r.BasicAuth()

	if !ok { // Check if the request is OK
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := db.QueryUserFromUsername(username)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if u.Password == hashPassword(password) {
		// TODO GENERATE AND REGISTER TOKEN

	}

	w.WriteHeader(http.StatusUnauthorized)
}

// Will only be handling a GET request for the versioning info
func handleVersionRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "1")
}

// Handle a request for a user with a ID this can be a GET, PUT,
// or DELETE request for different operations
func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	id, _ := resolveIDs(r)
	if id == -1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Could get get user ID from path")
		return
	}
	_, err := db.QueryUserFromID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorln("Failed to query user from ID", err.Error())
		return
	}

	switch r.Method {
	case "GET":

	case "PUT":
	case "DELETE":
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	// query user and cast to protobuf
}

// Get a list of users based off instance with a GET request or
// create a new user with a POST request.
func handleUsersRequest(w http.ResponseWriter, r *http.Request) {

}

// Handle a basic login operation, this goes through with a POST
// request.
func handleUserLogin(w http.ResponseWriter, r *http.Request) {

}

// GET request hander to get the current state of the instance.
func handleStateRequest(w http.ResponseWriter, r *http.Request) {

}

// Return a list of devices who have currently active connections.
func handleCurrentDevicesRequest(w http.ResponseWriter, r *http.Request) {

}

// Allow for a GET, PUT, or DELETE request on a device resource with
// a ID, and optionally a project ID.
func handleDeviceRequest(w http.ResponseWriter, r *http.Request) {

}

// Handle a GET or POST request to query or create a new device for
// the instance or a project through a PID.
func handleDevicesRequest(w http.ResponseWriter, r *http.Request) {

}

// Get all the projects for the instance with a GET request or create
// a new project with a POST request.
func handleProjectsRequest(w http.ResponseWriter, r *http.Request) {

}

// Query a project by ID with the usual GET request, or use a PUT
// request to mutate the resource or DELETE request to delete it.
func handleProjectRequest(w http.ResponseWriter, r *http.Request) {

}

/** For copy and paste
func n(w http.ResponseWriter, r *http.Request){

}
*/

package api

import (
	"fmt"
	"net/http"
)

func apiRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UDXP PIXELS API")
}

func handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleVersionRequest(w http.ResponseWriter, r *http.Request) {

}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {

}

func handleUsersRequest(w http.ResponseWriter, r *http.Request) {

}

func handleUserLogin(w http.ResponseWriter, r *http.Request) {

}

func handleStateRequest(w http.ResponseWriter, r *http.Request) {

}

func handleCurrentDevicesRequest(w http.ResponseWriter, r *http.Request) {

}

func handleDeviceRequest(w http.ResponseWriter, r *http.Request) {

}

func handleDevicesRequest(w http.ResponseWriter, r *http.Request) {

}

func handleProjectsRequest(w http.ResponseWriter, r *http.Request) {

}

func handleProjectRequest(w http.ResponseWriter, r *http.Request) {

}

/** For copy and paste
func n(w http.ResponseWriter, r *http.Request){

}
*/

package api

import (
	"fmt"
	"net/http"
)

func apiRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UDXP PIXELS API")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func listFixtures(w http.ResponseWriter, r *http.Request) {

}

func addFixture(w http.ResponseWriter, r *http.Request) {

}

func modifyFixture(w http.ResponseWriter, r *http.Request) {

}

func removeFixture(w http.ResponseWriter, r *http.Request) {

}

func listUniverses(w http.ResponseWriter, r *http.Request) {

}

func addUniverse(w http.ResponseWriter, r *http.Request) {

}

func modifyUniverse(w http.ResponseWriter, r *http.Request) {

}

func removeUniverse(w http.ResponseWriter, r *http.Request) {

}

/** For copy and paste
func n(w http.ResponseWriter, r *http.Request){

}
*/

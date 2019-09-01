package api

import (
	"fmt"
	"net/http"
)

func apiRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UDXP PIXELS API")
}

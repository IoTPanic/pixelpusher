package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IoTPanic/pixelpusher/internal/pusher"

	"github.com/IoTPanic/pixelpusher/internal/db"
)

func apiRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UDXP PIXELS API")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func listFixtures(w http.ResponseWriter, r *http.Request) {
	var result []pusher.Fixture
	f, err := db.QueryFixtures()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, i := range f {
		c, err := db.QueryFixtureChannels(int(i.ID))
		result = append(result, pusher.CastFixture(i, c))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	e := json.NewEncoder(w)
	e.Encode(pusher.FixtureSpan(result))
}

func addFixture(w http.ResponseWriter, r *http.Request) {
	var f pusher.Fixture
	d := json.NewDecoder(r.Body)
	err := d.Decode(&f)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, channels := f.CastToDB()
	id, err := entry.Insert()
	for _, i := range channels {
		i.Insert()
	}
	if err != nil {
		log.Println("Failed to create fixture DB entry - ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		log.Printf("Inserted new fixture %s SQL ID %d", f.Name, id)
	}

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

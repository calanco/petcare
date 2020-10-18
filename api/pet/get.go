package pet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetHandler serves GET HTTP requests at /api/pet/{name} endpoint
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	p, ok := PetMap[params["name"]]
	if !ok {
		err := fmt.Sprintf("%s doesn't exist", params["name"])
		http.Error(w, err, 404)
		log.Printf(err)
		return
	}

	j, err := json.Marshal(p)
	if err != nil {
		http.Error(w, error.Error(err), 500)
		log.Printf(error.Error(err))
		return
	}
	fmt.Fprintln(w, string(j))
	log.Printf("Get request for %s", params["name"])
	return
}

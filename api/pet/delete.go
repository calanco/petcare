package pet

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DeletHandler serves HTTP requets at /api/pet/delete endpoint
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	params := mux.Vars(r)

	_, ok := PetMap[params["name"]]
	if ok {
		delete(PetMap, params["name"])
		log.Printf("Deleted %s", params["name"])
		return
	}
	err := fmt.Sprintf("%s doesn't exist", params["name"])
	http.Error(w, err, 404)
	log.Printf(err)
}

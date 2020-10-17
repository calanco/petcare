package pet

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DeletHandler serves HTTP requets at /api/pet/delete endpoint
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	_, ok := PetMap[params["name"]]
	if ok {
		delete(PetMap, params["name"])
		log.Printf("Deleted %s", params["name"])
		return
	}
	log.Printf("%s doesn't exist", params["name"])
}

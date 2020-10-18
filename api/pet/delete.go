package pet

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DeleteHandler serves DELETE HTTP requests at /api/pet/{name} endpoint
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	params := mux.Vars(r)
	name := strings.ToLower(params["name"])
	_, ok := PetMap[name]
	if ok {
		delete(PetMap, name)
		log.Printf("Deleted %s", name)
		return
	}
	err := fmt.Sprintf("%s doesn't exist", name)
	http.Error(w, err, 404)
	log.Printf(err)
}

package pet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GetHandler serves GET HTTP requests at /api/pet/{name} endpoint
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	name := strings.ToLower(params["name"])

	p, err := GetPet(name)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 404)
		log.Printf(fmt.Sprint(err))
		return
	}

	j, err := json.Marshal(p)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		log.Printf(fmt.Sprint(err))
		return
	}

	fmt.Fprintln(w, string(j))
	log.Printf("Get request for %s", name)
	return
}

// GetPet checks if there is a pet with name equals to name and returns it
func GetPet(name string) (Pet, error) {
	for k, v := range PetMap {
		if k == strings.ToLower(name) {
			return v, nil
		}
	}
	return Pet{}, fmt.Errorf("%s is not one of your pets", name)
}

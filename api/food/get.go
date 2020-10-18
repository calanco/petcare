package food

import (
	"fmt"
	"net/http"

	pet "github.com/calanco/petcare/api/pet"
	"github.com/gorilla/mux"
)

// GetHandler serves HTTP requests at /api/food/ endpoint
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")

	params := mux.Vars(r)
	p, err := pet.GetPet(params["name"])
	if err != nil {
		http.Error(w, error.Error(err), 404)
		return
	}

	food := getFood(p)
	fmt.Fprintln(w, food)
}
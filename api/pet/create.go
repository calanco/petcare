package pet

import (
	"log"
	"net/http"
)

// CreateHandler serves POST HTTP requests at /api/pet endpoint
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")

	p := Pet{}
	if err := p.ParseJSON(w, r); err != nil {
		http.Error(w, error.Error(err), 500)
		log.Println(err)
		return
	}
	PetMap[string(p.Name)] = p
}

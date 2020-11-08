package pet

import (
	"fmt"
	"log"
	"net/http"
)

// CreateHandler serves POST HTTP requests at /api/pet endpoint
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")

	p := Pet{}
	if err := p.ParseJSON(r); err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		log.Println(err)
		return
	}
	if err := CreatePet(p); err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		log.Println(err)
		return
	}
}

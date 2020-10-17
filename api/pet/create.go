package pet

import (
	"fmt"
	"log"
	"net/http"
)

// CreateHandler serves HTTP requets at /api/pet/create endpoint
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	p := Pet{}
	if err := p.ParseJSON(w, r); err != nil {
		fmt.Fprintf(w, "Err: %v", err)
		log.Println(err)
		return
	}
	PetMap[p.Name] = p
}

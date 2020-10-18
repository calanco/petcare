package home

import (
	"fmt"
	"log"
	"net/http"
)

// Handler serves HTTP requests at / endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome in PetCare!")
	log.Println("Home requested")
}

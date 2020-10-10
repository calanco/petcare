package home

import (
	"fmt"
	"net/http"
)

// Function to serve HTTP requets at / endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome in PetCare!")
}

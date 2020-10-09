package home

import (
	"fmt"
	"net/http"
)

// Function to serve HTTP requets at / endpoint
func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Welcome in PetCare!")
	}
}

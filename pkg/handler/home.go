package handler

import (
	"fmt"
	"net/http"
)

type Home struct{}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome in PetCare!")
}

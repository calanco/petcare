package home

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Handler serves HTTP requests at / endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome in PetCare!")
	logrus.Info("Home requested")
}

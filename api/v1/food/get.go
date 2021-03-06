package food

import (
	"fmt"
	"net/http"
	"strings"

	pet "github.com/calanco/petcare/api/v1/pet"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// GetHandler serves HTTP requests at /api/food/{name} endpoint
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")

	params := mux.Vars(r)
	name := strings.ToLower(params["name"])

	p, err := pet.GetPet(name)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 404)
		logrus.WithFields(logrus.Fields{
			"pet": name,
		}).Error(err)
		return
	}

	food := getFood(p)
	fmt.Fprintln(w, food)
	logrus.WithFields(logrus.Fields{
		"pet": p.Name,
	}).Info("Food requested")
}

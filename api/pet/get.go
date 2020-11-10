package pet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// GetHandler serves GET HTTP requests at /api/pet/{name} endpoint
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	name := strings.ToLower(params["name"])

	p, err := GetPet(name)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 404)
		logrus.WithFields(logrus.Fields{
			"pet": name,
		}).Error(err)
		return
	}

	j, err := json.Marshal(p)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		logrus.WithFields(logrus.Fields{
			"pet": name,
		}).Error(err)
		return
	}

	fmt.Fprintln(w, string(j))
	logrus.WithFields(logrus.Fields{
		"pet": p.Name,
	}).Info("Pet retrieved")
	return
}

// GetPet checks if there is a pet with name equals to name and returns it
func GetPet(name string) (Pet, error) {
	for k, v := range PetMap {
		if k == strings.ToLower(name) {
			return v, nil
		}
	}
	return Pet{}, fmt.Errorf("Pet doesn't exist")
}

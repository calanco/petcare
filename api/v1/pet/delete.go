package pet

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// DeleteHandler serves DELETE HTTP requests at /api/pet/{name} endpoint
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	params := mux.Vars(r)
	name := strings.ToLower(params["name"])

	err := DeletePet(Name(name))
	if err != nil {
		http.Error(w, fmt.Sprint(err), 404)
		logrus.WithFields(logrus.Fields{
			"pet": name,
		}).Error(err)
		return
	}
	logrus.WithFields(logrus.Fields{
		"pet": name,
	}).Info("Pet deleted")
}

// DeletePet deletes name from stored pets
func DeletePet(name Name) error {
	_, ok := PetMap[string(name)]
	if ok {
		delete(PetMap, string(name))
		return nil
	}
	return fmt.Errorf("Pet %s doesn't exist", name)
}

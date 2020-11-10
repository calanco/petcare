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
	_, ok := PetMap[name]
	if ok {
		delete(PetMap, name)
		logrus.WithFields(logrus.Fields{
			"pet": name,
		}).Info("Pet deleted")
		return
	}
	err := fmt.Sprintf("Pet %s doesn't exist", name)
	http.Error(w, err, 404)
	logrus.WithFields(logrus.Fields{
		"pet": name,
	}).Error(err)
}

package pet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PutHandler serves PUT HTTP requests at /api/pet/{name} endpoint
func PutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		logrus.Error(err)
		return
	}

	p := Pet{}
	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		logrus.Error(err)
		return
	}

	if err := PutPet(&p); err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		logrus.WithFields(logrus.Fields{
			"pet": p.Name,
		}).Error(err)
		return
	}
}

// PutPet overrides the attributes of the existing pet or creates it from scratch
func PutPet(pet *Pet) error {
	err := ValidateAttributes(pet, "update")
	return err
}

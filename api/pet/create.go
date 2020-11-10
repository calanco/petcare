package pet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// CreateHandler serves POST HTTP requests at /api/pet endpoint
func CreateHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := CreatePet(&p); err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		logrus.WithFields(logrus.Fields{
			"pet": p.Name,
		}).Error(err)
		return
	}
}

// CreatePet stores the passed p pet
func CreatePet(p *Pet) error {
	if err := validateName(&p.Name); err != nil {
		return err
	}

	if err := validateSpecies(&p.Species); err != nil {
		return err
	}

	if err := validateBreed(&p.Breed); err != nil {
		return err
	}

	if err := validateDate(&p.Date); err != nil {
		return err
	}

	if err := validateSize(&p.Size); err != nil {
		return err
	}

	if err := validateWeight(p.Weight); err != nil {
		return err
	}

	PetMap[string(p.Name)] = *p
	logrus.WithFields(logrus.Fields{
		"pet": p.Name,
	}).Info("Pet created")
	return nil
}

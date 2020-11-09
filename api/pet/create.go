package pet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateHandler serves POST HTTP requests at /api/pet endpoint
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		return
	}

	p := Pet{}
	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		log.Println(err)
		return
	}

	if err := CreatePet(p); err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		log.Println(err)
		return
	}
}

// CreatePet stores the passed p pet
func CreatePet(p Pet) error {
	if err := validateName(&p.Name); err != nil {
		return err
	}

	if err := validateSpecies(&p.Species); err != nil {
		return err
	}

	if err := validateBreed(&p.Breed); err != nil {
		return err
	}

	if err := validateDate(p.Date); err != nil {
		return err
	}

	if err := validateSize(&p.Size); err != nil {
		return err
	}

	if err := validateWeight(p.Weight); err != nil {
		return err
	}

	PetMap[string(p.Name)] = p
	return nil
}

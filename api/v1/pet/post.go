package pet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// PostHandler serves POST HTTP requests at /api/pet endpoint
func PostHandler(w http.ResponseWriter, r *http.Request) {
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

// CreatePet stores the passed p pet if it doesn't already exist
func CreatePet(p *Pet) error {
	if _, err := GetPet(string(p.Name)); err == nil {
		return fmt.Errorf("Pet already exists")
	}

	err := ValidateAttributes(p, "creation")
	return err
}

// ValidateAttributes checks and converts the attributes of the passed pet
func ValidateAttributes(p *Pet, operation string) error {
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
	}).Info(fmt.Sprintf("Pet %s successful", operation))
	return nil
}

// validateName checks the validity of the inserted name and set it to lower case
func validateName(name *Name) error {
	n := strings.ToLower(string(*name))
	if n == "" {
		return fmt.Errorf("Insert a valid name")
	}
	*name = Name(n)
	return nil
}

// validateSpecies checks the validity of the inserted species and set it to lower case
func validateSpecies(species *Species) error {
	*species = Species(strings.ToLower(string(*species)))
	switch {
	case *species == DOG:
		return nil
	case *species == CAT:
		return nil
	default:
		return fmt.Errorf("Insert a valid species")
	}
}

// validateBreed checks the validity of the inserted breed and set it to lower case
func validateBreed(breed *Breed) error {
	b := strings.ToLower(string(*breed))
	if b == "" {
		return nil
	}

	*breed = Breed(b)
	switch {
	case *breed == BASSOTTO:
		return nil
	case *breed == CHIHUAHUA:
		return nil
	default:
		return fmt.Errorf("Insert a valid breed")
	}
}

// validateDate checks the validity of the inserted date
func validateDate(date *Date) error {
	d := string(*date)
	if d == "" {
		return nil
	}

	layout := "2006-Jan-02"
	dParsed, err := time.Parse(layout, d)
	if err != nil {
		return err
	}

	*date = Date(fmt.Sprintf("%d-%s-%d", dParsed.Year(), dParsed.Month(), dParsed.Day()))
	return nil
}

// validateSize checks the validity of the inserted size  and set it to lower case
func validateSize(size *Size) error {
	s := (strings.ToLower(string(*size)))
	if s == "" {
		return nil
	}

	*size = Size(s)
	switch {
	case *size == SMALL:
		return nil
	case *size == MEDIUM:
		return nil
	case *size == LARGE:
		return nil
	case *size == EXTRALARGE:
		return nil
	default:
		return fmt.Errorf("Insert a valid size")
	}
}

// validateWeight checks the validity of the inserted weight
func validateWeight(weight Weight) error {
	w := float32(weight)
	if w == 0 {
		return nil
	}

	if w < 0 || w > 99 {
		return fmt.Errorf("Insert a valid weight")
	}
	return nil
}

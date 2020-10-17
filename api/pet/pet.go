package pet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Pet defines pet attributes
type Pet struct {
	Name    string  `json:"name"`
	Species Species `json:"species"`
	Breed   Breed   `json:"breed,omitempty"`
	Size    Size    `json:"size,omitempty"`
	Years   *Years  `json:"years,omitempty"`
	Months  Months  `json:"months,omitempty"`
	Weight  Weight  `json:"weight,omitempty"`
}

// PetMap contains stored pets temporarily
var PetMap = make(map[string]Pet)

// ParseJSON marshals JSON data in Pet struct
func (p *Pet) ParseJSON(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &p); err != nil {
		return err
	}
	if p.Name == "" {
		return errors.New("No name defined")
	}
	if p.Species == "" {
		return errors.New("No species defined")
	}
	return nil
}

// GetPet checks if there is a pet with name equals to nama
func GetPet(name string) (Pet, error) {
	for k, v := range PetMap {
		if k == name {
			return v, nil
		}
	}
	return Pet{}, fmt.Errorf("%s is not one of your pets", name)
}

package pet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Pet defines pet attributes
type Pet struct {
	Name    Name    `json:"name"`
	Species Species `json:"species"`
	Breed   Breed   `json:"breed,omitempty"`
	Size    Size    `json:"size,omitempty"`
	Date    Date    `json:"date,omitempty"`
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
		return fmt.Errorf("No name defined")
	}
	if p.Species == "" {
		return fmt.Errorf("No species defined")
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

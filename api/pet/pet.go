package pet

import (
	"fmt"
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

// GetPet checks if there is a pet with name equals to name and returns it
func GetPet(name string) (Pet, error) {
	for k, v := range PetMap {
		if k == name {
			return v, nil
		}
	}
	return Pet{}, fmt.Errorf("%s is not one of your pets", name)
}

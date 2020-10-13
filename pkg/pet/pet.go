package pet

import "fmt"

// Pet defines pet attributes
type Pet struct {
	Name    string  `json:"name"`
	Species Species `json:"Species"`
	Breed   Breed   `json:"breed,omitempty"`
	Size    Size    `json:"size,omitempty"`
	Years   *Years  `json:"years,omitempty"`
	Months  Months  `json:"months,omitempty"`
	Weight  Weight  `json:"weight,omitempty"`
}

var PetMap = make(map[string]Pet)

// Check if there is a pet with name equals to nama
func GetPet(name string) (Pet, error) {
	for k, v := range PetMap {
		if k == name {
			return v, nil
		}
	}
	return Pet{}, fmt.Errorf("%s is not one of your pets", name)
}

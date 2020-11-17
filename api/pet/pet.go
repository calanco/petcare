package pet

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

// CleanPets destroys all the stored pets
func CleanPets() {
	PetMap = make(map[string]Pet)
}

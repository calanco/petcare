package pet

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

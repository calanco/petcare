package pet

// Pet defines pet attributes
type Pet struct {
	Name  string `json:"name"`
	Breed Breed  `json:"breed"`
	Size  Size   `json:"size"`
}

var PetMap = make(map[string]Pet)

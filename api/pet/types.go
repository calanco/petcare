package pet

// Name is a type used by Pet struct. It's mandatory to use
type Name string

// Species is a type used by Pet struct. It's mandatory to use
type Species string

// Possible values of Species type
const (
	DOG Species = "dog"
	CAT         = "cat"
)

// Breed is a type used by Pet struct
type Breed string

// Possible values of Breed type
const (
	BASSOTTO  Breed = "bassotto"
	CHIHUAHUA       = "chihuahua"
)

// Size is a type used by Pet struct
type Size string

// Possible values of Size type
const (
	SMALL      Size = "small"
	MEDIUM          = "medium"
	LARGE           = "large"
	EXTRALARGE      = "extralarge"
)

// Date is a type used by Pet struct
type Date string

// Weight is a type used by Pet struct
type Weight float32

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

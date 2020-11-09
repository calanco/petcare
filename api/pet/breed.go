package pet

import (
	"fmt"
	"strings"
)

// Breed is a type used by Pet struct
type Breed string

// Possible values of Breed type
const (
	BASSOTTO  Breed = "bassotto"
	CHIHUAHUA       = "chihuahua"
)

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

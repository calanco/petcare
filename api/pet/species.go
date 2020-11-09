package pet

import (
	"fmt"
	"strings"
)

// Species is a type used by Pet struct
type Species string

// Possible values of Species type
const (
	DOG Species = "dog"
	CAT         = "cat"
)

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

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

// UnmarshalJSON povides a json unmarshal function to force Species to assume fixed values
func (s *Species) UnmarshalJSON(data []byte) error {
	str := strings.ToLower(strings.Trim(string(data), `"`))
	switch {
	case str == "dog":
		*s = DOG
	case str == "cat":
		*s = CAT
	default:
		return fmt.Errorf("%s not allowed as species", str)
	}
	return nil
}

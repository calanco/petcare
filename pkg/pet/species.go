package pet

import (
	"fmt"
	"strings"
)

// Type used by Pet struct
type Species string

const (
	DOG Species = "dog"
	CAT         = "cat"
)

// Provide a json unmarshal function to force species to assume fixed values
func (s *Species) UnmarshalJSON(data []byte) error {
	str := strings.ToUpper(strings.Trim(string(data), `"`))
	switch {
	case str == "DOG":
		*s = DOG
	case str == "CAT":
		*s = CAT
	default:
		return fmt.Errorf("%s not allowed as species", str)
	}
	return nil
}

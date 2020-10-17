package pet

import (
	"fmt"
	"strings"
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

// UnmarshalJSON provides a json unmarshal function to force size to assume fixed values
func (s *Size) UnmarshalJSON(data []byte) error {
	str := strings.ToUpper(strings.Trim(string(data), `"`))
	switch {
	case str == "SMALL":
		*s = SMALL
	case str == "MEDIUM":
		*s = MEDIUM
	case str == "LARGE":
		*s = LARGE
	case str == "EXTRALARGE":
		*s = EXTRALARGE
	default:
		return fmt.Errorf("%s not allowed as size", str)
	}
	return nil
}

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

// UnmarshalJSON provides a json unmarshal function to force Size to assume fixed values
func (s *Size) UnmarshalJSON(data []byte) error {
	str := strings.ToLower(strings.Trim(string(data), `"`))
	switch {
	case str == "small":
		*s = SMALL
	case str == "medium":
		*s = MEDIUM
	case str == "large":
		*s = LARGE
	case str == "extralarge":
		*s = EXTRALARGE
	default:
		return fmt.Errorf("%s not allowed as size", str)
	}
	return nil
}

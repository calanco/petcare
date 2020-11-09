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

// validateSize checks the validity of the inserted size  and set it to lower case
func validateSize(size *Size) error {
	s := (strings.ToLower(string(*size)))
	if s == "" {
		return nil
	}

	*size = Size(s)
	switch {
	case *size == SMALL:
		return nil
	case *size == MEDIUM:
		return nil
	case *size == LARGE:
		return nil
	case *size == EXTRALARGE:
		return nil
	default:
		return fmt.Errorf("Insert a valid size")
	}
}

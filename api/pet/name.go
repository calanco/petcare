package pet

import (
	"fmt"
	"strings"
)

// Name is a type used by Pet struct. It's mandatory to use
type Name string

// validateName checks the validity of the inserted name and set it to lower case
func validateName(name *Name) error {
	n := strings.ToLower(string(*name))
	if n == "" {
		return fmt.Errorf("Insert a valid name")
	}
	*name = Name(n)
	return nil
}

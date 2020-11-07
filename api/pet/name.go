package pet

import (
	"fmt"
	"strings"
)

// Name is a type used by Pet struct. It's mandatory to use
type Name string

// UnmarshalJSON provides a json unmarshal function to set Name
func (n *Name) UnmarshalJSON(data []byte) error {
	*n = Name(strings.ToLower(strings.Trim(string(data), `"`)))
	fmt.Println(*n)
	return nil
}

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

// UnmarshalJSON provides a json unmarshal function to force breed to assume fixed values
func (b *Breed) UnmarshalJSON(data []byte) error {
	str := strings.ToUpper(strings.Trim(string(data), `"`))
	switch {
	case str == "BASSOTTO":
		*b = BASSOTTO
	case str == "CHIHUAHUA":
		*b = CHIHUAHUA
	default:
		return fmt.Errorf("%s not allowed as breed", str)
	}
	return nil
}

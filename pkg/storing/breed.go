package storing

import (
	"fmt"
	"strings"
)

type Breed string

const (
	BASSOTTO  Breed = "bassotto"
	CHIHUAHUA       = "chihuahua"
)

func (b *Breed) UnmarshalJSON(data []byte) error {
	str := strings.ToUpper(strings.Trim(string(data), `"`))
	switch {
	case str == "BASSOTTO":
		*b = BASSOTTO
	case str == "CHIHUAHUA":
		*b = CHIHUAHUA
	default:
		return fmt.Errorf("Razza sconosciuta: %s", str)
	}
	return nil
}

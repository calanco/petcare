package storing

import (
	"fmt"
	"strings"
)

type Size string

const (
	SMALL      Size = "small"
	MEDIUM          = "medium"
	LARGE           = "large"
	EXTRALARGE      = "extralarge"
)

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
		return fmt.Errorf("Taglia sconosciuta: %s", str)
	}
	return nil
}

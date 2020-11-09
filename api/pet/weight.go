package pet

import (
	"fmt"
)

// Weight is a type used by Pet struct
type Weight float32

// validateWeight checks the validity of the inserted weight
func validateWeight(weight Weight) error {
	w := float32(weight)
	if w == 0 {
		return nil
	}

	if w < 0 || w > 99 {
		return fmt.Errorf("Insert a valid weight")
	}
	return nil
}

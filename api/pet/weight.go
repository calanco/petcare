package pet

import (
	"fmt"
	"strconv"
	"strings"
)

// Weight is a type used by Pet struct
type Weight float32

// UnmarshalJSON provides a json unmarshal function to force weight to assume fixed values
func (w *Weight) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return fmt.Errorf("%s not allowed as weight", str)
	}
	if f <= 0 || f > 99 {
		return fmt.Errorf("%f not allowed as weight", f)
	}
	*w = Weight(f)
	return nil
}

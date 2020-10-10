package pet

import (
	"fmt"
	"strconv"
	"strings"
)

// Type used by Pet struct
type Years int

// Provide a json unmarshal function to force years to assume fixed values
func (y *Years) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)

	i, err := strconv.Atoi(str)
	if err != nil {
		return fmt.Errorf("%s not allowed as years", str)
	}

	if i < 0 || i > 99 {
		return fmt.Errorf("%d not allowed as years", i)
	}

	*y = Years(i)
	return nil
}

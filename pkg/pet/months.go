package pet

import (
	"fmt"
	"strconv"
	"strings"
)

// Type used by Pet struct
type Months int

// Provide a json unmarshal function to force months to assume fixed values
func (m *Months) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)

	i, err := strconv.Atoi(str)
	if err != nil {
		return fmt.Errorf("%s not allowed as months", str)
	}

	if i < 1 || i > 12 {
		return fmt.Errorf("%d not allowed as months", i)
	}

	*m = Months(i)
	return nil
}

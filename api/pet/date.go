package pet

import (
	"fmt"
	"strings"
	"time"
)

// Date is a type used by Pet struct
type Date string

// UnmarshalJSON provides a json unmarshal function to set Date to a proper value
func (d *Date) UnmarshalJSON(data []byte) error {
	layout := "2006-Jan-02"
	date, err := time.Parse(layout, string(strings.Trim(string(data), `"`)))
	if err != nil {
		return err
	}
	*d = Date(fmt.Sprintf("%d-%s-%d", date.Year(), date.Month(), date.Day()))
	return nil
}

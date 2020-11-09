package pet

import (
	"time"
)

// Date is a type used by Pet struct
type Date string

// validateDate checks the validity of the inserted date
func validateDate(date Date) error {
	d := string(date)
	if d == "" {
		return nil
	}

	layout := "2006-Jan-02"
	_, err := time.Parse(layout, d)
	if err != nil {
		return err
	}
	return nil
}

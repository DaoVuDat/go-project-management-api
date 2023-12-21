package util

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

func IsUTCTime(layout string) validation.Rule {
	return validation.By(func(value interface{}) error {
		if timePtr, ok := value.(*time.Time); ok && timePtr != nil {
			// Parsing the time value using the specified layout
			_, err := time.Parse(layout, timePtr.Format(layout))
			if err != nil {
				return fmt.Errorf("must be in UTC time with layout %s", layout)
			}
		}
		return nil
	})

}

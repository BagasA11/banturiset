package validators

import (
	"time"

	tz "github.com/bagasa11/banturiset/timezone"
	v "github.com/go-playground/validator/v10"
)

func ValidateStartTime(fl v.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := tz.GetTime(time.Now())
		if date.Before(today) {
			return false
		}
	}
	return true
}

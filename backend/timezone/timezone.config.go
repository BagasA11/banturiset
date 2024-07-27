package timezone

import (
	"time"
)

const Timezone = "Asia/Jakarta"

var Location *time.Location

func SetLocation(tz string) error {
	var err error
	Location, err = time.LoadLocation(Timezone)
	return err
}

func GetTime(t time.Time) time.Time {
	return t.In(Location)
}

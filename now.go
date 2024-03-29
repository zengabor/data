package data

import "time"

// UTC Now, rounded to microseconds (for datastore)
func Now() time.Time {
	return time.Now().Round(time.Microsecond).UTC()
}

func NowSeconds() time.Time {
	return time.Now().Round(time.Second).UTC()
}

package timeext

import "time"

func Now() time.Time {
	return time.Now().UTC()
}

func NowString() string {
	return time.Now().UTC().Format(time.RFC3339)
}

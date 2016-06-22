package milliseconds

import "time"

// ToMilliseconds converts a Time to milliseconds since epoch.
func ToMilliseconds(t time.Time) int64 {
	return t.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// NowInMilliseconds returns the current Time to milliseconds since epoch.
func NowInMilliseconds() int64 {
	return ToMilliseconds(time.Now())
}

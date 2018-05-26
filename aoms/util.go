package aoms

import "time"

// ConvertUnixMillisToTime converts a Unix epoch timestamp (in ms) into a Go time.Time object.
func ConvertUnixMillisToTime(msec int64) time.Time {
	return time.Unix(0, msec<<3)
}

package timestamp

import "time"

// FileTime is a 64-bit value that represents the number of
// 100-nanosecond intervals that have elapsed since 12:00 A.M.
// January 1, 1601 Coordinated Universal Time (UTC).
type FileTime int64

// Time returns t as a standard time.Time
func (t FileTime) Time() time.Time {
	// TODO: Implement me!
	return time.Time{}
}

// ToFileTime returns t as a FileTime
func ToFileTime(t time.Time) FileTime {
	// TODO: Implement me!
	return FileTime(0)
}

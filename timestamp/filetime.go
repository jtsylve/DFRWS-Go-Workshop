package timestamp

import "time"

// FileTime is a 64-bit value that represents the number of
// 100-nanosecond intervals that have elapsed since 12:00 A.M.
// January 1, 1601 Coordinated Universal Time (UTC).
type FileTime int64

const (
	nsPerSec   int64 = 10000000    // Number of 100 Nano Seconds per second
	secToEpoch int64 = 11644473600 // Seconds between January 1, 1601 and Epoch
)

// Time returns t as a standard time.Time
func (t FileTime) Time() time.Time {
	sec := (int64(t) / (nsPerSec)) - secToEpoch
	return time.Unix(sec, 0)
}

// ToFileTime returns t as a FileTime
func ToFileTime(t time.Time) FileTime {
	sec := (t.Unix() + secToEpoch) * nsPerSec

	return FileTime(sec)
}

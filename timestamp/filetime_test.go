package timestamp

import (
	"testing"
	"time"
)

type convTest struct {
	FileTime FileTime
	UnixTime int64
}

var convTable = []convTest{
	convTest{102054214930000000, -1439052107},   // Monday, May 26, 1924 7:18:13am UTC
	convTest{-102054214930000000, -21849895093}, // Sunday, August 8, 1277 4:41:47pm UTC
	convTest{130835257070000000, 1439052107},    // Saturday, August 8, 2015 4:41:47pm UTC
	convTest{-130835257070000000, -24727999307}, // Wednesday, May 27, 1186 7:18:13am UTC

	convTest{116444735990000000, -1}, // Wednesday, December 31, 1969 11:59:59pm UTC
	convTest{116444736000000000, 0},  // Thursday, January 1, 1970 12:00:00am UTC
	convTest{116444736010000000, 1},  // Thursday, January 1, 1970 12:00:01am UTC
}

func TestFileTime(t *testing.T) {
	now := time.Now()
	ft := ToFileTime(now)

	if ft.Time().Unix() != now.Unix() {
		t.Errorf(
			"Current time conversion failed.  Current Time: %v FileTime: %v\n",
			now.Unix(),
			ft.Time().Unix(),
		)
	}
}

func TestTime(t *testing.T) {
	for _, ft := range convTable {
		if FileTime(ft.FileTime).Time().Unix() != ft.UnixTime {
			t.Errorf(
				"Time() failed. Input: %v Returned: %v Expected: %v\n",
				ft.FileTime,
				FileTime(ft.FileTime).Time().Unix(),
				ft.UnixTime,
			)
		}
	}
}

func TestToFileTime(t *testing.T) {
	for _, ft := range convTable {
		if ToFileTime(time.Unix(ft.UnixTime, 0)) != ft.FileTime {
			t.Errorf(
				"ToFileTime() failed. Input: %v Returned: %v Expected: %v\n ",
				ft.UnixTime,
				ToFileTime(time.Unix(ft.UnixTime, 0)),
				ft.FileTime,
			)
		}
	}
}

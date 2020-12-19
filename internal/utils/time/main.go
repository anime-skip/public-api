package time

import (
	"time"
)

// CurrentTimeMS ...
func CurrentTimeMS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// CurrentTimeSec ...
func CurrentTimeSec() int64 {
	return time.Now().Unix()
}

package date

import (
	"time"
)

func Today() time.Time {
	year, month, day := time.Now().Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	return date
}

func Date(unix int64) int64 {
	year, month, day := time.Unix(unix, 0).Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	return date.Unix()
}

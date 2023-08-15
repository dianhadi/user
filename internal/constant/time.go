package constant

import "time"

const (
	OneHour  = 1 * time.Hour
	OneDay   = 24 * OneHour
	OneWeek  = 7 * OneDay
	OneMonth = 30 * OneDay

	TimeFormat = "2006-01-02 15:04:05.999999 -0700 -0700"
)

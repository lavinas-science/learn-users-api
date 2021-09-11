package dates

import "time"

const (
	dateLayout = "2006-01-02T15:04:05Z"
	dbLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(dateLayout)
}

func GetNowDb() string {
	return GetNow().Format(dbLayout)
}

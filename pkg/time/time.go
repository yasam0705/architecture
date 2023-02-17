package time

import "time"

const (
	Date = "2006-01-02"
)

func DateToStringRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func StringToDateRFC3339(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

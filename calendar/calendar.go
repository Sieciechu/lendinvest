package calendar

import "time"

const (
	layoutISO = "2006-01-02"
)

func NewDate(isoDate string) time.Time {
	t, _ := time.Parse(layoutISO, isoDate)
	return t
}

func GetLastDayOfMonth(date time.Time) int {
	m := date.Month()

	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 2:
		if !isLeapYear(date.Year()) {
			return 28
		}
		return 29
	default:
		return 30
	}
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

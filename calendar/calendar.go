package calendar

import "time"

const (
	layoutISO = "2006-01-02"
)

// NewDate - helper function to easily create
//	date from ISO format YYYY-MM-DD
func NewDate(isoDate string) time.Time {
	t, _ := time.Parse(layoutISO, isoDate)
	return t
}

// GetLastDayOfMonth - returns the last day of the month of given date
// Example:
//	GetLastDayOfMonth(2019-07-14) => july has 31 days, so returns 31
//	GetLastDayOfMonth(2019-02-28) => february has 28 days, so returns 28 (or 29 if year is leap)
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

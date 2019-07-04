package calendar

import (
	"fmt"
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	// given
	isoDates := []string{
		"2019-07-04",
		"2019-01-31",
		"2020-02-29",
	}

	for _, inDate := range isoDates {

		// when-then
		date := NewDate(inDate)
		outDate := fmt.Sprintf("%.4d-%.2d-%.2d", date.Year(), int(date.Month()), date.Day())

		if inDate != outDate {
			t.Errorf("Expected %s, but got %s", inDate, outDate)
		}
	}
}
func TestGetLastDayOfMonth(t *testing.T) {
	cases := []struct {
		inDate      time.Time
		expectedDay int
	}{
		{NewDate("2018-01-01"), 31},
		{NewDate("2020-02-14"), 29},
		{NewDate("2019-11-12"), 30},
	}

	for _, c := range cases {
		lastDayOfMonth := GetLastDayOfMonth(c.inDate)

		if lastDayOfMonth != c.expectedDay {
			t.Errorf("Input month was %.2d, expected %d, but got %d",
				c.inDate.Month(), c.expectedDay, lastDayOfMonth)
		}
	}
}

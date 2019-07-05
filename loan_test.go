package lendinvest

import (
	"testing"
	"time"

	"github.com/Sieciechu/lendinvest/calendar"
)

func TestCalculateNumberOfPaychecks(t *testing.T) {

	date := calendar.NewDate

	cases := []struct {
		start          time.Time
		end            time.Time
		expectedNumber int
	}{
		{date("2015-01-01"), date("2015-01-01"), 1},
		{date("2015-01-31"), date("2015-01-31"), 1},
		{date("2015-01-16"), date("2015-02-01"), 2},
		{date("2015-01-01"), date("2015-02-01"), 2},
		{date("2015-01-01"), date("2015-01-31"), 1},
		{date("2015-10-02"), date("2015-11-15"), 2},
		{date("2015-10-02"), date("2016-11-15"), 14},
	}
	for _, c := range cases {
		got := calculateNumberOfPaychecks(c.start, c.end)
		if got != c.expectedNumber {
			t.Errorf("calculateNumberOfPaychecks(%q, %q), got %d , want %d",
				c.start, c.end, got, c.expectedNumber)
		}
	}
}

func TestGetNextPaymentDate(t *testing.T) {

	date := calendar.NewDate

	cases := []struct {
		start    time.Time
		end      time.Time
		expected time.Time
	}{
		{date("2018-07-01"), date("2018-07-01"), date("2018-07-01")},
		{date("2018-07-04"), date("2018-07-15"), date("2018-07-15")},
		{date("2018-07-04"), date("2018-08-15"), date("2018-08-01")},
		{date("2018-07-31"), date("2018-08-01"), date("2018-08-01")},
		{date("2018-07-07"), date("2018-12-01"), date("2018-08-01")},
	}

	for _, c := range cases {
		got := getNextPaymentDate(c.start, c.end)

		if !c.expected.Equal(got) {
			t.Errorf("Expected %q, but got %q", c.expected, got)
		}
	}
}

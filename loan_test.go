package lendinvest

import (
	"testing"
	"time"

	"github.com/Sieciechu/lendinvest/calendar"
)

func TestCalculateNumberOfPaychecks(t *testing.T) {

	cases := []struct {
		start          time.Time
		end            time.Time
		expectedNumber int
	}{
		{calendar.NewDate("2015-01-01"), calendar.NewDate("2015-01-01"), 1},
		{calendar.NewDate("2015-01-01"), calendar.NewDate("2015-01-31"), 1},
		{calendar.NewDate("2015-10-02"), calendar.NewDate("2015-11-15"), 2},
		{calendar.NewDate("2015-10-02"), calendar.NewDate("2016-11-15"), 14},
	}
	for _, c := range cases {
		got := calculateNumberOfPaychecks(c.start, c.end)
		if got != c.expectedNumber {
			t.Errorf("calculateNumberOfPaychecks(%q, %q), got %d , want %d",
				c.start, c.end, got, c.expectedNumber)
		}
	}
}

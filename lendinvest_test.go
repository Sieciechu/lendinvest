package lendinvest

import "testing"
import "time"

func TestCalculateNumberOfPaychecks(t *testing.T) {

	cases := []struct {
		start          time.Time
		end            time.Time
		expectedNumber int
	}{
		{newDate("2015-01-01"), newDate("2015-01-01"), 1},
		{newDate("2015-01-01"), newDate("2015-01-31"), 1},
		{newDate("2015-10-02"), newDate("2015-11-15"), 2},
		{newDate("2015-10-02"), newDate("2016-11-15"), 14},
	}
	for _, c := range cases {
		got := calculateNumberOfPaychecks(c.start, c.end)
		if got != c.expectedNumber {
			t.Errorf("calculateNumberOfPaychecks(%q, %q), got %d , want %d",
				c.start, c.end, got, c.expectedNumber)
		}
	}
}

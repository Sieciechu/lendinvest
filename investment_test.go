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

func TestGetNearestPeriodEndDate(t *testing.T) {

	date := calendar.NewDate

	cases := []struct {
		start    time.Time
		end      time.Time
		expected time.Time
	}{
		{date("2018-07-01"), date("2018-07-01"), date("2018-07-01")},
		{date("2018-07-04"), date("2018-07-15"), date("2018-07-15")},
		{date("2018-07-04"), date("2018-08-15"), date("2018-07-31")},
		{date("2018-07-31"), date("2018-08-01"), date("2018-07-31")},
		{date("2018-07-07"), date("2018-12-01"), date("2018-07-31")},
	}

	for _, c := range cases {
		got := getNearestPeriodEndDate(c.start, c.end)

		if !c.expected.Equal(got) {
			t.Errorf("Expected %q, but got %q", c.expected, got)
		}
	}
}

func TestCalculateMoneyToPayForPeriod(t *testing.T) {

	date := calendar.NewDate

	cases := []struct {
		money          Cash
		start, end     time.Time
		percent        uint
		expectedIncome Cash
	}{
		{1000, date("2019-10-03"), date("2019-11-15"), 3, 28.06},
		{500, date("2019-10-10"), date("2019-11-15"), 6, 21.29},
	}
	_ = cases

	// given
	for _, c := range cases {
		i := investment{
			investedMoney:            c.money,
			startDate:                c.start,
			endDate:                  c.end,
			monthlyInterestPercetage: c.percent,
		}

		// when
		moneyForFirstMonth := i.calculateMoneyToPayForPeriod(c.start, date("2019-10-31"))

		// then
		if moneyForFirstMonth != c.expectedIncome {
			t.Errorf("Expected %f, but got %f", c.expectedIncome, moneyForFirstMonth)
		}
	}
}

func TestCalculatePaychecks(t *testing.T) {
	date := calendar.NewDate

	// given
	i := investment{
		investedMoney:            1000,
		startDate:                date("2019-10-03"),
		endDate:                  date("2019-11-15"),
		monthlyInterestPercetage: 3,
		paychecks:                nil,
	}

	// when
	i.calculatePaychecks()

	// then
	if n := len(i.paychecks); n != 3 {
		t.Errorf("Expected 3 paychecks, but got %d", n)
	}
	if m := i.paychecks[0].moneyToPay; m != 28.06 {
		t.Errorf("Expected 28.06 for first paycheck, but got %s", m)
	}
	if m := i.paychecks[2].moneyToPay; m != 1000 {
		t.Errorf("Expected to return the invested 1000, but got %s", m)
	}
}

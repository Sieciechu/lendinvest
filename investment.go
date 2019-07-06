package lendinvest

import (
	"math"
	"time"

	"github.com/Sieciechu/lendinvest/calendar"
)

type investment struct {
	investor                 Investor
	investedMoney            Cash
	monthlyInterestPercetage uint
	startDate                time.Time
	endDate                  time.Time
	paychecks                []paycheck
}

type paycheck struct {
	investor      Investor
	periodStart   time.Time
	periodEnd     time.Time
	dateOfPayment time.Time
	moneyToPay    Cash
	paid          bool
	title         string
}

// Factory function to create new investment and calculates future paychecks
// so we know in advance what money we will have to transfer
// when time of paycheck will occur
func newInvestment(i Investor, money Cash, monthlyInterestPercetage uint,
	start time.Time, end time.Time) investment {

	investment := investment{investor: i,
		investedMoney:            money,
		startDate:                start,
		endDate:                  end,
		monthlyInterestPercetage: monthlyInterestPercetage,
		paychecks:                nil}

	investment.calculatePaychecks()

	return investment
}

// Support method for factory newInvestment()
// Creates paychecks according to investment
// 	plus additional last one being the return of invested money
func (investment *investment) calculatePaychecks() {

	investment.paychecks = nil

	count := calculateNumberOfPaychecks(investment.startDate, investment.endDate)

	periodStart := investment.startDate
	end := investment.endDate
	for k := 0; k < count; k++ {

		periodEndDate := getNearestPeriodEndDate(periodStart, end)

		moneyToPay := investment.calculateMoneyToPayForPeriod(periodStart, periodEndDate)
		p := paycheck{investor: investment.investor,
			periodStart: periodStart, periodEnd: periodEndDate, dateOfPayment: periodEndDate.AddDate(0, 0, 1),
			moneyToPay: moneyToPay, paid: false, title: "income from investment"}

		investment.paychecks = append(investment.paychecks, p)

		periodStart = periodEndDate.AddDate(0, 0, 1)
	}

	returnOfInvestedMoney := paycheck{investor: investment.investor,
		periodStart: periodStart, periodEnd: end, dateOfPayment: end,
		moneyToPay: investment.investedMoney, paid: false, title: "return of investment"}

	investment.paychecks = append(investment.paychecks, returnOfInvestedMoney)
}

// Support method for calculatePaychecks()
// Calculates number of paychecks
func calculateNumberOfPaychecks(start, end time.Time) (n int) {
	for end.After(start) || end.Equal(start) {
		n++
		start = start.AddDate(0, 1, (-start.Day() + 1))
	}
	return
}

// Support method for calculatePaychecks()
func getNearestPeriodEndDate(start, end time.Time) time.Time {

	if end.Year() == start.Year() && end.Month() == start.Month() {
		return end
	}

	return start.AddDate(0, 0, (-start.Day() + calendar.GetLastDayOfMonth(start))) // set last day of given month
}

// It calculates how much money should be paid for each payment period (each paycheck)
// Method used in loop as support method for calculatePaychecks()
func (investment *investment) calculateMoneyToPayForPeriod(start, end time.Time) Cash {

	maxDaysInCurrentPaymentMonth := float64(calendar.GetLastDayOfMonth(start))

	// If start date and end date are same, end.Sub() would return 0, but we invested for this one day
	// therefore we must add one day (24 hours)
	actualNumberOfDays := (end.Sub(start).Hours() + 24) / 24.0

	var percent float64 = float64(investment.monthlyInterestPercetage) / 100.0

	result := (float64(investment.investedMoney) * percent / maxDaysInCurrentPaymentMonth) * actualNumberOfDays

	moneyToPay := Cash(math.Round(result*100) / 100) // round float to 2 decimal places

	return moneyToPay
}

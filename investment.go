package lendinvest

import (
	"time"

	"github.com/Sieciechu/lendinvest/calendar"
)

type investment struct {
	investor                 *Investor
	investedMoney            Cash
	monthlyInterestPercetage uint
	startDate                time.Time
	endDate                  time.Time
	paychecks                []paycheck
}

type paycheck struct {
	investor      *Investor
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
func newInvestment(i *Investor, m Cash, monthlyInterestPercetage uint,
	start time.Time, end time.Time) investment {

	investment := investment{investor: i,
		investedMoney: m,
		startDate:     start,
		endDate:       end,
		paychecks:     nil}

	investment.calculatePaychecks()

	return investment
}

// Support method for factory newInvestment()
// Creates paychecks according to investment
func (investment *investment) calculatePaychecks() {

	investment.paychecks = nil
	
	count := calculateNumberOfPaychecks(investment.startDate, investment.endDate)
	
	periodStart := investment.startDate
	end := investment.endDate
	for k := 0; k < count; k++ {

		nextPaymentDate := getNextPaymentDate(periodStart, end)

		moneyToPay := investment.calculateMoneyToPayForPeriod(periodStart, nextPaymentDate)
		p := paycheck{investor: investment.investor,
			periodStart: periodStart, periodEnd: nextPaymentDate, dateOfPayment: nextPaymentDate,
			moneyToPay: moneyToPay, paid: false, title: "income from investment"}

		investment.paychecks = append(investment.paychecks, p)

		periodStart = nextPaymentDate.AddDate(0, 0, 1)
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
func getNextPaymentDate(start, end time.Time) time.Time {

	if end.Year() == start.Year() && end.Month() == start.Month() {
		return end
	}

	return start.AddDate(0, 1, (-start.Day() + 1)) // set date to 1st day of next month
}

// It calculates how much money should be paid for each payment period (each paycheck)
// Method used in loop as support method for calculatePaychecks()
func (investment *investment) calculateMoneyToPayForPeriod(start, end time.Time) Cash {

	maxDaysInCurrentPaymentMonth := float64(calendar.GetLastDayOfMonth(start))
	actualNumberOfDays := end.Sub(start).Hours() / 24.0

	var percent float64 = (1.0 + float64(investment.monthlyInterestPercetage)/100.0)

	var moneyToPay Cash = Cash(
		float64(investment.investedMoney) * percent / maxDaysInCurrentPaymentMonth * actualNumberOfDays)

	return moneyToPay
}

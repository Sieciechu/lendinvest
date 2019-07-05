package lendinvest

import (
	"time"

	"github.com/Sieciechu/lendinvest/calendar"
)

type loan struct {
	start    time.Time
	end      time.Time
	tranches map[trancheID]tranche
}

type trancheID string

type tranche struct {
	id                       trancheID
	maximumInvestment        Cash
	investmentsLeft          Cash
	monthlyInterestPercetage uint
	investment               []investment
}

type investment struct {
	investor                 *Investor
	investedMoney            Cash
	monthlyInterestPercetage uint
	startDate                time.Time
	endDate                  time.Time
	paychecks                []paycheck
}

func newInvestment(i *Investor, m Cash, monthlyInterestPercetage uint,
	start time.Time, end time.Time) investment {

	n := calculateNumberOfPaychecks(start, end)

	investment := investment{investor: i,
		investedMoney: m,
		startDate:     start,
		endDate:       end,
		paychecks:     make([]paycheck, n)}

	investment.calculatePaychecks()

	return investment
}

func (investment *investment) calculatePaychecks() {

	periodStart := investment.startDate
	end := investment.endDate

	for k, size := 0, len(investment.paychecks); k < size; k++ {

		nextPaymentDate := getNextPaymentDate(periodStart, end)

		moneyToPay := investment.calculateMoneyToPayForPeriod(periodStart, nextPaymentDate)
		p := paycheck{investor: investment.investor,
			periodStart: periodStart, periodEnd: nextPaymentDate, dateOfPayment: nextPaymentDate,
			moneyToPay: moneyToPay, paid: false, title: "income from investment"}

		investment.paychecks[k] = p

		periodStart = nextPaymentDate.AddDate(0, 0, 1)
	}

	returnOfInvestedMoney := paycheck{investor: investment.investor,
		periodStart: periodStart, periodEnd: end, dateOfPayment: end,
		moneyToPay: investment.investedMoney, paid: false, title: "return of investment"}

	investment.paychecks = append(investment.paychecks, returnOfInvestedMoney)
}

func (investment *investment) calculateMoneyToPayForPeriod(start, end time.Time) Cash {

	maxDaysInCurrentPaymentMonth := float64(calendar.GetLastDayOfMonth(start))
	actualNumberOfDays := end.Sub(start).Hours() / 24.0

	var percent float64 = (1.0 + float64(investment.monthlyInterestPercetage)/100.0)

	var moneyToPay Cash = Cash(
		float64(investment.investedMoney) * percent / maxDaysInCurrentPaymentMonth * actualNumberOfDays)

	return moneyToPay
}

func getNextPaymentDate(start, end time.Time) time.Time {

	if end.Year() == start.Year() && end.Month() == start.Month() {
		return end
	}

	return start.AddDate(0, 1, (-start.Day() + 1)) // set date to 1st day of next month
}

func calculateNumberOfPaychecks(start, end time.Time) (n int) {
	for end.After(start) || end.Equal(start) {
		n++
		start = start.AddDate(0, 1, (-start.Day() + 1))
	}
	return
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

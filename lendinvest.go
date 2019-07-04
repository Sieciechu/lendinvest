package lendinvest

import (
	"time"

	"github.com/Sieciechu/lendinvest/calendar"
)

type Lendinvest struct {
	loans     []loan
	paychecks []*paycheck
}

type loan struct {
	start    time.Time
	end      time.Time
	tranches map[trancheId]tranche
}

type Cash float64

type Investor interface {
	lendMoney(money Cash) (Cash, error)
	takeMoney(money Cash)
}

type trancheId string

type tranche struct {
	id                       trancheId
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

	paymentDate := investment.startDate
	end := investment.endDate

	for k, size := 0, len(investment.paychecks); k < size; k++ {

		nextPaymentDate := getNextPaymentDate(paymentDate, end)
		numberOfDays := nextPaymentDate.Sub(paymentDate).Hours() / 24.0

		daysInCurrentPaymentMonth := float64(31)

		var percent float64 = (1.0 + float64(investment.monthlyInterestPercetage)/100.0)

		var moneyToPay Cash = Cash(
			float64(investment.investedMoney) * percent / daysInCurrentPaymentMonth * numberOfDays)

		p := paycheck{investor: investment.investor, dateOfPayment: nextPaymentDate,
			moneyToPay: moneyToPay, paid: false, title: "income from investment"}

		investment.paychecks[k] = p

		paymentDate = nextPaymentDate
	}

	returnOfInvestedMoney := paycheck{investor: investment.investor, dateOfPayment: paymentDate,
		moneyToPay: investment.investedMoney, paid: false, title: "return of investment"}

	investment.paychecks = append(investment.paychecks, returnOfInvestedMoney)
}

func getNextPaymentDate(start, end time.Time) time.Time {

	if end.Year() == start.Year() && end.Month() == start.Month() {
		return end.AddDate(0, 0, 0) // AddDate the easiest way to return copy time struct
	}

	return start.AddDate(0, 1, (-start.Day() + 1)) // set date to 1st day of next month
}

func calculateNumberOfPaychecks(start, end time.Time) (n int) {
	if start.Equal(end) {
		return 1
	}
	for i := start; end.After(i); i = i.AddDate(0, 1, 0) {
		n++
	}
	return
}

type paycheck struct {
	investor      *Investor
	dateOfPayment time.Time
	moneyToPay    Cash
	paid          bool
	title         string
}

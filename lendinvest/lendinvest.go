package lendinvest

import (
	"time"
)

const (
	layoutISO = "2006-01-02"
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

func newDate(isoDate string) (time.Time, error) {
	return time.Parse(layoutISO, isoDate)
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
	investor      *Investor
	investedMoney Cash
	startDate     time.Time
	endDate       time.Time
	paychecks     []paycheck
}

type paycheck struct {
	investor      *Investor
	dateOfPayment time.Time
	moneyToPay    Cash
	paid          bool
	title         string
}

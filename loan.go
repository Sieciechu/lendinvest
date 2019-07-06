package lendinvest

import (
	"errors"
	"fmt"
	"time"
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

func (l *loan) makeInvestment(i InvestmentRequest) (investment *investment, err error) {

	ok, err := l.checkInvestmentDates(i.startDate, i.endDate)
	if !ok {
		return
	}

	t := l.tranches[i.tranche]

	investment, err = t.makeInvestment(i)
	return
}

func (l *loan) checkInvestmentDates(investmentStart, investmentEnd time.Time) (ok bool, err error) {

	if investmentStart.Before(l.start) {
		err = errors.New("Cannot invest before loan start date")
		return
	}

	if investmentEnd.After(l.end) {
		err = errors.New("Cannot invest after loan end date")
		return
	}

	return true, nil
}

func (t *tranche) makeInvestment(i InvestmentRequest) (*investment, error) {

	m, err := i.investor.LendMoney(i.investedMoney)
	if err != nil {
		return nil, err
	}

	if m > t.investmentsLeft {
		err = fmt.Errorf("Current maximum investment to tranche '%s' is %s, while you wanted to invest %s",
			t.id, t.investmentsLeft, m)
		return nil, err
	}

	t.investmentsLeft -= m

	t.investment = append(t.investment,
		 newInvestment(i.investor, m, t.monthlyInterestPercetage, i.startDate, i.endDate))

	createdInvestment := &t.investment[len(t.investment)-1]

	return createdInvestment, nil
}

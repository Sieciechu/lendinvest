package lendinvest

import (
	"errors"
	"fmt"
	"time"
)

// loan struct contains information about loan and it's tranches
type loan struct {
	start    time.Time
	end      time.Time
	tranches map[TrancheID]tranche
}

// simple type for TrancheID
type TrancheID string

// Tranche contains information about a Tranche
//	* maximumInvestment is the initial "capacity"
//  * investmentsLeft is actual "capacity" left to invest
//	* investments - contain information about each investment made on the Tranche;
//		they are generated automatically on makeInvestment()
type tranche struct {
	id                        TrancheID
	maximumInvestment         Cash
	investmentsLeft           Cash
	monthlyInterestPercentage uint
	investments               []investment
}

// loan.makeInvestment - makes investment in the loan's Tranche according to investment request.
func (l *loan) makeInvestment(i InvestmentRequest) (investment *investment, err error) {

	ok, err := l.checkInvestmentDates(i.StartDate, i.EndDate)
	if !ok {
		return
	}

	t, exists := l.tranches[i.Tranche]
	if !exists {
		err = fmt.Errorf("There is no such tranche like '%s'", i.Tranche)
		return
	}

	investment, err = t.makeInvestment(i)
	return
}

// Support method for loan.makeInvestment. Checks if given dates are between loan's start-end dates
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

// Tranche.makeInvestment - makes investment in the Tranche according to the investment request
//	After each investment the "investmentsLeft" is lowered by invested money, so total investments
//	cannot be larger than maximumInvestment
func (t *tranche) makeInvestment(i InvestmentRequest) (*investment, error) {

	m, err := i.Inv.LendMoney(i.MoneyToInvest)
	if err != nil {
		return nil, err
	}

	if m > t.investmentsLeft {
		err = fmt.Errorf("Current maximum investment to Tranche '%s' is %s, while you wanted to invest %s",
			t.id, t.investmentsLeft, m)
		return nil, err
	}

	t.investmentsLeft -= m

	t.investments = append(t.investments,
		newInvestment(i.Inv, m, t.monthlyInterestPercentage, i.StartDate, i.EndDate))

	createdInvestment := &t.investments[len(t.investments)-1]

	return createdInvestment, nil
}

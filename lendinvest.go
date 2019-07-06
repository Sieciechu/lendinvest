package lendinvest

import (
	"fmt"
	"time"
)

// Lendinvest struct containing information about loans in which
//	investors can invest and information about future paychecks - to know
// 	when, to whom and how much cash should investor be paid.
//
//	Important: paychecks are generated automatically (by Lendinvest.MakeInvestment method)
type Lendinvest struct {
	loans     []loan
	paychecks []*paycheck
}

// Cash - simple type for representing money/cash
type Cash float64

func (c Cash) String() string {
	return fmt.Sprintf("%.2f", c)
}

// Investor - anyone who implements this interface can invest in Lendinvest
type Investor interface {
	LendMoney(money Cash) (Cash, error)
	TakeMoney(money Cash)
}

// InvestmentRequest is the DTO struct containing neccessary data to make investment request,
// 	so we know: who invests, how much, the target, investment start date and investment end date
type InvestmentRequest struct {
	investor      Investor
	moneyToInvest Cash
	loanID        int
	tranche       trancheID
	startDate     time.Time
	endDate       time.Time
}

// MakeInvestment - method to make investment according to given investment request
//	If something goes wrong, it returns false and error message
func (l *Lendinvest) MakeInvestment(i InvestmentRequest) (ok bool, err error) {

	loan := l.loans[i.loanID]

	investment, err := loan.makeInvestment(i)
	if nil != err {
		for _, p := range investment.paychecks {
			l.paychecks = append(l.paychecks, &p)
		}
	}
	return
}

// PayPaychecks - pay to assigned investors paychecks for given date
func (l *Lendinvest) PayPaychecks(date time.Time) {
	for i := range l.paychecks {
		if !l.paychecks[i].dateOfPayment.Equal(date) || true == l.paychecks[i].paid {
			continue
		}
		l.paychecks[i].investor.TakeMoney(l.paychecks[i].moneyToPay)
		l.paychecks[i].paid = true
	}
}

package lendinvest

import (
	"fmt"
	"time"
)

// Lendinvest struct containing information about loans in which
//	investors can invest and information about future paychecks - to know
// 	when, to whom and how much cash should Inv be paid.
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
	Inv           Investor
	MoneyToInvest Cash
	LoanID        int
	Tranche       TrancheID
	StartDate     time.Time
	EndDate       time.Time
}

// MakeInvestment - method to make investment according to given investment request
//	If something goes wrong, it returns false and error message
func (li *Lendinvest) MakeInvestment(i InvestmentRequest) (ok bool, err error) {

	loan := li.loans[i.LoanID]

	investment, err := loan.makeInvestment(i)
	if nil != err {
		return
	}

	for i := range investment.paychecks {
		li.paychecks = append(li.paychecks, &investment.paychecks[i])
	}
	return
}

// PayPaychecks - pay to assigned investors paychecks for given date
func (li *Lendinvest) PayPaychecks(date time.Time) {
	for i := range li.paychecks {
		if !li.paychecks[i].dateOfPayment.Equal(date) || true == li.paychecks[i].paid {
			continue
		}
		li.paychecks[i].investor.TakeMoney(li.paychecks[i].moneyToPay)
		li.paychecks[i].paid = true
	}
}

// AddLoan add loan to lendinvest
func (li *Lendinvest) AddLoan(l loan) {
	li.loans = append(li.loans, l)
}

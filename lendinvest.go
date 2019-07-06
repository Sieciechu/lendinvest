package lendinvest

import (
	"fmt"
	"time"
)

type Lendinvest struct {
	loans     []loan
	paychecks []*paycheck
}

type Cash float64

func (c Cash) String() string {
	return fmt.Sprintf("%.2f", c)
}

type Investor interface {
	LendMoney(money Cash) (Cash, error)
	TakeMoney(money Cash)
}

type InvestmentRequest struct {
	investor      Investor
	investedMoney Cash
	loanID        int
	tranche       trancheID
	startDate     time.Time
	endDate       time.Time
}

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

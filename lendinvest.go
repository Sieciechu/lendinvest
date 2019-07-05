package lendinvest

import "fmt"

type Lendinvest struct {
	loans     []loan
	paychecks []*paycheck
}

type Cash float64

func (c Cash) String() string {
	return fmt.Sprintf("%.2f", c)
}

type Investor interface {
	lendMoney(money Cash) (Cash, error)
	takeMoney(money Cash)
}

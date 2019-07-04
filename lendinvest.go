package lendinvest

type Lendinvest struct {
	loans     []loan
	paychecks []*paycheck
}

type Cash float64

type Investor interface {
	lendMoney(money Cash) (Cash, error)
	takeMoney(money Cash)
}

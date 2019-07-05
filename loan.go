package lendinvest

import (
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


package lendinvest

import (
	"github.com/Sieciechu/lendinvest/calendar"
)

// LoanBuilder struct (as name says) for building loans
type LoanBuilder struct {
	l loan
}

// NewLoan inits new loan with start and end date
//	start and end date must be ISO format: YYYY-MM-DD
func (lb *LoanBuilder) NewLoan(start, end string) {
	lb.l = loan{
		start:    calendar.NewDate(start),
		end:      calendar.NewDate(end),
		tranches: make(map[TrancheID]tranche)}
}

// AddTranche - adds Tranche to the loan
func (lb *LoanBuilder) AddTranche(tID TrancheID, maxInvestment Cash, monthlyPercentIncome uint) {
	lb.l.tranches[tID] = tranche{
		id:                        tID,
		maximumInvestment:         maxInvestment,
		investmentsLeft:           maxInvestment,
		monthlyInterestPercentage: monthlyPercentIncome}
}

// Create - returns copies of defined loan
func (lb *LoanBuilder) Create() loan {
	return lb.l
}

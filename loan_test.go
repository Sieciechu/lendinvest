package lendinvest

import (
	"errors"
	"testing"

	"github.com/Sieciechu/lendinvest/calendar"
)

var date = calendar.NewDate

type InvestorMock struct{}

func (f *InvestorMock) LendMoney(money Cash) (Cash, error) {
	return 0, errors.New("Cannot lend money")
}
func (f *InvestorMock) TakeMoney(money Cash) {}

func TestWhenInvestorCannotLendMoneyThenInvestmentShouldNotBeMade(t *testing.T) {

	// given
	john := InvestorMock{}

	trancheA := tranche{
		id:                        "A",
		maximumInvestment:         2000,
		investmentsLeft:           2000,
		monthlyInterestPercentage: 3,
		investments:               nil,
	}

	investmentRequest := InvestmentRequest{
		&john, 2000, 0, "A", date("2019-01-01"), date("2019-01-31"),
	}

	// when
	investment, err := trancheA.makeInvestment(investmentRequest)

	// then
	if nil == err || nil != investment {
		t.Errorf("Expected error, but got investment")
	}
}

func TestWhenInvestmentExceedsTrancheAvailableInvestmentsItShouldBeNotAllowedToInvest(t *testing.T) {
	// given
	john := User{1000}

	trancheA := tranche{
		id:                        "A",
		maximumInvestment:         2000,
		investmentsLeft:           30,
		monthlyInterestPercentage: 3,
		investments:               nil,
	}

	investmentRequest := InvestmentRequest{
		&john, 1000, 0, "A", date("2019-01-01"), date("2019-01-31"),
	}

	// when
	investment, err := trancheA.makeInvestment(investmentRequest)

	// then
	if nil == err || nil != investment {
		t.Errorf("Expected error, but got investment")
	}
}

func TestTrancheMakeInvestment(t *testing.T) {
	// given
	john := User{1000}

	trancheA := tranche{
		id:                        "A",
		maximumInvestment:         2000,
		investmentsLeft:           1500,
		monthlyInterestPercentage: 3,
		investments:               nil,
	}

	investmentRequest := InvestmentRequest{
		&john, 1000, 0, "A", date("2019-01-01"), date("2019-01-31"),
	}

	// when
	investment, err := trancheA.makeInvestment(investmentRequest)

	// then
	if nil != err || nil == investment {
		t.Errorf("Expected investment, but got error")
	}
	if investment.investedMoney != investmentRequest.moneyToInvest {
		t.Errorf("Expected to invest %s money, but %s was invested",
			investmentRequest.moneyToInvest, investment.investedMoney)
	}
	if investment.investor != investmentRequest.investor {
		t.Errorf("Investition was made by someone else")
	}
	if investment.startDate != investmentRequest.startDate {
		t.Errorf("Expected to start invest on %s, but is started on %s",
			investmentRequest.startDate, investment.startDate)
	}
	if investment.endDate != investmentRequest.endDate {
		t.Errorf("Expected to end invest on %s, but is ended on %s",
			investmentRequest.endDate, investment.endDate)
	}
}

func TestCheckInvestmentDates(t *testing.T) {
	// given
	l := loan{start: date("2019-01-01"), end: date("2019-02-16")}

	cases := []struct {
		start          string
		end            string
		expectedResult bool
	}{
		{"2019-01-01", "2019-01-01", true},
		{"2019-01-01", "2019-01-17", true},
		{"2019-01-01", "2019-02-16", true},

		{"2019-01-01", "2019-02-17", false},
		{"2018-12-31", "2019-02-16", false},
		{"2018-12-31", "2019-02-17", false},
	}

	for _, c := range cases {

		// when
		ok, _ := l.checkInvestmentDates(date(c.start), date(c.end))

		// then
		if ok != c.expectedResult {
			t.Errorf("Expected %v, but got %v", c.expectedResult, ok)
		}
	}
}

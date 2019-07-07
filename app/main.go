package main

import (
	"github.com/Sieciechu/lendinvest"
	"github.com/Sieciechu/lendinvest/calendar"

	"fmt"
)

func main() {
	li := lendinvest.Lendinvest{}
	lb := lendinvest.LoanBuilder{}

	lb.NewLoan("2019-10-01", "2019-11-15")
	lb.AddTranche("A", 1000, 3)
	lb.AddTranche("B", 1000, 6)
	loan := lb.Create()

	li.AddLoan(loan)
	fmt.Println(`Scenario
- Given a loan (start 2019-10-01 end 2019-11-15).
- Given the loan has 2 tranches called A and B (3% and 6% monthly interest rate) each with
1,000 pounds amount available.`)

	someUser := lendinvest.User{Money: lendinvest.Cash(1000.0)}
	fmt.Printf("At the beginning user has %s money\n", someUser.Money)

	_, err := li.MakeInvestment(lendinvest.InvestmentRequest{
		Inv:           &someUser,
		MoneyToInvest: 1000,
		LoanID:        0,
		Tranche:       "A",
		StartDate:     calendar.NewDate("2019-10-03"),
		EndDate:       calendar.NewDate("2019-11-15")})

	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Printf("The user is making investment of 1000 on 2019-10-03")
	fmt.Printf("Just after making the investment the user has %s money\n", someUser.Money)

	li.PayPaychecks(calendar.NewDate("2019-11-01"))
	fmt.Printf("At the 1st of November the user got income of %s money\n", someUser.Money)

}

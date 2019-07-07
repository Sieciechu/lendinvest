package lendinvest

import "errors"

// User - a type who will be able to invest
type User struct {
	Money Cash
}

// Implementation of the Investor interface
func (u *User) LendMoney(money Cash) (Cash, error) {
	if money > u.Money {
		return 0, errors.New("User has not enough money to lend")
	}
	u.Money -= money

	return money, nil
}

// Implementation of the Investor interface
func (u *User) TakeMoney(money Cash) {
	u.Money += money
}

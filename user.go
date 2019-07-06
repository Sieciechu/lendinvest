package lendinvest

import "errors"

// User - a type who will be able to invest
type User struct {
	money Cash
}

// Implementation of the Investor interface
func (u *User) LendMoney(money Cash) (Cash, error) {
	if money > u.money {
		return 0, errors.New("User has not enough money to lend")
	}
	u.money -= money

	return money, nil
}

// Implementation of the Investor interface
func (u *User) TakeMoney(money Cash) {
	u.money += money
}

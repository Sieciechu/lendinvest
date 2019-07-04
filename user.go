package lendinvest

import "errors"

type User struct {
	money Cash
}

func (u *User) LendMoney(money Cash) (Cash, error) {
	if money > u.money {
		return 0, errors.New("User has not enough money to lend")
	}
	u.money -= money

	return money, nil
}
func (u *User) TakeMoney(money Cash) {
	u.money += money
}

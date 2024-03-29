package lendinvest

import (
	"testing"
)

func TestWhenUserHasEnoughMoneyMoneyIsLent(t *testing.T) {
	u := User{Money: 7.0}

	lentMoney, err := u.LendMoney(5.0)

	if nil != err {
		t.Errorf("Wanted to lend %f money, but got error '%s'", 5.0, err)
	}
	if lentMoney != 5.0 {
		t.Errorf("Wanted to lend %f money, but got %f", 5.0, lentMoney)
	}
	if u.Money != 2.0 {
		t.Errorf("After lending %f money, user should have left with %f, but has %f", 5.0, 2.0, u.Money)
	}

}
func TestWhenUserHasNotEnoughMoneyMoneyShouldNotBeLent(t *testing.T) {
	u := User{Money: 2.0}

	_, err := u.LendMoney(5.0)

	if nil == err {
		t.Errorf("Wanted to lend more money than user has, expected error, but got nothing")
	}
}

func TestWhenUserTakesMoneyThenHeHasMoreMoney(t *testing.T) {
	u := User{Money: 2.0}

	u.TakeMoney(0.5)

	if 2.5 != u.Money {
		t.Errorf("User took %f money and should have %f, but has %f", 0.5, 2.5, u.Money)
	}
}

func TestUserShouldBeInvestor(t *testing.T) {

	var _ Investor = &User{Money: 2.0}
}

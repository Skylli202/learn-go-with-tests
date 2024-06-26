package arrays

import (
	"testing"

	"github.com/Skylli202/learn-go-with-tests/generics"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	generics.AssertEqual(t, newBalanceFor(riya), 200)
	generics.AssertEqual(t, newBalanceFor(chris), 0)
	generics.AssertEqual(t, newBalanceFor(adil), 175)
}

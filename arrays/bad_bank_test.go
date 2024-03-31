package arrays

import (
	"testing"

	"github.com/Skylli202/learn-go-with-tests/generics"
)

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}

	generics.AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	generics.AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	generics.AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}
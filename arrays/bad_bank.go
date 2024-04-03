package arrays

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewAccount(name string, balance float64) Account {
	return Account{Name: name, Balance: balance}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if a.Name == transaction.From {
		a.Balance = a.Balance - transaction.Sum
	}
	if a.Name == transaction.To {
		a.Balance = a.Balance + transaction.Sum
	}

	return a
}

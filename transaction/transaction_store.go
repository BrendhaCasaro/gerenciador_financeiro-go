package transaction

type TransactionStore struct {
	store []Transaction
}

var ts TransactionStore

func Insert(transaction Transaction) {
	ts.store = append(ts.store, transaction)
}

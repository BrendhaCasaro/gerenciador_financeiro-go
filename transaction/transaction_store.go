package transaction

import (
	"fmt"
)

type TransactionStore struct {
	store []Transaction
}

func (ts *TransactionStore) Insert(transaction Transaction) {
	ts.store = append(ts.store, transaction)
	fmt.Printf("%+v\n", ts.store)
}

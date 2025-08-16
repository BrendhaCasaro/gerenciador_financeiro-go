package transaction

import (
	"fmt"
	"github.com/google/uuid"
)

type TransactionStore struct {
	store []Transaction
}

func (ts *TransactionStore) Insert(transaction Transaction) {
	ts.store = append(ts.store, transaction)
	fmt.Printf("%+v\n", ts.store)
}

func (ts *TransactionStore) SearchById(id uuid.UUID) Transaction {
	for _, transaction := range ts.store {
		if transaction.id == id {
			// fmt.Printf("%T\n", transaction.id)
			// fmt.Printf("%T\n", id)
			return transaction
		}
	}

	return Transaction{}
}

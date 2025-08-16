package transaction

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type TransactionStore struct {
	store []Transaction
}

func (ts *TransactionStore) MarshalJson(transaction Transaction) ([]byte, error) {
	return json.Marshal(transaction)
}

func (ts *TransactionStore) Insert(transaction Transaction) {
	ts.store = append(ts.store, transaction)
	fmt.Printf("%+v\n", ts.store)
}

func (ts *TransactionStore) SearchById(id uuid.UUID) Transaction {
	for _, transaction := range ts.store {
		if transaction.Id == id {
			return transaction
		}
	}

	return Transaction{}
}

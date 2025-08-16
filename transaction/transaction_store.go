package transaction

import (
	"encoding/json"
	"errors"
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
}

func (ts *TransactionStore) SearchById(id uuid.UUID) (*Transaction, error) {
	for _, transaction := range ts.store {
		if transaction.Id == id {
			return &transaction, nil
		}
	}

	return nil, errors.New("Transaction not found")
}

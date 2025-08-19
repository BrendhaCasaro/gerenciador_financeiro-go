package transaction

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type TransactionStore struct {
	store []*Transaction
}

func (ts *TransactionStore) MarshalJson() ([]byte, error) {
	return json.Marshal(ts.store)
}

func (ts *TransactionStore) Insert(transaction *Transaction) {
	ts.store = append(ts.store, transaction)
}

func (ts *TransactionStore) SearchById(id uuid.UUID) (*Transaction, error) {
	for _, transaction := range ts.store {
		if transaction.Id == id {
			return transaction, nil
		}
	}

	return nil, errors.New("Transaction not found")
}

func (ts *TransactionStore) TotalAmount() float64 {
	acc := 0.0
	for _, transaction := range ts.store {
		if transaction.DeletedAt.IsZero() {
			acc += transaction.Value
		}
	}
	return acc
}

func (ts *TransactionStore) SoftDelete(id uuid.UUID) error {
	transaction, err := ts.SearchById(id)
	if err != nil {
		return fmt.Errorf("soft delete failed: %v", err)
	}

	transaction.Delete()

	return nil
}

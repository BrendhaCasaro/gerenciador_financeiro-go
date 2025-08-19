package transaction

import (
	"encoding/json"
	"errors"
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
}

func (ts *TransactionStore) SearchById(id uuid.UUID) (*Transaction, error) {
	for i := range ts.store {
		if ts.store[i].Id == id {
			return &ts.store[i], nil
		}
	}

	return nil, errors.New("Transaction not found")
}

func (ts *TransactionStore) TotalAmount() float64 {
	acc := 0.0
	for _, transaction := range ts.store {
		acc += transaction.Value
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

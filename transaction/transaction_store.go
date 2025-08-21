package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

func (ts *TransactionStore) HardDelete(id uuid.UUID) {
	for i := range ts.store {
		if ts.store[i].Id == id {
			ts.store[i] = ts.store[len(ts.store)-1]
			ts.store = ts.store[:len(ts.store)-1]
			break
		}
	}
}

func (ts *TransactionStore) EditById(id uuid.UUID, fields TransactionFields) {
	type TransactionFields struct {
		name        string
		description string
		value       float64
		realizedAt  time.Time
	}

	tx, err := ts.SearchById(id)

	if err != nil {
		fmt.Errorf("Have a error: %v", err)
		return
	}
}

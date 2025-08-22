package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

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

func (ts *TransactionStore) EditById(id uuid.UUID, uft UpdateFieldsTransaction) {
	tx, err := ts.SearchById(id)
	if err != nil {
		fmt.Printf("Have a error: %v", err)
		return
	}

	tx.Update(uft)
}

func (ts *TransactionStore) ExpensesAmount() float64 {
	acc := 0.0
	for _, transaction := range ts.store {
		if transaction.DeletedAt.IsZero() && transaction.Value < 0.0 {
			acc += transaction.Value
		}
	}
	return acc
}

func (ts *TransactionStore) IncomeAmount() float64 {
	acc := 0.0
	for _, transaction := range ts.store {
		if transaction.DeletedAt.IsZero() && transaction.Value > 0.0 {
			acc += transaction.Value
		}
	}
	return acc
}

func (ts *TransactionStore) SearchByName(name string) ([]*Transaction, error) {
	var results []*Transaction
	for _, transaction := range ts.store {
		if strings.Contains(transaction.Name, name) {
			results = append(results, transaction)
		}
	}

	if len(results) == 0 {
		return nil, errors.New("Transaction not found")
	}

	return results, nil
}

func (ts *TransactionStore) FilterByValue(value float64) {}

func (ts *TransactionStore) FilterByType() {}

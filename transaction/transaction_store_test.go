package transaction

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	ts := TransactionStore{}
	ts.Insert(NewTransaction("teste", "teste", 100, time.Now()))

	if ts.store == nil {
		t.Fatalf("The fields of transaction should not be nil")
	}
}

func TestSearchById(t *testing.T) {
	ts := TransactionStore{}

	ts.Insert(NewTransaction("teste", "teste", 100, time.Now()))
	ts.Insert(NewTransaction("teste2", "teste2", 200, time.Now()))

	transactionSearch := ts.store[1].id
	transactionFound := ts.SearchById(transactionSearch)

	if transactionFound.id != transactionSearch {
		t.Fatalf("The ID founded is different of ID/transaction wanted")
	}
}

package transaction

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	ts := TransactionStore{}
	ts.Insert(NewTransaction("teste", "teste", 100, time.Now()))

	if len(ts.store) == 0 {
		t.Fatalf("The fields of transaction should not be nil")
	}
}

func TestSearchById(t *testing.T) {
	ts := TransactionStore{}

	ts.Insert(NewTransaction("teste", "teste", 100, time.Now()))
	ts.Insert(NewTransaction("teste2", "teste2", 200, time.Now()))

	transactionSearch := ts.store[1].Id
	transactionFound, err := ts.SearchById(transactionSearch)
	if err != nil {
		t.Fatalf("Have a error")
	}

	if transactionFound.Id != transactionSearch {
		t.Fatalf("The ID founded is different of ID/transaction wanted")
	}
}

func TestMarshalJson(t *testing.T) {
	ts := TransactionStore{}
	transaction := NewTransaction("teste", "teste", 100, time.Now())
	transactionJson, err := ts.MarshalJson(transaction)

	if transactionJson == nil {
		t.Fatalf("The transaction was not converted for a JSON: %v", err)
	}
}

func TestTotalAmount(t *testing.T) {
	ts := TransactionStore{}

	ts.Insert(NewTransaction("teste1", "teste1", 100, time.Now()))
	ts.Insert(NewTransaction("teste2", "teste2", 200, time.Now()))

	if ts.TotalAmount() != 300 {
		t.Fatalf("the total received was not what was expected")
	}
}

func TestSoftDelete(t *testing.T) {
	ts := TransactionStore{}
	ts.Insert(NewTransaction("teste1", "teste1", 100, time.Now()))

	ts.SoftDelete(ts.store[0].Id)

	if ts.store[0].DeletedAt.IsZero() {
		t.Fatalf("The function didn't change the time of DeletedAt")
	}
}

package transaction

import (
	"fmt"
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

	transactionSearch := ts.store[1].Id
	transactionFound := ts.SearchById(transactionSearch)

	if transactionFound.Id != transactionSearch {
		t.Fatalf("The ID founded is different of ID/transaction wanted")
	}
}

func TestMarshalJson(t *testing.T) {
	ts := TransactionStore{}
	transaction := NewTransaction("teste", "teste", 100, time.Now())
	transactionJson, err := ts.MarshalJson(transaction)

	if transactionJson == nil {
		t.Fatalf("The transaction was not converted for a JSON")
		fmt.Println(err)
	} else {
		fmt.Println(string(transactionJson))
	}
}

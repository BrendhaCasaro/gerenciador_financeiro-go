package transaction

import (
	"fmt"
	"testing"
	"time"
)

func TestMarshalJson(t *testing.T) {
	ts := TransactionStore{}
	tx := NewTransaction("teste", "teste", 100, time.Now())
	ts.Insert(tx)

	storeJson, err := ts.MarshalJson()

	fmt.Println(string(storeJson))

	if err != nil {
		t.Fatalf("The slice was not converted for a JSON: %v", err)
	}
}

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
		t.Fatalf("Have a error: %v", err)
	}

	if transactionFound.Id != transactionSearch {
		t.Fatalf("The ID founded is different of ID/transaction wanted")
	}
}

func TestTotalAmount(t *testing.T) {
	ts := TransactionStore{}

	ts.Insert(NewTransaction("teste1", "teste1", -100, time.Now()))
	ts.Insert(NewTransaction("teste2", "teste2", 200, time.Now()))
	tx := NewTransaction("teste3", "teste3", 100, time.Now())
	ts.Insert(tx)

	tx.Delete()

	if ts.TotalAmount() != 100 {
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

func TestHardDelete(t *testing.T) {
	ts := TransactionStore{}
	ts.Insert(NewTransaction("teste1", "teste1", 100, time.Now()))
	ts.Insert(NewTransaction("teste2", "teste2", 200, time.Now()))
	tx := NewTransaction("teste3", "teste3", 100, time.Now())
	ts.Insert(tx)

	ts.HardDelete(tx.Id)

	if len(ts.store) != 2 {
		t.Fatalf("The function didn't remove the transaction")
	}
}

func TestEditById(t *testing.T) {
	ts := TransactionStore{}
	tx := NewTransaction("teste3", "teste3", 100, time.Now())
	ts.Insert(tx)

	ts.EditById(tx.Id, UpdateFieldsTransaction{
		Name:  "Compra",
		Value: 200,
	})

	if tx.Name != "Compra" && tx.Value != 200 {
		t.Fatalf("The function not change the name and the value of transaction")
	}
}

func TestExpensesAmount(t *testing.T) {

}

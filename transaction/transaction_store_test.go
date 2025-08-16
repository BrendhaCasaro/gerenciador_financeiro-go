package transaction

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	ts := TransactionStore{}
	ts.Insert(NewTransaction("teste", "teste", 100, time.Now()))

	if ts.store == nil {
		t.Fatalf("The fields of transaction don't should not be nil")
	}
}

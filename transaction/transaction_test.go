package transaction

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestTransaction(t *testing.T) {
	transaction := NewTransaction("teste", "teste", 244, time.Now())

	if transaction.Id == uuid.Nil {
		t.Fatalf("Expected id have a something different of nil")
	}
}

// testar name, campos que recebem valores de outras de fora

func TestDeleteTransaction(t *testing.T) {
	transaction := NewTransaction("teste", "teste", 244, time.Now())

	transaction.Delete()

	if transaction.DeletedAt.IsZero() {
		t.Fatalf("The function Delete didn't work")
	}
}

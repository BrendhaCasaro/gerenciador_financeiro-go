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

func TestUpdate(t *testing.T) {
	transaction := NewTransaction("teste", "teste", 244, time.Now().AddDate(-2, -2, -2))
	now := time.Now()

	transaction.Update(UpdateFieldsTransaction{
		Name:       "Compra",
		Value:      377,
		RealizedAt: now,
	})

	if transaction.Name != "Compra" {
		t.Error("Field name has not been changed")
	}

	if transaction.Value != 377 {
		t.Error("Field value has not been changed")
	}

	if transaction.RealizedAt != now {
		t.Error("Field RealizedAt has not been changed")
	}
}

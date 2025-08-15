package transaction

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestTransaction(t *testing.T) {
	transaction := NewTransaction("teste", "teste", 244, time.Now())

	if transaction.id == uuid.Nil {
		t.Fatalf("Expected id have a something different of nil")
	}
}

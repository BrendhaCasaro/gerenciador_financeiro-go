package transaction

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestTransaction(t *testing.T) {
	transaction := NewTransaction("teste", "teste", 244, time.Now())

	if transaction.Id == uuid.Nil {
		t.Fatalf("Expected id have a something different of nil")
	}
}

// testar name, campos que recebem valores de outras de fora
// testar delete

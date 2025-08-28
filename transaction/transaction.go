package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	insertedAt  time.Time
	RealizedAt  time.Time `json:"realized_at"`
	deletedAt   time.Time
}

func NewTransaction(name string, description string, value float64, realizedAt time.Time) *Transaction {
	return &Transaction{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
		Value:       value,
		insertedAt:  time.Now(),
		RealizedAt:  realizedAt,
	}
}

func (t *Transaction) Delete() {
	t.deletedAt = time.Now()
}

type UpdateFieldsTransaction struct {
	Name        string
	Description *string
	Value       float64
	RealizedAt  time.Time
}

func (t *Transaction) Update(uft UpdateFieldsTransaction) {
	if uft.Name != "" {
		t.Name = uft.Name
	}

	if uft.Description != nil {
		t.Description = *uft.Description
	}

	if uft.Value != 0.0 {
		t.Value = uft.Value
	}

	if !uft.RealizedAt.IsZero() {
		t.RealizedAt = uft.RealizedAt
	}
}

type TransactionType int

const (
	Income TransactionType = iota
	Expense
)

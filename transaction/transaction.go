package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id          uuid.UUID
	Name        string
	Description string
	Value       float64
	InsertedAt  time.Time
	RealizedAt  time.Time
	DeletedAt   time.Time
}

func NewTransaction(name string, description string, value float64, realizedAt time.Time) *Transaction {
	return &Transaction{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
		Value:       value,
		InsertedAt:  time.Now(),
		RealizedAt:  realizedAt,
	}
}

func (t *Transaction) Delete() {
	t.DeletedAt = time.Now()
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

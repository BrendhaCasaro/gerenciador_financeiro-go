package transaction

import (
	"github.com/google/uuid"
	"time"
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

func NewTransaction(name string, description string, value float64, realizedAt time.Time) Transaction {
	return Transaction{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
		Value:       value,
		InsertedAt:  time.Now(),
		RealizedAt:  realizedAt,
	}
}

func Delete() Transaction {
	return Transaction{
		DeletedAt: time.Now(),
	}
}

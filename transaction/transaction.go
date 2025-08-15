package transaction

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	id          uuid.UUID
	name        string
	description string
	value       float64
	insertedAt  time.Time
	realizedAt  time.Time
	deletedAt   time.Time
}

func NewTransaction(name string, description string, value float64, realizedAt time.Time) Transaction {
	return Transaction{
		id:          uuid.New(),
		name:        name,
		description: description,
		value:       value,
		insertedAt:  time.Now(),
		realizedAt:  realizedAt,
	}
}

func Delete() Transaction {
	return Transaction{
		deletedAt: time.Now(),
	}
}

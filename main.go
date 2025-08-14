package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	id          int
	name        string
	description string
	value       float64
	insertedAt  time.Time
	realizedAt  time.Time
	deletedAt   time.Time
}

type TransactionStore struct {
	store []Transaction
}

func (transaction Transaction) Insert() bool {
	return true
}

func main() {}

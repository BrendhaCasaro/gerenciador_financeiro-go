package main

import (
	"time"

	"github.com/BrendhaCasaro/gerenciador_financeiro-go/transaction"
)

func main() {
	ts := transaction.TransactionStore{}
	ts.Insert(transaction.NewTransaction("teste", "teste", 100, time.Now()))
}

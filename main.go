package main

import (
	"github.com/BrendhaCasaro/gerenciador_financeiro-go/transaction"
	"time"
)

func main() {
	transaction.Insert(transaction.NewTransaction("teste", "teste", 100, time.Now()))
}

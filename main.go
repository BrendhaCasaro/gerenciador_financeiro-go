package main

import (
	"fmt"
	"time"

	"github.com/BrendhaCasaro/gerenciador_financeiro-go/transaction"
)

func main() {
	ts := transaction.TransactionStore{}
	tx := transaction.NewTransaction("teste", "teste", 100, time.Now())

	ts.Insert(tx)

	store, err := (ts.MarshalJSON())
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf(string(store))
}

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/BrendhaCasaro/gerenciador_financeiro-go/internal/api"
	"github.com/BrendhaCasaro/gerenciador_financeiro-go/transaction"
)

// Listar transação (na mesma tela, é necessário ter a diferença entre as receitas e as despesas, junto com as mesmas)
// Adicionar uma transação
// Deletar uma transação
// Editar uma transação

func main() {
	store := transaction.TransactionStore{}
	server := api.NewServer(store)
	server.ts.Insert(transaction.NewTransaction("teste", "teste", 100, time.Now()))


	http.HandleFunc("GET /health-check", s.HandleHealthCheck)
	http.HandleFunc("GET /transactions", s.HandleListTransactions)

	http.ListenAndServe(":42069", nil)
	log.Println("Server starting on :8080")
}

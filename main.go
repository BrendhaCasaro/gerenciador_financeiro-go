package main

import (
	"log"
	"net/http"

	"github.com/BrendhaCasaro/gerenciador_financeiro-go/internal/api"
	"github.com/BrendhaCasaro/gerenciador_financeiro-go/transaction"
)

// Listar transação (na mesma tela, é necessário ter a diferença entre as receitas e as despesas, junto com as mesmas)
// Adicionar uma transação
// Deletar uma transação
// Achar uma transação específica
// Editar uma transação
// filtrar por tipo e valor

func main() {
	store := transaction.TransactionStore{}
	server := api.NewServer(&store)

	server.DefineRoutes()
	log.Println("Server starting on :42069")
	http.ListenAndServe(":42069", nil)
}

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/BrendhaCasaro/gerenciador_financeiro-go/transaction"
)

// Listar transação (na mesma tela, é necessário ter a diferença entre as receitas e as despesas, junto com as mesmas)
// Adicionar uma transação
// Deletar uma transação
// Editar uma transação

type Server struct {
	ts transaction.TransactionStore
}

func HandleHealthCheck(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello World")
}

func (s *Server) HandleListTransactions(w http.ResponseWriter, _ *http.Request) {
	JSONResponse, err := json.Marshal(s.ts.ListTransactions())
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(JSONResponse)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {
	s := Server{}
	s.ts.Insert(transaction.NewTransaction("teste", "teste", 100, time.Now()))

	http.HandleFunc("GET /health-check", HandleHealthCheck)
	http.HandleFunc("GET /transactions", s.HandleListTransactions)

	http.ListenAndServe(":42069", nil)
	log.Println("Server starting on :8080")
}

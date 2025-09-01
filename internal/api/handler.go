package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/BrendhaCasaro/gerenciador_financeiro-go/transaction"
)

type Server struct {
	ts *transaction.TransactionStore
}

func NewServer(ts *transaction.TransactionStore) *Server {
	return &Server{
		ts: ts,
	}
}

func (s *Server) HandleHealthCheck(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello World")
}

func (s *Server) HandleListTransactions(w http.ResponseWriter, _ *http.Request) {
	jsonResponse, err := json.Marshal(s.ts.ListTransactions())
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func (s *Server) HandleAddTransaction(w http.ResponseWriter, r *http.Request) {
	// receber o body da request
	// converter de json para struct transaction os dados da transação
	// inserir a struct na store
	// retornar 201 da execução
	// retornar no header um campo location que o campo seja o id da transaction

	var tx transaction.Transaction
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, "Error: decoding body", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/BrendhaCasaro/gerenciador_financeiro-go/transaction"
	"github.com/google/uuid"
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

func (s *Server) HandleListTransactions(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var transactions []*transaction.Transaction

	if t := query.Get("type"); t != "" {
		switch t {
		case "income":
			transactions = s.ts.FilterByType(transaction.Income)

		case "expense":
			transactions = s.ts.FilterByType(transaction.Expense)

		default:
			http.Error(w, "Invalid type filter", http.StatusBadRequest)
			return
		}
	}

	if query.Has("init") && query.Has("end") {
		init, err := strconv.ParseFloat(query.Get("init"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter", http.StatusUnprocessableEntity)
			log.Printf("Error to convert to float the init of filter: %v", err)
			return
		}

		end, err := strconv.ParseFloat(query.Get("end"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter", http.StatusUnprocessableEntity)
			log.Printf("Error to convert to float the init of filter: %v", err)
			return
		}

		transactions = s.ts.FilterByValue(init, end)
	}

	if transactions == nil {
		transactions = s.ts.ListTransactions()
	}

	jsonResponse, err := json.Marshal(transactions)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		log.Printf("Error marshalling JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonResponse); err != nil {
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		log.Printf("Error writing response: %v", err)
	}
}

func (s *Server) HandleAddTransaction(w http.ResponseWriter, r *http.Request) {
	var tx transaction.Transaction
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusInternalServerError)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	tInserted := s.ts.Insert(transaction.NewTransaction(tx.Name, tx.Description, tx.Value, tx.RealizedAt))

	w.Header().Set("Location", "/transactions/"+tInserted.Id.String())
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) HandleDeleteTransaction(w http.ResponseWriter, r *http.Request) {
	idReq, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Error parsing ID: %v", err)
		return
	}
	_, err = s.ts.SearchByID(idReq)
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		log.Printf("Transaction not found: %v", err)
		return
	}

	s.ts.SoftDelete(idReq)
}

func (s *Server) HandleFindTransaction(w http.ResponseWriter, r *http.Request) {
	idReq, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Error parsing ID: %v", err)
		return
	}

	transaction, err := s.ts.SearchByID(idReq)
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		log.Printf("Transaction not found: %v", err)
		return
	}

	jsonResponse, err := json.Marshal(transaction)
	if err != nil {
		http.Error(w, "Internal error", http.StatusUnprocessableEntity)
		log.Printf("Error marshalling JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Internal error", http.StatusUnprocessableEntity)
		log.Printf("Error writing response: %v", err)
	}
}

func (s *Server) HandleEditTransaction(w http.ResponseWriter, r *http.Request) {
	idReq, errParse := uuid.Parse(r.PathValue("id"))
	if errParse != nil {
		http.Error(w, "Internal error", http.StatusUnprocessableEntity)
		log.Printf("Error parsing ID: %v", errParse)
		return
	}

	var tx transaction.UpdateFieldsTransaction
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, "Internal error", http.StatusUnprocessableEntity)
		log.Printf("Error to decode the request body to a transaction type %v", err)
		return
	}

	s.ts.EditByID(idReq, tx)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

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
			http.Error(w, "Error to return the transactions", http.StatusUnprocessableEntity)
			log.Printf("Error to parse to float of init filter: %v", err)
			return
		}

		end, err := strconv.ParseFloat(query.Get("end"), 64)
		if err != nil {
			http.Error(w, "Error to return the transactions", http.StatusUnprocessableEntity)
			log.Printf("Error to parse to float of init filter: %v", err)
			return
		}

		transactions = s.ts.FilterByValue(init, end)
	}

	if transactions == nil {
		transactions = s.ts.ListTransactions()
	}

	jsonResponse, err := json.Marshal(transactions)
	if err != nil {
		http.Error(w, "Error serializing transactions", http.StatusInternalServerError)
		log.Printf("Error marshalling json: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonResponse); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		log.Printf("Error writing response: %v", err)
	}
}

func (s *Server) HandleAddTransaction(w http.ResponseWriter, r *http.Request) {
	var tx transaction.Transaction
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, "Error: decoding body", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	tInserted := s.ts.Insert(transaction.NewTransaction(tx.Name, tx.Description, tx.Value, tx.RealizedAt))

	w.Header().Set("Location", "/transactions/"+tInserted.Id.String())
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) HandleDeleteTransaction(w http.ResponseWriter, r *http.Request) {
	idReq, errParse := uuid.Parse(r.PathValue("id"))
	if errParse != nil {
		http.Error(w, "Error to parse the ID", http.StatusUnprocessableEntity)
		log.Printf("Error to parse the ID: %v", errParse)
		return
	}
	_, err := s.ts.SearchByID(idReq)
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		log.Printf("Error to find transaction: %v", err)
		return
	}

	s.ts.SoftDelete(idReq)
}

func (s *Server) HandleFindTransaction(w http.ResponseWriter, r *http.Request) {
	idReq, errParse := uuid.Parse(r.PathValue("id"))
	if errParse != nil {
		http.Error(w, "Error to parse the ID", http.StatusUnprocessableEntity)
		log.Printf("Error to parse the ID: %v", errParse)
		return
	}

	transaction, err := s.ts.SearchByID(idReq)
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		log.Printf("Error to find transaction: %v", err)
		return
	}

	jsonResponse, errMarshal := json.Marshal(transaction)
	if errMarshal != nil {
		http.Error(w, "Error to return a transaction", http.StatusUnprocessableEntity)
		log.Printf("Error to marshal the transaction: %v", errMarshal)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write(jsonResponse)
	if errWrite != nil {
		http.Error(w, "Error to return the transaction", http.StatusUnprocessableEntity)
		log.Printf("Error writing response of request: %v", err)
		return
	}
}

func (s *Server) HandleEditTransaction(w http.ResponseWriter, r *http.Request) {
	idReq, errParse := uuid.Parse(r.PathValue("id"))
	if errParse != nil {
		http.Error(w, "Internal error", http.StatusUnprocessableEntity)
		log.Printf("Error to parse the ID: %v", errParse)
		return
	}

	var tx transaction.UpdateFieldsTransaction
	errDecoder := json.NewDecoder(r.Body).Decode(&tx)
	if errDecoder != nil {
		http.Error(w, "Internal error", http.StatusUnprocessableEntity)
		log.Printf("Error to decode the request body to a transaction type %v", errDecoder)
	}

	s.ts.EditByID(idReq, tx)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

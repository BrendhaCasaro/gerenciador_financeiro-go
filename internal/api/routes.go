package api

import (
	"net/http"
)

func (s *Server) DefineRoutes() {
	http.HandleFunc("GET /health-check", s.HandleHealthCheck)
	http.HandleFunc("GET /transactions", s.HandleListTransactions)
	http.HandleFunc("POST /transactions", s.HandleAddTransaction)
	http.HandleFunc("DELETE /transactions/{id}", s.HandleDeleteTransaction)
	http.HandleFunc("GET /transactions/{id}", s.HandleFindTransaction)
	http.HandleFunc("PATCH /transactions/{id}", s.HandleEditTransaction)
}

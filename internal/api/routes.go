package api

import (
	"net/http"
)

func (s *Server) DefineRoutes() {
	http.HandleFunc("GET /health-check", s.HandleHealthCheck)
	http.HandleFunc("GET /transactions", s.HandleListTransactions)
	http.HandleFunc("POST /transactions", s.HandleAddTransaction)
}

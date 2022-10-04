package routes

import (
	"waysbook/handlers"
	psql "waysbook/pkg/dbConnection"
	"waysbook/pkg/middlewares"
	"waysbook/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(psql.DB)
  	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transaction", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction", middlewares.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction", middlewares.Auth(h.UpdateTransaction)).Methods("PATCH")
}
package routes

import (
	"waysbook/handlers"
	psql "waysbook/pkg/dbConnection"
	"waysbook/pkg/middlewares"
	"waysbook/repositories"

	"github.com/gorilla/mux"
)

func BookPurchasedRoutes(r *mux.Router) {
	bookRepository := repositories.RepositoryBookPurchased(psql.DB)
	h := handlers.HanlderBookPurchasedRepository(bookRepository)

	r.HandleFunc("/purchased", middlewares.Auth(h.FindBooksPurcased)).Methods("GET")
	r.HandleFunc("/add-purchased", middlewares.Auth(h.CreateBookPurchased)).Methods("POST")
}
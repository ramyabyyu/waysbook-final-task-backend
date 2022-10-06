package routes

import (
	"waysbook/handlers"
	psql "waysbook/pkg/dbConnection"
	"waysbook/pkg/middlewares"
	"waysbook/repositories"

	"github.com/gorilla/mux"
)

func BookPurchasedRoutes(r *mux.Router) {
	bookPurchasedRepository := repositories.RepositoryBookPurchased(psql.DB)
	h := handlers.HanlderBookPurchased(bookPurchasedRepository)

	r.HandleFunc("/add-book-purchased", h.CreateBookPurchased).Methods("POST")
	r.HandleFunc("/book-purchaseds", middlewares.Auth(h.FindBooksPurcased)).Methods("GET")
}
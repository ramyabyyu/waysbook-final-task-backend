package routes

import (
	"waysbook/handlers"
	psql "waysbook/pkg/dbConnection"
	"waysbook/pkg/middlewares"
	"waysbook/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(psql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", middlewares.Auth(h.FindCartsByUserID)).Methods("GET")
	r.HandleFunc("/cart", middlewares.Auth(h.AddCart)).Methods("POST")
	r.HandleFunc("/cart", middlewares.Auth(h.DeleteCart)).Methods("DELETE")
}
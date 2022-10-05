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
	h := handlers.HanlderCart(cartRepository)

	r.HandleFunc("/carts", middlewares.Auth(h.FindCartItems)).Methods("GET")
	r.HandleFunc("/cart", middlewares.Auth(h.CreateCartItem)).Methods("POST")
}
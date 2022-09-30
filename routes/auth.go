package routes

import (
	"waysbook/handlers"
	psql "waysbook/pkg/dbConnection"
	"waysbook/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authRepository := repositories.RepositoryAuth(psql.DB)
	h := handlers.HandlerAuth(authRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/become-seller", h.BecomeSeller).Methods("POST")
}
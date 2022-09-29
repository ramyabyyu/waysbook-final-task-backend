package routes

import (
	"waysbook/handlers"
	psql "waysbook/pkg/dbConnection"
	"waysbook/pkg/middlewares"
	"waysbook/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(psql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.GetAllUser).Methods("GET")
	r.HandleFunc("/user", middlewares.Auth(h.GetUserByID)).Methods("GET")
}
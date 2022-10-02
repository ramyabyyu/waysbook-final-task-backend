package routes

import "github.com/gorilla/mux"

func RoutesInit(r *mux.Router) {
	AuthRoutes(r)
	BookRoutes(r)
	UserRoutes(r)
	CartRoutes(r)
}
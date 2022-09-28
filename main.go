package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"waysbook/database"
	psql "waysbook/pkg/dbConnection"
	"waysbook/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// ENV config
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load .env file")
	}

	// Database Init
	psql.DatabaseInit()

	// Run Migration
	database.RunMigration()

	// Mux Router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Congrats! Your Dumbass API is now setup!")
	})

	routes.RoutesInit(r.PathPrefix("/api/v1").Subrouter())

	// Config CORS
	var allowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var allowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD"})
	var allowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = os.Getenv("PORT");
	fmt.Println("Your server at http://localhost:"+port)
	http.ListenAndServe(":" + port, handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(r))
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"waysbook/database"
	psql "waysbook/pkg/dbConnection"

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

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Congrats! Your Dumbass API is now setup!")
	})

	port := "8080"
	fmt.Println("Your server at http://localhost:"+port)
	http.ListenAndServe(":"+port, r)
}
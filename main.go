package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

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
package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func IsSeller(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
		isSeller := userInfo["is_seller"]

		if isSeller != true {
			w.WriteHeader(http.StatusForbidden)
			response := Result{Code: http.StatusForbidden, Message: "Sorry, you can't access this page"}
			json.NewEncoder(w).Encode(response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
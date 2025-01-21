package auth

import (
	"fmt"
	"net/http"
	"os"
)
func AuthMiddleware(next http.Handler) http.Handler {

	var API_KEY = os.Getenv("API_KEY")

	fmt.Println(API_KEY)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		apiKey := r.Header.Get("Authorization")
		
		if apiKey != API_KEY {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
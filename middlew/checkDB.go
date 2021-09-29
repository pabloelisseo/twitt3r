package middlew

import (
	"net/http"

	"github.com/pabloelisseo/twitt3r/bd"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() != nil {
			http.Error(w, "Database connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

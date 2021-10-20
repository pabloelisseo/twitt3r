package middlew

import (
	"net/http"

	"github.com/pabloelisseo/twitt3r/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() != nil {
			http.Error(w, "Database connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

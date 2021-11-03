package middleware

import (
	"net/http"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*if bd.CheckConnection() == 0 {
			http.Error(w, "Sin conexion a DB", http.StatusInternalServerError)
			return
		}*/
		next(w, r)
	}
}

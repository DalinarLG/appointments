package middlewares

import (
	"DalinarLG/appointments/db"
	"net/http")


func CheckDb(next http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckDB() {
			http.Error(w, "Database error", 400)
			return
		}
		next.ServeHTTP(w, r)
	}
}
package middlewares

import (
	"DalinarLG/appointments/routers"
	"net/http"
)

func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := routers.ProccessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error: "+err.Error(), 400)
			return
		}

		next.ServeHTTP(w, r)
	}
}

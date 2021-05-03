package handlers

import (
	"DalinarLG/appointments/middlewares"
	"DalinarLG/appointments/routers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)



func Handlers(){
	r := mux.NewRouter()

	r.HandleFunc("/registeruser", middlewares.CheckDb(routers.UserRegister)).Methods("POST")
	r.HandleFunc("/loginuser", middlewares.CheckDb(routers.Userlogin)).Methods("GET")
	r.HandleFunc("deluser", middlewares.CheckDb(middlewares.ValidateToken(routers.DelUser))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(r)

	http.ListenAndServe(":"+PORT, handler)
}
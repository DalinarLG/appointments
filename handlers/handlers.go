package handlers

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)



func Handlers(){
	r := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(r)

	http.ListenAndServe(":"+PORT, handler)
}
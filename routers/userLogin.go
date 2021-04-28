package routers

import (
	"DalinarLG/appointments/db"
	"DalinarLG/appointments/jwt"
	"DalinarLG/appointments/models"
	"encoding/json"
	"net/http"
	"text/template"
	"time"
)



func Userlogin(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error: "+err.Error(),400)
		return
	}

	user.Password = template.HTMLEscapeString(user.Password)


	user, status := db.Userlogin(user.Cedula, user.Password)
	if !status {
		http.Error(w, "Cedula or password incorrect", 400)
		return
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		http.Error(w, "Error generating token "+err.Error(), 400)
		return
	}

	tokenResponse := models.Userlogin {
		Token: token,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tokenResponse)

	http.SetCookie(w,&http.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(24*time.Hour),
	})
}
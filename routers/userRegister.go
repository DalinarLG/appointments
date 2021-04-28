package routers

import (
	"DalinarLG/appointments/db"
	"DalinarLG/appointments/models"
	"encoding/json"
	"net/http"

	"text/template"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	if len(user.Name) == 0 {
		http.Error(w, "Name is required", 400)
		return
	}

	if len(user.Lastname) == 0 {
		http.Error(w, "Lastname is required", 400)
		return
	}

	if len(user.Password) == 0 || len(user.Password) < 6 {
		http.Error(w, "Password must have 6 characters", 400)
		return
	}

	user.Name = template.HTMLEscapeString(user.Name)
	user.Lastname = template.HTMLEscapeString(user.Lastname)
	user.Password = template.HTMLEscapeString(user.Password)

	_, found, _ := db.Checkced(user.Cedula)
	if found {
		http.Error(w, "User already exists", 400)
		return
	}

	status, err := db.UserRegister(user)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Could no register user: "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

package routers

import (
	"DalinarLG/appointments/db"	
	"net/http"
	"strconv"
)

func DelUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	ced_id := r.URL.Query().Get("ced_id")
	if len(ced_id) < 1 {
		http.Error(w, "Cédula or Id required ", 400)
		return
	}
	cedid, _ := strconv.Atoi(ced_id)

	user,_, err := db.Checkced(UserID)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	if user.Role != 1 {
		http.Error(w, "User not authorized for this operation", 400)
		return
	}
	
	err = db.DelUser(cedid)	
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

}

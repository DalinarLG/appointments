package validations


import (	
	"DalinarLG/appointments/db"
)

func RoleValidation(ced_id int)bool {
	u, _, err := db.Checkced(ced_id)
	if err != nil {		
		return false
	}

	if u.Role != 1 {		
		return false
	}

	return true
}
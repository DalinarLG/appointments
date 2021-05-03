package db

import (
	"DalinarLG/appointments/models"	
	"log"
)


func Checkced(ced_id int )(models.User, bool, error){
	var user models.User
	stmt, err := DB.Prepare("select * from user where cedula = ? or id = ?")
	if err != nil{
		log.Println("Error: "+err.Error())
		return user, false, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(ced_id, ced_id).Scan(&user.ID, &user.Cedula, &user.Name, &user.Lastname, &user.Cellphone, &user.Role, &user.Password)
	if err != nil {
		log.Println("Error: "+err.Error())
		return user, false, err
	}

	return user, true,  nil
}
package db

import (
	"DalinarLG/appointments/models"	
	"log"
)


func UserRegister(u models.User)(bool, error){
	stmt, err := DB.Prepare("INSERT INTO user (cedula, name, lastname, celphone,role, password) values (?,?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	defer stmt.Close()
	u.Password = Encrypt(u.Password)
	_, err = stmt.Exec(u.Cedula, u.Name, u.Lastname, u.Cellphone, u.Role, u.Password)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}
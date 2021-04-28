package db

import (
	"DalinarLG/appointments/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)


func Userlogin(ced int, pass string)(models.User, bool){

	user, found, _ := Checkced(ced)
	if !found {
		log.Println("User not found")
		return user, false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		log.Println(err.Error())
		return user, false
	}

	return user, true
}
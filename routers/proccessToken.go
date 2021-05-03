package routers

import (
	"DalinarLG/appointments/db"
	"DalinarLG/appointments/models"
	"errors"	
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var UserID int
var UserCED int

func ProccessToken(t string) (*models.Claims, bool, error) {
	myKey := []byte("jetstream960")
	claims := &models.Claims{}
	splitToken := strings.Split(t, "Bearer")	

	if len(splitToken) != 2 {
		return claims, false, errors.New("error proccessing token split")
	}

	token := strings.TrimSpace(splitToken[1])

	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token)(interface{}, error){
		return myKey, nil
	})
	
	if err == nil {
		_, found, _ := db.Checkced(claims.ID)		
		if found {
			UserID = claims.ID
			UserCED = claims.Cedula
		}

		return claims, true, nil
	}

	if !tk.Valid {
		return claims, false, errors.New("error proccessing token tk valid")
	}

	return claims, false, err
}

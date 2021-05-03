package jwt

import (
	"DalinarLG/appointments/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(u models.User) (string, error) {
	mykey := []byte("jetstream960")
	payload := jwt.MapClaims{
		"id":       u.ID,
		"name":     u.Name,
		"cedula":   u.Cedula,
		"lastname": u.Lastname,
		"role":     u.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(mykey)
	if err != nil {
		return "", err
	}

	return signedToken, nil

}

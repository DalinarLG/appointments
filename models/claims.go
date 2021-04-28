package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID     int `json:"id,omitempty"`
	Cedula int `json:"cedula,omitempty"`
	jwt.StandardClaims
}

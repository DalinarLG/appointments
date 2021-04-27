package models

type User struct {
	ID        int    `json:"id"`
	Cedula    int    `json:"cedula,omitempty"`
	Name      string `json:"name,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Cellphone int    `json:"cellphone,omitempty"`
	Role      int    `json:"role,omitempty"`
	Password  string `json:"password,omitempty"`
}

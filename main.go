package main

import (
	"DalinarLG/appointments/db"
	"DalinarLG/appointments/handlers"
	"log"
)


func main() {
	if !db.CheckDB() {
		log.Fatal("No database connection")
		return
	}

	handlers.Handlers()
}
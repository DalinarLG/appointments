package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB = getDB()

func getDB() *sql.DB {
	dbName := "appointments"
	dbPass := ""
	dbUser := "root"

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal("Error " + err.Error())
	}

	log.Println("Database Connected...")
	return db
}

func CheckDB() bool {
	err := DB.Ping()
	if err != nil {
		return false
	}
	return err == nil
}

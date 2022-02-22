package database

import (
	"database/sql"
	"log"
)

// databse object for the connection
var DBConn *sql.DB

func SetupDatabase() {
	var err error
	DBConn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/studentdb")
	if err != nil {
		log.Fatal(err)
	}

	defer DBConn.Close()

}

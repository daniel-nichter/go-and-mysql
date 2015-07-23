package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: 05-little-bobby-tables NAME")
	}
	db, err := sql.Open("mysql", "daniel:foo@tcp(localhost:3306)/?interpolateParams=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT user FROM mysql.user WHERE user = ?", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Throw away the reuslts. We're only trying to break interpolateParams.
}

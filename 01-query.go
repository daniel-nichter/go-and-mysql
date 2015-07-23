package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
	9 Connect	daniel@localhost on
	9 Query		SELECT @@max_allowed_packet
	9 Query		SELECT user FROM mysql.user
	9 Quit
*/

func main() {
	db, err := sql.Open("mysql", "daniel:foo@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT user FROM mysql.user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user string
	for rows.Next() {
		err := rows.Scan(&user)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

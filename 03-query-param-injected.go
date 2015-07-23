package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
	13 Connect	daniel@localhost on
	13 Query	SELECT @@max_allowed_packet
	13 Query	SELECT user FROM mysql.user WHERE user = 'daniel'
	13 Quit
*/

func main() {
	db, err := sql.Open("mysql", "daniel:foo@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT user FROM mysql.user WHERE user = 'daniel'")
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

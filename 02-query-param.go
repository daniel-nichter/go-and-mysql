package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
	12 Connect		daniel@localhost on
	12 Query		SELECT @@max_allowed_packet
	12 Prepare		SELECT user FROM mysql.user WHERE user = ?
	12 Execute		SELECT user FROM mysql.user WHERE user = 'daniel'
	12 Close stmt
	12 Quit
*/

func main() {
	db, err := sql.Open("mysql", "daniel:foo@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user := "daniel"
	rows, err := db.Query("SELECT user FROM mysql.user WHERE user = ?", user)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// reuse user var
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

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
	20 Connect	daniel@localhost on
	20 Query	SELECT @@max_allowed_packet
	20 Query	SELECT user FROM mysql.user WHERE user = 'daniel'
	20 Quit
*/

func main() {
	db, err := sql.Open("mysql", "daniel:foo@tcp(localhost:3306)/?interpolateParams=true")
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

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
	9 Connect	daniel@localhost on
	9 Query		SELECT @@max_allowed_packet
	9 Query		START TRANSACTION
	9 Query		SELECT 1
	9 Query		SELECT 2
	9 Query		COMMIT
	9 Quit
*/

func main() {
	db, err := sql.Open("mysql", "daniel:foo@tcp(localhost:3306)/?interpolateParams=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	var number int
	tx.QueryRow("SELECT 1").Scan(&number)
	log.Println(number)
	tx.QueryRow("SELECT 2").Scan(&number)
	log.Println(number)

	// Commit can fail with synchronous replication
	// (Percona XtraDB Cluster and MariaDB Galera Cluster)
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

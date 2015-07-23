package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// This code has a leak and will die after 16,382 executions. Imagine this is
// your long-running API. ~16k executions could take  minutes, hours, or days.
// Multiple by N-many servers, and the failure seems to appear randomly without
// a single, unifying cause. The fix is extremely simple.

func main() {
	db, err := sql.Open("mysql", "daniel:foo@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	n := 0
	for {
		if err := doImportantWork(db); err != nil {
			log.Fatal("Failed after doing important work only ", n, " times :-(")
		}
		n++
	}
}

func doImportantWork(db *sql.DB) error {
	stmt, err := db.Prepare("SELECT 'something important'")
	if err != nil {
		return err
	}

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var something string
		err := rows.Scan(&something)
		if err != nil {
			return err
		}
		// throw away the result
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
